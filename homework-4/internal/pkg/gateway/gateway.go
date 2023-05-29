package gateway

import (
	"context"
	"log"
	"sync"

	"homework/internal/model"
	completestep "homework/internal/pkg/gateway/steps/complete"
	createstep "homework/internal/pkg/gateway/steps/create"
	processstep "homework/internal/pkg/gateway/steps/process"
)

func New(create *createstep.Implementation, process *processstep.Implementation, complete *completestep.Implementation) *Implementation {
	return &Implementation{
		create:   create,
		process:  process,
		complete: complete,
	}
}

type Implementation struct {
	create   *createstep.Implementation
	process  *processstep.Implementation
	complete *completestep.Implementation
}

func (i *Implementation) Process(order model.Order, workerID model.WorkerID) (model.Order, error) {

	// Создание заказа
	orderCreated, err := i.create.Create(workerID, order)
	if err != nil {
		return order, err
	}

	// Обработка заказа
	orderProcessed, err := i.process.Process(orderCreated)
	if err != nil {
		return orderCreated, err
	}

	// Завершение заказа
	orderCompleted, err := i.complete.Complete(orderProcessed)
	if err != nil {
		return orderProcessed, err
	}

	return orderCompleted, nil
}

func (i *Implementation) Pipeline(ctx context.Context, workerID model.WorkerID, order <-chan model.Order, completedOrder chan<- model.Order) {
	// Процесс обработки заказов
	createCh := i.create.Pipeline(ctx, workerID, order)
	processCh := i.process.Pipeline(ctx, createCh)
	completeCh := i.complete.Pipeline(ctx, processCh)

	for order := range completeCh {
		if order.Err != nil {
			log.Printf("error while processing order: [%d], err: [%v]", order.Order.ID, order.Err)
		} else {
			completedOrder <- order.Order
		}
	}
}

func (i *Implementation) PipelineFan(ctx context.Context, workerID model.WorkerID, order <-chan model.Order, completedOrder chan<- model.Order) {
	createCh := i.create.Pipeline(ctx, workerID, order)

	const limit = 3
	fanOutProgress := make([]<-chan model.PipelineOrder, limit)
	for it := 0; it < limit; it++ {
		fanOutProgress[it] = i.process.Pipeline(ctx, createCh)
	}

	pipeline := i.complete.Pipeline(ctx, fanIn(ctx, fanOutProgress))

	for order := range pipeline {
		if order.Err != nil {
			log.Printf("error while processing order: [%d], err: [%v]", order.Order.ID, order.Err)
		} else {
			completedOrder <- order.Order
		}
	}
}

func fanIn(ctx context.Context, chans []<-chan model.PipelineOrder) <-chan model.PipelineOrder {
	muliteplexed := make(chan model.PipelineOrder)

	var wg sync.WaitGroup
	for _, ch := range chans {
		wg.Add(1)

		go func(ch <-chan model.PipelineOrder) {
			defer wg.Done()
			for v := range ch {
				select {
				case <-ctx.Done():
					return
				case muliteplexed <- v:
				}
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(muliteplexed)
	}()

	return muliteplexed
}
