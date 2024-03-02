package services

import (
	"github.com/sebigdev/walletpay/modules/dto"
	"github.com/sebigdev/walletpay/modules/entities"
	"github.com/sebigdev/walletpay/modules/repositories"
	"github.com/sebigdev/walletpay/modules/responses"
	"github.com/sebigdev/walletpay/modules/vo"
)

type ITransactionService interface {
	Submit(userId string, trx dto.CreateTransaction) (responses.TransactionResponse, error)
	GetTransactions(userId string) ([]responses.TransactionFullResponse, error)
}

type TransactionService struct {
	TransactionRepo repositories.ITransactionRepository
	UserRepo        repositories.IUserRepository
}

func NewTransactionService(repo repositories.ITransactionRepository, userRepo repositories.IUserRepository) ITransactionService {
	return &TransactionService{
		TransactionRepo: repo,
		UserRepo:        userRepo,
	}
}

func (ts *TransactionService) Submit(userId string, trx dto.CreateTransaction) (responses.TransactionResponse, error) {

	person, err := ts.UserRepo.GetUserById(userId)
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
	if err := person.VerifyPin(trx.Pin); err != nil {
		return responses.TransactionResponse{}, err
	}

	convertedMoney := vo.NewMoneyConverted(amount, toCurr, fromCurr)
	toMoney := vo.NewMoney(vo.Amount(convertedMoney), toCurr)
	fromMoney := vo.NewMoney(amount, fromCurr)

	//CLEANED UP TRANSACTION OBJECTS
	beneficiary := entities.NewBeneficiary(toWallet, *toMoney, creditorName)
	debitor := entities.NewDebitor(person, fromWallet, *fromMoney)

	transaction := entities.NewTransaction(*debitor, *beneficiary, trx.Description)

	//DEBIT WALLET
	if err := person.Withdraw(*fromMoney, fromWallet); err != nil {
		return responses.TransactionResponse{}, err
	}
	//CREDIT WALLET
	if err := person.Deposit(*toMoney, toWallet); err != nil {
		return responses.TransactionResponse{}, err
	}

	//SUBMIT PAYMENT
	if err := ts.TransactionRepo.Submit(transaction); err != nil {
		return responses.TransactionResponse{}, err
	}

	if err := ts.UserRepo.UpdatePerson(person.MapToDao()); err != nil {
		return responses.TransactionResponse{}, err
	}
	return responses.TransactionResponse{
		TransactionReference: transaction.PaymentRef,
		TransactionId:        transaction.ID,
	}, nil
}

func (ts *TransactionService) GetTransactions(userId string) ([]responses.TransactionFullResponse, error) {
	person, err := ts.UserRepo.GetUserById(userId)
	if err != nil {
		return []responses.TransactionFullResponse{}, err
	}
	daos, err := ts.TransactionRepo.GetTransaction(person.GetUserID())

	if err != nil {
		return []responses.TransactionFullResponse{}, err
	}

	var transactions []responses.TransactionFullResponse

	for _, p := range *daos {
		transactions = append(transactions, *entities.ToResponse(p))
	}
	return transactions, nil
}
