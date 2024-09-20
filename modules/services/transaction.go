package services

import (
	"github.com/sebigdev/walletpay/modules/dto"
	"github.com/sebigdev/walletpay/modules/entities"
	"github.com/sebigdev/walletpay/modules/repositories"
	"github.com/sebigdev/walletpay/modules/responses"
	"github.com/sebigdev/walletpay/modules/vo"
)

type ITransactionService interface {
	Initiate(userId string, trx dto.CreateTransaction) (responses.TransactionResponse, error)
	GetTransactions(userId string) ([]responses.TransactionFullResponse, error)
	SubmitTransaction(transaction *entities.Transaction) (responses.TransactionResponse, error)
}

type TransactionService struct {
	TransactionRepo repositories.ITransactionRepository
	UserRepo        repositories.IUserRepository
	EventBus        *EventBus
}

func NewTransactionService(repo repositories.ITransactionRepository, userRepo repositories.IUserRepository,
	eventBus *EventBus) ITransactionService {
	return &TransactionService{
		TransactionRepo: repo,
		UserRepo:        userRepo,
		EventBus:        eventBus,
	}
}

func (ts *TransactionService) Initiate(userId string, trx dto.CreateTransaction) (responses.TransactionResponse, error) {

	owner, err := ts.UserRepo.GetUserById(userId)
	if err != nil {
		return responses.TransactionResponse{}, err
	}
	toWallet, err := vo.NewWalletNumber(trx.CreditorWalletAddress)
	if err != nil {
		return responses.TransactionResponse{}, err
	}
	fromWallet, err := vo.NewWalletNumber(trx.DebitorWalletAddress)
	if err != nil {
		return responses.TransactionResponse{}, err
	}
	toCurr, err := vo.NewCurrency(trx.CreditorCurrency)
	if err != nil {
		return responses.TransactionResponse{}, err
	}
	fromCurr, err := vo.NewCurrency(trx.DebitorCurrency)
	if err != nil {
		return responses.TransactionResponse{}, err
	}
	amount, err := vo.NewAmount(trx.Amount, false)
	if err != nil {
		return responses.TransactionResponse{}, err
	}
	creditorName, err := vo.NewCreditor(trx.CreditorName)
	if err != nil {
		return responses.TransactionResponse{}, err
	}
	if err := owner.VerifyPin(trx.Pin); err != nil {
		return responses.TransactionResponse{}, err
	}

	convertedMoney := vo.NewMoneyConverted(amount, toCurr, fromCurr)
	toMoney := vo.NewMoney(vo.Amount(convertedMoney), toCurr)
	fromMoney := vo.NewMoney(amount, fromCurr)

	creditor, err := ts.UserRepo.GetWalletOwner(toWallet.String())
	if err != nil {
		return responses.TransactionResponse{}, err
	}

	//CLEANED UP TRANSACTION OBJECTS
	beneficiary := entities.NewBeneficiary(creditor.GetUserID(), toWallet, *toMoney, creditorName)
	debitor := entities.NewDebitor(owner, fromWallet, *fromMoney)

	transaction := entities.NewTransaction(*debitor, *beneficiary, trx.Description)

	//DEBIT WALLET
	if err := owner.Withdraw(*fromMoney, fromWallet); err != nil {
		return responses.TransactionResponse{}, err
	}

	//CREDIT WALLET
	if err := creditor.Deposit(*toMoney, toWallet); err != nil {
		return responses.TransactionResponse{}, err
	}

	if err := ts.UserRepo.UpdatePerson(owner.MapToDao()); err != nil {
		return responses.TransactionResponse{}, err
	}

	if err := ts.UserRepo.UpdatePerson(creditor.MapToDao()); err != nil {
		return responses.TransactionResponse{}, err
	}

	ts.EventBus.Publish(*ToEvent(*transaction, TransactionCreated))

	return responses.TransactionResponse{
		Status:  "Processing",
		Message: "Transaction is being processed",
	}, nil

}

func (ts *TransactionService) GetTransactions(userId string) ([]responses.TransactionFullResponse, error) {
	owner, err := ts.UserRepo.GetUserById(userId)
	if err != nil {
		return []responses.TransactionFullResponse{}, err
	}
	daos, err := ts.TransactionRepo.GetTransaction(owner.GetUserID())

	if err != nil {
		return []responses.TransactionFullResponse{}, err
	}

	var transactions []responses.TransactionFullResponse

	for _, p := range *daos {
		transactions = append(transactions, *entities.ToResponse(p, userId))
	}
	return transactions, nil
}

func (ts *TransactionService) SubmitTransaction(transaction *entities.Transaction) (responses.TransactionResponse, error) {
	//SUBMIT PAYMENT

	if err := ts.TransactionRepo.Submit(transaction); err != nil {
		return responses.TransactionResponse{}, err
	}

	return responses.TransactionResponse{
		Status:  "Processed",
		Message: "Transaction processed successfully",
	}, nil
}
