package model

import (
	"bytes"
	"encoding/json"
	"time"
)

type OrderState string

const (
	OrderStateCreated   = OrderState("created")   // Заказ создан
	OrderStateProcessed = OrderState("processed") // Заказ обработан
	OrderStateComplete  = OrderState("complete")  // Заказ выполнен
)

type OrderID uint64
type ProductID uint64
type WarehouseID uint64
type DeliveryPointID uint64
type WorkerID uint64

type PipelineOrder struct {
	Order Order
	Err   error
}

type Order struct {
	ID              OrderID
	ProductID       ProductID
	WarehouseID     WarehouseID
	DeliveryPointID DeliveryPointID
	WorkerID        WorkerID
	Tracking        []OrderTracking
}

type OrderTracking struct {
	State OrderState
	Start time.Time
}

func (order *Order) OrderToJsonFormat() (string, error) {
	str, err := json.Marshal(order)
	if err != nil {
		return "", err
	}

	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, []byte(str), "", "    "); err != nil {
		return "", err
	}
	return prettyJSON.String(), nil
}
