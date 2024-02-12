package services

import (
	"CrashCourse/GoApp/src/modules/dto"
	"CrashCourse/GoApp/src/modules/entities"
	"CrashCourse/GoApp/src/modules/repositories"
	"fmt"
)

type IWalletService interface {
	AddWallet(userId string) error
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

func (ws *walletService) AddWallet(userId string) error {
	person, err := ws.UserRepository.GetUserById(userId)
	if err != nil {
		return err
	}

	newWallet, err := entities.NewWallet(person.GetUserID())
	if err != nil {
		return err
	}
	for _, wa := range *person.GetWallets() {
		if walletAlreayExist(*newWallet, wa) {
			return fmt.Errorf("wallet of type %s and currency %s already created", wa.Type, wa.Amount.Currency)
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

	if err := person.Deposit(depositReq.Amount, depositReq.WalletNumber); err != nil {
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
	if err := person.Withdraw(withdrawReq.Amount, withdrawReq.WalletNumber); err != nil {
		return err
	}

	dao := person.MapToDao()
	err = ws.UserRepository.UpdatePerson(dao)

	if err != nil {
		return err
	}
	return nil
}

func walletAlreayExist(rWallet entities.Wallet, nWallet entities.Wallet) bool {
	return rWallet.Type == nWallet.Type && rWallet.Amount.Currency == nWallet.Amount.Currency
}
