package services

import (
	"sync"
	"time"
)

type Event struct {
	Data      interface{}
	Topic     string
	Timestamp time.Time
}

type EventBus struct {
	rm          sync.RWMutex
	subscribers map[string][]chan<- Event
	closed      bool
}

func NewEventBus() *EventBus {
	return &EventBus{
		subscribers: make(map[string][]chan<- Event),
	}
}

func (eb *EventBus) Subscribe(topic string, data chan Event) {
	eb.rm.Lock()
	defer eb.rm.Unlock()
	eb.subscribers[topic] = append(eb.subscribers[topic], data)
}

func (eb *EventBus) Publish(data Event) {
	eb.rm.RLock()
	defer eb.rm.RUnlock()

	if eb.closed {
		return
	}

	for _, channel := range eb.subscribers[data.Topic] {
		go func(channel chan<- Event) {
			channel <- data
		}(channel)
	}

}

func (eb *EventBus) Close() {
	eb.rm.Lock()
	defer eb.rm.Unlock()

	if !eb.closed {
		eb.closed = true
		for _, subs := range eb.subscribers {
			for _, ch := range subs {
				close(ch)
			}
		}
	}
}

func New(eb *EventBus, opts ...Option) {
	busOpt := &BusOptions{}
	for _, option := range opts {
		option(busOpt)
	}

	handler := NewEventHandler(busOpt.WalletService, busOpt.UserService, busOpt.TransactionService)
	makeSub := func(topic string, buf int64) chan Event {
		channel := make(chan Event, buf)
		eb.Subscribe(topic, channel)
		return channel
	}

	//HANDLERS
	walletChannel := makeSub(WalletCreated, 1)
	userChannel := makeSub(PinCreated, 1)
	transactionChannel := makeSub(TransactionCreated, 1)

	//GOROUTINES
	go handler.WalletCreatedHandler(walletChannel)
	go handler.PinCreatedHandler(userChannel)
	go handler.NotifyTransactionHandler(transactionChannel)

	time.Sleep(time.Second)
}

type Option func(*BusOptions)

func WithWalletService(walletService IWalletService) Option {
	return func(bo *BusOptions) {
		bo.WalletService = walletService
	}
}

func WithUserService(userService IUserService) Option {
	return func(bo *BusOptions) {
		bo.UserService = userService
	}
}

func WithTransactionService(trxService ITransactionService) Option {
	return func(bo *BusOptions) {
		bo.TransactionService = trxService
	}
}

type BusOptions struct {
	WalletService      IWalletService
	UserService        IUserService
	TransactionService ITransactionService
}
