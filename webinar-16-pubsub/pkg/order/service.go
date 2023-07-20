package order

import "fmt"

type OrderNotifier interface {
	NotifyAll(order any)
}

type Service struct {
	notifier OrderNotifier
}

func NewService(on OrderNotifier) *Service {
	return &Service{
		notifier: on,
	}
}

func (s *Service) ProcessOrder() {

	// validation

	// save to database

	// other order processing logic

	s.notifier.NotifyAll("order_processed")

	fmt.Println("Finished order processing")
}
