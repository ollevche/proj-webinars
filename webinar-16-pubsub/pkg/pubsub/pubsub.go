package pubsub

import "fmt"

type Subscriber interface {
	GetNotified(event any) error
	GetID() string
}

type message struct {
	event   any
	errChan chan<- error
}

type Service struct {
	subscribers chan Subscriber
	events      chan message
	stop        chan struct{}
}

func NewService() *Service {
	s := &Service{
		subscribers: make(chan Subscriber),
		events:      make(chan message),
		stop:        make(chan struct{}),
	}

	go s.processEvents()

	return s
}

func (s *Service) processEvents() {
	subscribers := make(map[string]Subscriber)

	for {
		select {
		case e := <-s.events:
			fmt.Println("Got event to process")
			for _, sub := range subscribers {
				if err := sub.GetNotified(e.event); err != nil {
					e.errChan <- err
				}
			}
			close(e.errChan)
			fmt.Println("Finished event processing")
		case s := <-s.subscribers:
			subscribers[s.GetID()] = s
		case <-s.stop:
			fmt.Println("Got stop signal")
			s.stop <- struct{}{}
			return
		}
	}
}

func (s *Service) AddSubscriber(sub Subscriber) {
	s.subscribers <- sub
}

func (s *Service) Publish(event any) <-chan error {
	errCh := make(chan error)

	s.events <- message{
		event:   event,
		errChan: errCh,
	}

	return errCh
}

func (s *Service) Stop() {
	fmt.Println("Sending stop signal")
	s.stop <- struct{}{}
	<-s.stop
}
