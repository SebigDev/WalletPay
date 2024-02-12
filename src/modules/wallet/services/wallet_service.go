package services

import (
	"CrashCourse/GoApp/src/modules/user/repositories"
	"CrashCourse/GoApp/src/modules/wallet/daos"
	"CrashCourse/GoApp/src/modules/wallet/entities"
)

type IWalletService interface {
	AddWallet(userId string) error
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

	newWallet := entities.NewWallet(person.UserId)
	person.Wallets = append(person.Wallets, *daos.ToDao(newWallet))
	err = ws.UserRepository.UpdatePerson(*person)

	if err != nil {
		return err
	}
	return nil
}
