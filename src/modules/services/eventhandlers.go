package services

import (
	"log"
	"time"
)

var (
	WalletCreated string = "WalletCreated"
	PinCreated    string = "PinCreated"
)

type IEvent interface {
	WalletCreatedEvent | PinCreatedEvent
}

type WalletCreatedEvent struct {
	UserId   string
	Currency string
}

type PinCreatedEvent struct {
	UserId string
	Pin    string
}

func ToEvent[T IEvent](event T, topic string) *Event {
	return &Event{
		Data:      event,
		Topic:     topic,
		Timestamp: time.Now().UTC(),
	}
}

type IEventHandler interface {
	WalletCreatedHandler(event <-chan Event)
	PinCreatedHandler(eventData <-chan Event)
}

type eventHandler struct {
	WalletService IWalletService
	UserService   IUserService
}

func NewEventHandler(walletService IWalletService, userService IUserService) IEventHandler {
	return &eventHandler{
		UserService:   userService,
		WalletService: walletService,
	}
}

func (e *eventHandler) WalletCreatedHandler(eventData <-chan Event) {
	log.Println("Wallet created event listening...")
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
	log.Println("Pin created event listening...")
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
