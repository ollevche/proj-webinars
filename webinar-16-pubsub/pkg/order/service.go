package order

import "fmt"

type OrderNotifier interface {
	NotifyAll(order any)
}

type OrderPublisher interface {
	Publish(event any)
}

type Service struct {
	notifier  OrderNotifier
	publisher OrderPublisher
}

func NewService(on OrderNotifier, op OrderPublisher) *Service {
	return &Service{
		notifier:  on,
		publisher: op,
	}
}

func (s *Service) ProcessOrder() {

	// validation

	// save to database

	// other order processing logic

	if s.notifier != nil {
		s.notifier.NotifyAll("order_processed")
	}

	if s.publisher != nil {
		s.publisher.Publish("order_processed")
	}

	fmt.Println("Finished order processing")
}
