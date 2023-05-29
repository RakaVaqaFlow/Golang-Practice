package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	"homework/internal/model"
	"homework/internal/pkg/gateway"
	completestep "homework/internal/pkg/gateway/steps/complete"
	createstep "homework/internal/pkg/gateway/steps/create"
	processstep "homework/internal/pkg/gateway/steps/process"
	"homework/internal/pkg/printer"
	"homework/internal/pkg/producer"
)

func main() {

	// Создаем группу ожидания для горутины вывода завершенных заказов
	var wg sync.WaitGroup

	// Создаем канал, cодержащий структуры заказов с инициализированным полем ID товара
	orders := producer.Orders()

	// Создаем шаги шлюза
	create := createstep.New()
	process := processstep.New()
	complete := completestep.New()

	// Создаем канал для заверешенных заказов
	ordersCompleted := make(chan model.Order)

	// Запускаем горутину для вывода завершенных заказов
	wg.Add(1)
	printer.PrintCompletedOrder(&wg, ordersCompleted)

	// Создаем сервер шлюза
	server := gateway.New(create, process, complete)

	start := time.Now().UTC()
	for order := range orders {
		// Обратотка заказа (создание, обработка, завершение)
		orderCompleted, err := server.Process(order, 0)
		if err != nil {
			log.Printf("error while processing order[%d], err: [%v]", orderCompleted.ID, err)
		} else {
			ordersCompleted <- orderCompleted
		}
	}
	close(ordersCompleted)
	wg.Wait()
	fmt.Printf("Total duration: %f", time.Since(start).Seconds())

}
