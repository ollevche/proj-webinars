package order

import "fmt"

type OrderNotifier interface {
	NotifyAll(order any)
}

type OrderPublisher interface {
	Publish(event any) <-chan error
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
		errChan := s.publisher.Publish("order_processed")
		go func() {
			for err := range errChan {
				if err != nil {
					fmt.Println("Failed to publish order processed event")
				}
			}
		}()
	}

	fmt.Println("Finished order processing")
}
