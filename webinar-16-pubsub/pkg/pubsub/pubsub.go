package pubsub

type Subscriber interface {
	GetNotified(event any)
	GetID() string
}

type Service struct {
	subscribers chan Subscriber
	events      chan any
}

func NewService() *Service {
	s := &Service{
		subscribers: make(chan Subscriber),
		events:      make(chan any),
	}

	go s.processEvents()

	return s
}

func (s *Service) processEvents() {
	subscribers := make(map[string]Subscriber)

	for {
		select {
		case e := <-s.events:
			for _, sub := range subscribers {
				sub.GetNotified(e)
			}
		case s := <-s.subscribers:
			subscribers[s.GetID()] = s
		}
	}
}

func (s *Service) AddSubscriber(sub Subscriber) {
	s.subscribers <- sub
}

func (s *Service) Publish(event any) {
	s.events <- event
}
