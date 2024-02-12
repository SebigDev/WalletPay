package daos

import (
	r "CrashCourse/GoApp/src/modules/user/responses"
	"CrashCourse/GoApp/src/modules/wallet/entities"

	"github.com/google/uuid"
)

type WalletDao struct {
	Number  string    `bson:"number"`
	Balance float64   `bson:"balance"`
	Owner   uuid.UUID `bson:"owner"`
	Type    string    `bson:"type"`
}

func MapToDaoFrom(wallets []entities.Wallet) *[]WalletDao {
	var daos []WalletDao
	for _, wallet := range wallets {
		daos = append(daos, *ToDao(&wallet))
	}
	return &daos
}

func MapDaoToResponse(wallets []WalletDao) *[]r.WalletResponse {
	var res []r.WalletResponse
	for _, wallet := range wallets {
		res = append(res, r.WalletResponse{
			Number:  wallet.Number,
			Balance: wallet.Balance,
			Type:    wallet.Type,
		})
	}
	return &res
}

func ToDao(w *entities.Wallet) *WalletDao {
	return &WalletDao{
		Number:  w.Number,
		Balance: w.Balance,
		Owner:   w.Owner,
		Type:    string(w.Type),
	}
}
