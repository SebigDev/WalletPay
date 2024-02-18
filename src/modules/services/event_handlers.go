package services

import (
	"CrashCourse/GoApp/src/modules/vo"
	"log"
	"time"
)

var (
	WalletCreated string = "WalletCreated"
)

type WalletCreatedEvent struct {
	UserId    string
	CreatedAt time.Time
	Currency  string
}

type IEventHandler interface {
	WalletCreatedHandler(event <-chan Event)
}

type EventHandler struct {
	WalletService IWalletService
}

func NewEventHandler(walletService IWalletService) IEventHandler {
	return &EventHandler{
		WalletService: walletService,
	}
}

func (e *EventHandler) WalletCreatedHandler(eventData <-chan Event) {
	log.Println("Wallet created event listening...")
	for event := range eventData {
		data, ok := event.Data.(WalletCreatedEvent)
		if !ok {
			log.Println("Invalid event")
			continue
		}
		log.Printf("Creating wallet for user: %s", data.UserId)
		if err := e.WalletService.AddWallet(data.UserId, string(vo.EURO)); err != nil {
			log.Fatal(err)
		}

		log.Println("Wallet created successfully...")
	}
}
