package printer

import (
	"fmt"
	"log"
	"sync"

	"homework/internal/model"
)

func PrintCompletedOrder(wg *sync.WaitGroup, ordersCh <-chan model.Order) {
	go func() {
		defer wg.Done()
		for order := range ordersCh {
			orderStringInJsonFormat, err := order.OrderToJsonFormat()
			if err != nil {
				log.Printf("error while converting order[%d] to json format, err: [%v]", order.ID, err)
			} else {
				fmt.Println(orderStringInJsonFormat)
			}
		}
	}()
}
