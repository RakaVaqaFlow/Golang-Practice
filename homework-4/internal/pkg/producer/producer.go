package producer

import (
	"homework/internal/config"
	"homework/internal/model"
	"math/rand"
)

func Orders() <-chan model.Order {
	result := make(chan model.Order, config.NumberOfOrders)
	go func() {
		defer close(result)
		for i := 0; i < config.NumberOfOrders; i++ {
			result <- model.Order{
				ID:        model.OrderID(uint64(i)),
				ProductID: model.ProductID(rand.Uint64()),
				Tracking:  make([]model.OrderTracking, 0, 3),
			}
		}
	}()
	return result
}
