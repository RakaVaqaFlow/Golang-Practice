package complete

import (
	"context"
	"fmt"
	"time"

	"homework/internal/model"
)

type Implementation struct{}

func New() *Implementation {
	return &Implementation{}
}

func (i *Implementation) Complete(order model.Order) (model.Order, error) {

	// Проверка обработан ли заказ
	if len(order.Tracking) > 1 && order.Tracking[1].State != model.OrderStateProcessed {
		return order, fmt.Errorf("order is not processed")
	}

	// Инициализируется пункт выдачи заказа - результат суммы ID товара и ID склада
	order.DeliveryPointID = model.DeliveryPointID(uint64(order.ProductID) + uint64(order.WarehouseID))

	// В массив состояний добавляется новое состояние "Завершен"
	order.Tracking = append(order.Tracking, model.OrderTracking{
		State: model.OrderStateComplete,
		Start: time.Now().UTC(),
	})

	return order, nil
}

func (i *Implementation) Pipeline(ctx context.Context, orders <-chan model.PipelineOrder) <-chan model.PipelineOrder {
	// Создаем канал для завершенных заказов
	outCh := make(chan model.PipelineOrder)

	// Запускаем горутину для завершения заказов
	go func() {
		defer close(outCh)
		for order := range orders {

			// Если заказ был обработан с ошибкой, то просто отправляем его в следующий шаг
			if order.Err != nil {
				select {
				case <-ctx.Done():
					return
				case outCh <- order:
				}
			}

			// Завершаем заказ
			orderR, err := i.Complete(order.Order)
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
