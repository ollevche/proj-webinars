package main

import (
	"fmt"
	"time"
	"webinar/pkg/order"
	"webinar/pkg/pubsub"
)

type LoggingSubscriber struct {
}

func (l LoggingSubscriber) GetNotified(subject any) {
	fmt.Printf("Logger subscriber got notified about %v\n", subject)
}

func (l LoggingSubscriber) GetID() string {
	return "logger"
}

type UserNotifierSubscriber struct {
}

func (l UserNotifierSubscriber) GetNotified(subject any) {
	time.Sleep(2 * time.Second)
	fmt.Printf("User notifier subscriber got notified about %v\n", subject)
}

func (l UserNotifierSubscriber) GetID() string {
	return "user_notifier"
}

func main() {
	pubsubService := pubsub.NewService()

	pubsubService.AddSubscriber(LoggingSubscriber{})

	pubsubService.AddSubscriber(UserNotifierSubscriber{})

	orderService := order.NewService(nil, pubsubService)

	orderService.ProcessOrder()

	time.Sleep(3 * time.Second)
}
