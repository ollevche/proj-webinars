package main

import (
	"image"
	"webinar/pkg/imageprocessing"
	"webinar/pkg/order"
)

func main() {
	var i image.Image

	i = imageprocessing.ConvertToGray(i)
	i = imageprocessing.Resize(i, 100, 200)

	var ei = imageprocessing.NewEditableImage(i)

	ei.ApplyBlackAndWhiteFilter()
	ei.ConvertToGray()
	i = ei.Finalize()

	var orderStore order.OrderStore // = store.New()

	var notifier order.Notifier // = notifier.New()

	orderProcessor := order.NewProcessor(orderStore, notifier)

	orderProcessor.ProcessOrder(order.ProcessingOrder{})

	// orderCreator := order.NewCreator(orderStore)

	// orderResource := rest.NewOrderResource(orderProcessor)

	// create mux router

	// orderResource.RegisterRoutes(muxRouter)
}
