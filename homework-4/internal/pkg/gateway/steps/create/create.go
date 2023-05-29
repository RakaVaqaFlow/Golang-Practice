package create

import (
	"context"
	"time"

	"homework/internal/model"
)

type Implementation struct{}

func New() *Implementation {
	return &Implementation{}
}

func (i *Implementation) Create(workerID model.WorkerID, order model.Order) (model.Order, error) {
	// Присваиваем заказу workerID, который его обрабатывает
	order.WorkerID = workerID

	// В массив состояний добавляется новое состояние "Создан"
	order.Tracking = append(order.Tracking, model.OrderTracking{
		State: model.OrderStateCreated,
		Start: time.Now().UTC(),
	})

	return order, nil
}

func (i *Implementation) Pipeline(ctx context.Context, workerID model.WorkerID, orderCh <-chan model.Order) <-chan model.PipelineOrder {
	// Создаем канал для отправки созданных заказов
	outCh := make(chan model.PipelineOrder)

	// Запускаем горутину для создания заказов
	go func() {
		defer close(outCh)
		for order := range orderCh {
			orderR, err := i.Create(workerID, order)
			select {
			case <-ctx.Done():
				return
			case outCh <- model.PipelineOrder{
				Order: orderR,
				Err:   err,
			}:
			}
		}
	}()

	return outCh
}
