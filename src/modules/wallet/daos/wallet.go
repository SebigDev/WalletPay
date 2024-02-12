package daos

import (
	"CrashCourse/GoApp/src/modules/user/responses"
	"CrashCourse/GoApp/src/modules/wallet/entities"

	"github.com/google/uuid"
)

type WalletDao struct {
	Number  string    `bson:"number"`
	Balance float64   `bson:"balance"`
	Owner   uuid.UUID `bson:"owner"`
}

func MapToDaoFrom(wallets []entities.Wallet) *[]WalletDao {
	var daos []WalletDao
	for _, wallet := range wallets {
		daos = append(daos, WalletDao(wallet))
	}
	return &daos
}

func MapDaoToResponse(wallets []WalletDao) *[]responses.WalletResponse {
	var resps []responses.WalletResponse
	for _, wallet := range wallets {
		resps = append(resps, responses.WalletResponse{Number: wallet.Number, Balance: wallet.Balance})
	}
	return &resps
}
