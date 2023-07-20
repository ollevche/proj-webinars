package observer

type Observer interface {
	GetNotified(subject any)
	GetID() string
}

type ObserverRegistrar struct {
	observers map[string]Observer
}

func NewObserverRegistrar() *ObserverRegistrar {
	return &ObserverRegistrar{
		observers: make(map[string]Observer),
	}
}

func (s *ObserverRegistrar) NotifyAll(subject any) {
	for _, o := range s.observers {
		o.GetNotified(subject)
	}
}

func (s *ObserverRegistrar) Register(o Observer) {
	s.observers[o.GetID()] = o
}

func (s *ObserverRegistrar) Unregister(o Observer) {
	delete(s.observers, o.GetID())
}
