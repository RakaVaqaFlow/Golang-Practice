package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"homework/internal/config"
	"homework/internal/model"
	"homework/internal/pkg/gateway"
	completestep "homework/internal/pkg/gateway/steps/complete"
	createstep "homework/internal/pkg/gateway/steps/create"
	processstep "homework/internal/pkg/gateway/steps/process"
	"homework/internal/pkg/printer"
	"homework/internal/pkg/producer"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var workerWaitGroup sync.WaitGroup
	var printerWaitGroup sync.WaitGroup
	// Создаем канал, cодержащий структуры заказов с инициализированным полем ID товара
	orders := producer.Orders()

	// Создаем шаги шлюза
	create := createstep.New()
	process := processstep.New()
	complete := completestep.New()

	// Создаем канал для заверешенных заказов
	ordersCompleted := make(chan model.Order)
	// Запускаем горутину для вывода завершенных заказов
	printerWaitGroup.Add(1)
	printer.PrintCompletedOrder(&printerWaitGroup, ordersCompleted)

	// Создаем сервер шлюза
	server := gateway.New(create, process, complete)

	// Запуск воркеров
	start := time.Now().UTC()

	for i := 0; i < config.NumberOfWorkers; i++ {
		workerWaitGroup.Add(1)
		go worker(ctx, &workerWaitGroup, server, model.WorkerID(uint64(i)), orders, ordersCompleted)
	}
	workerWaitGroup.Wait()
	close(ordersCompleted)
	printerWaitGroup.Wait()
	fmt.Printf("Total duration: %f", time.Since(start).Seconds())
}

func worker(ctx context.Context, wg *sync.WaitGroup, server *gateway.Implementation, workerID model.WorkerID, orders <-chan model.Order, ordersCompleted chan<- model.Order) {
	defer wg.Done()
	server.PipelineFan(ctx, workerID, orders, ordersCompleted)
}
