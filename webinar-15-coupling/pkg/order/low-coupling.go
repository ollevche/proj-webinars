package order

import "fmt"

type OrderStore interface {
	CreateOrder(ProcessingOrder) error
}

type Notifier interface {
	NotifyUser(ProcessingOrder) error
}

type Processor struct {
	orderStore OrderStore
	notifier   Notifier
}

func NewProcessor(os OrderStore, n Notifier) *Processor {
	return &Processor{
		orderStore: os,
		notifier:   n,
	}
}

type ProcessingOrder struct {
	ID     int
	UserID int
	Items  []string
}

func (op *Processor) ProcessOrder(o ProcessingOrder) error {
	// 1. validate (перевірити чи є оплата)

	// 2. save
	if err := op.orderStore.CreateOrder(o); err != nil {
		return fmt.Errorf("unable to create order: %w", err)
	}

	// 3. notify
	if err := op.notifier.NotifyUser(o); err != nil {
		return fmt.Errorf("unable to notify user: %w", err)
	}

	return nil
}
