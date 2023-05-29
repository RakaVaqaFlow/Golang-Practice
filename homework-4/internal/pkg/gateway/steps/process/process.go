package process

import (
	"context"
	"fmt"
	"homework/internal/config"
	"homework/internal/model"
	"time"
)

type Implementation struct{}

func New() *Implementation {
	return &Implementation{}
}

func (i *Implementation) Process(order model.Order) (model.Order, error) {

	// Проверка создан ли заказ
	if len(order.Tracking) > 0 && order.Tracking[0].State != model.OrderStateCreated {
		return order, fmt.Errorf("order is not created")
	}

	// Для наглядности работы пайплайна, добавляем задержку
	time.Sleep(config.ProcessStepDuration * time.Second)

	// Инициализируется склад для заказа - результат взятия ID товара по модулю 2
	order.WarehouseID = model.WarehouseID(uint64(order.ProductID) % 2)

	// В массив состояний добавляется новое состояние "Обработан"
	order.Tracking = append(order.Tracking, model.OrderTracking{
		State: model.OrderStateProcessed,
		Start: time.Now().UTC(),
	})

	return order, nil
}

func (i *Implementation) Pipeline(ctx context.Context, orders <-chan model.PipelineOrder) <-chan model.PipelineOrder {
	// Создаем канал для обработанных заказов
	outCh := make(chan model.PipelineOrder)

	// Запускаем горутину для обработки заказов
	go func() {
		defer close(outCh)
		for order := range orders {

			// Если заказ был создан с ошибкой, то просто отправляем его в следующий шаг
			if order.Err != nil {
				select {
				case <-ctx.Done():
					return
				case outCh <- order:
				}
			}

			// Обрабатываем заказ
			orderR, err := i.Process(order.Order)
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
