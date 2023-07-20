package main

import (
	"fmt"
	"time"
	"webinar/pkg/observer"
	"webinar/pkg/order"
)

type LoggerObserver struct {
}

func (l LoggerObserver) GetNotified(subject any) {
	fmt.Printf("Logger observer got notified about %v\n", subject)
}

func (l LoggerObserver) GetID() string {
	return "logger"
}

type UserNotifierObserver struct {
}

func (l UserNotifierObserver) GetNotified(subject any) {
	time.Sleep(2 * time.Second)
	fmt.Printf("User notifier observer got notified about %v\n", subject)
}

func (l UserNotifierObserver) GetID() string {
	return "user_notifier"
}

func main() {
	orderObserver := observer.NewObserverRegistrar()

	orderObserver.Register(LoggerObserver{})

	orderObserver.Register(UserNotifierObserver{})

	orderService := order.NewService(orderObserver)

	orderService.ProcessOrder()
}
