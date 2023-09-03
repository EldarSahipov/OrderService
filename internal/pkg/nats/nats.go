package nats

import (
	"OrderService/internal/models"
	"encoding/json"
)

func UnmarshalTheMessage(str string) (*models.Order, error) {
	var order models.Order
	err := json.Unmarshal([]byte(str), &order)
	if err != nil {
		return nil, err
	}

	for i := range order.Items {
		order.Items[i].OrderUid = order.OrderUid
	}

	return &order, nil
}
