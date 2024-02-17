package services

import (
	"CrashCourse/GoApp/src/modules/dto"
	"CrashCourse/GoApp/src/modules/entities"
	"CrashCourse/GoApp/src/modules/repositories"
	"CrashCourse/GoApp/src/modules/vo"
	"fmt"
)

type IWalletService interface {
	AddWallet(userId, currency string) error
	Deposit(userId string, depositReq dto.DepositRequest) error
	Withdraw(userId string, withdrawReq dto.WithdrawRequest) error
}

type walletService struct {
	UserRepository repositories.IUserRepository
}

func NewWalletService(userRepo repositories.IUserRepository) IWalletService {
	return &walletService{
		UserRepository: userRepo,
	}
}

func (ws *walletService) AddWallet(userId, currency string) error {
	person, err := ws.UserRepository.GetUserById(userId)
	if err != nil {
		return err
	}
	owner, err := vo.NewOwner(person.GetUserID())
	if err != nil {
		return err
	}

	amount, err := vo.NewAmount(0, true)
	if err != nil {
		return err
	}
	curr, err := vo.NewCurrency(currency)
	if err != nil {
		return err
	}

	money := entities.NewMoney(amount, curr)
	newWallet := entities.NewWallet(owner, *money)

	for _, wa := range *person.GetWallets() {
		if walletAlreadyExist(*newWallet, wa) {
			return fmt.Errorf("customer already has a wallet of type %s and currency in %s", wa.Type, wa.Money.Currency)
		}
	}

	person.SetWallet(*newWallet)
	dao := person.MapToDao()
	err = ws.UserRepository.UpdatePerson(dao)

	if err != nil {
		return err
	}
	return nil
}

func (ws *walletService) Deposit(userId string, depositReq dto.DepositRequest) error {
	person, err := ws.UserRepository.GetUserById(userId)
	if err != nil {
		return err
	}

	amount, err := vo.NewAmount(depositReq.Amount, false)
	if err != nil {
		return err
	}

	currency, err := vo.NewCurrency(depositReq.Currency)
	if err != nil {
		return err
	}
	walletNumber, err := vo.NewWalletNumber(depositReq.WalletNumber)
	if err != nil {
		return err
	}

	money := entities.NewMoney(amount, currency)
	if err := person.Deposit(*money, walletNumber); err != nil {
		return err
	}

	dao := person.MapToDao()

	err = ws.UserRepository.UpdatePerson(dao)

	if err != nil {
		return err
	}
	return nil
}

func (ws *walletService) Withdraw(userId string, withdrawReq dto.WithdrawRequest) error {
	person, err := ws.UserRepository.GetUserById(userId)
	if err != nil {
		return err
	}

	amount, err := vo.NewAmount(withdrawReq.Amount, false)
	if err != nil {
		return err
	}
	currency, err := vo.NewCurrency(withdrawReq.Currency)
	if err != nil {
		return err
	}
	walletNumber, err := vo.NewWalletNumber(withdrawReq.WalletNumber)
	if err != nil {
		return err
	}

	money := entities.NewMoney(amount, currency)
	if err := person.Withdraw(*money, walletNumber); err != nil {
		return err
	}

	dao := person.MapToDao()
	err = ws.UserRepository.UpdatePerson(dao)

	if err != nil {
		return err
	}
	return nil
}

func walletAlreadyExist(rWallet entities.Wallet, nWallet entities.Wallet) bool {
	return rWallet.Type == nWallet.Type && rWallet.Money.Currency == nWallet.Money.Currency
}
