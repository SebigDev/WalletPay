package entities

import (
	"CrashCourse/GoApp/src/modules/vo"

	"github.com/google/uuid"
)

type RequestStatus string

var (
	Acknowledged RequestStatus = "Acknowledged"
	Processed    RequestStatus = "Processed"
	Declined     RequestStatus = "Declined"
	Pending      RequestStatus = "P ending"
)

type PayRequest struct {
	Id       string
	To       string
	Amount   float64
	Currency string
	UserId   string
	Status   RequestStatus
}

func NewPayRequest(to vo.WalletNumber, userId vo.Owner, amount vo.Amount, currency vo.Currency) *PayRequest {
	return &PayRequest{
		Id:       uuid.NewString(),
		Status:   Pending,
		To:       to.String(),
		Amount:   float64(amount),
		Currency: currency.String(),
		UserId:   userId.String(),
	}
}

func (p *PayRequest) AcknowledgeRequest() {
	p.Status = Acknowledged
}

func (p *PayRequest) DeclineRequest() {
	p.Status = Declined
}

func (p *PayRequest) CompleteRequest(amount float64, to string) {
	p.Status = Processed
	p.Amount = amount
}
