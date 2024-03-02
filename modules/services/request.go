package services

import (
	"github.com/sebigdev/walletpay/modules/dto"
	"github.com/sebigdev/walletpay/modules/entities"
	"github.com/sebigdev/walletpay/modules/repositories"
	"github.com/sebigdev/walletpay/modules/vo"
)

type IPaymentRequestService interface {
	Request(userId string, pr dto.CreatePayRequest) (string, error)
	AcknowledgeRequest(userId, requestId string) error
}

type paymentRequestService struct {
	PaymentRequestRepo repositories.IPaymentRequestRepository
	UserRepo           repositories.IUserRepository
}

func NewPaymentRequestService(repo repositories.IPaymentRequestRepository, userRepo repositories.IUserRepository) IPaymentRequestService {
	return &paymentRequestService{
		PaymentRequestRepo: repo,
		UserRepo:           userRepo,
	}
}

func (pr *paymentRequestService) Request(userId string, req dto.CreatePayRequest) (string, error) {
	person, err := pr.UserRepo.GetUserById(userId)
	if err != nil {
		return "", err
	}
	wallet, err := vo.NewWalletNumber(req.CreditorWallet)
	if err != nil {
		return "", err
	}

	currency, err := vo.NewCurrency(req.Currency)
	if err != nil {
		return "", err
	}

	amount, err := vo.NewAmount(req.Amount, false)
	if err != nil {
		return "", err
	}

	if err := person.VerifyPin(req.Pin); err != nil {
		return "", err
	}

	pRequest := entities.NewPayRequest(wallet, vo.Owner(req.RequestPartyId), amount, currency)

	resp, err := pr.PaymentRequestRepo.MakeRequest(pRequest)
	if err != nil {
		return "", err
	}
	return resp, nil
}

func (pr *paymentRequestService) AcknowledgeRequest(userId, requestId string) error {
	_, err := pr.UserRepo.GetUserById(userId)
	if err != nil {
		return err
	}

	if err := pr.PaymentRequestRepo.AcknowlegeRequest(requestId); err != nil {
		return err
	}

	return nil
}
