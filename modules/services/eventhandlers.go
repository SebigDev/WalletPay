package services

import (
	"log"
	"time"

	"github.com/sebigdev/walletpay/modules/entities"
)

var (
	WalletCreated      string = "WalletCreated"
	PinCreated         string = "PinCreated"
	TransactionCreated string = "TransactionCreated"
)

type IEvent interface {
	WalletCreatedEvent | PinCreatedEvent | entities.Transaction
}

type WalletCreatedEvent struct {
	UserId   string
	Currency string
}

type PinCreatedEvent struct {
	UserId string
	Pin    string
}

type TransactionEvent struct {
	AmountCredited    float64
	Currency          string
	DateOfTransaction time.Time
	Receiver          string
	Sender            string
	Direction         string
	SenderId          string
	ReceiverId        string
}

func ToEvent[T IEvent](event T, topic string) *Event {
	return &Event{
		Data:      event,
		Topic:     topic,
		Timestamp: time.Now().UTC(),
	}
}

type IEventHandler interface {
	WalletCreatedHandler(eventData <-chan Event)
	PinCreatedHandler(eventData <-chan Event)
	NotifyTransactionHandler(eventData <-chan Event)
}

type eventHandler struct {
	WalletService      IWalletService
	UserService        IUserService
	TransactionService ITransactionService
}

func NewEventHandler(walletService IWalletService, userService IUserService, trxService ITransactionService) IEventHandler {
	return &eventHandler{
		UserService:        userService,
		WalletService:      walletService,
		TransactionService: trxService,
	}
}

func (e *eventHandler) WalletCreatedHandler(eventData <-chan Event) {
	log.Println("WalletCreatedHandler event listening...")
	for event := range eventData {
		data, ok := event.Data.(WalletCreatedEvent)
		if !ok {
			log.Println("Invalid event")
			continue
		}
		log.Printf("Creating wallet for user: %s", data.UserId)
		if err := e.WalletService.AddWallet(data.UserId, data.Currency); err != nil {
			log.Fatal(err)
		}

		log.Println("Wallet created successfully...")
	}
}

func (e *eventHandler) PinCreatedHandler(eventData <-chan Event) {
	log.Println("PinCreatedHandler event listening...")
	for event := range eventData {
		data, ok := event.Data.(PinCreatedEvent)
		if !ok {
			log.Println("Invalid event")
			continue
		}
		log.Printf("Creating Pin for user: %s", data.UserId)
		if err := e.UserService.AddPin(data.UserId, data.Pin); err != nil {
			log.Fatal(err)
		}

		log.Println("Pin created successfully...")

	}
}

func (e *eventHandler) NotifyTransactionHandler(eventData <-chan Event) {
	log.Println("NotifyTransactionHandler event listening...")
	for event := range eventData {
		data, ok := event.Data.(entities.Transaction)
		if !ok {
			log.Println("Invalid event")
			continue
		}
		log.Printf("TransactionEvent : %+v", data)
		resp, err := e.TransactionService.SubmitTransaction(&data)
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("Response: %+v", resp)
	}
}
