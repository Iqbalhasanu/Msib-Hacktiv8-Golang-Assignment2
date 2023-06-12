package dto

import (
	"time"
)

type NewOrderRequest struct {
	OrderedAt    time.Time     `json:"orderedAt"`
	CustomerName string        `json:"customerName"`
	Items        []ItemRequest `json:"items"`
}

func (o *NewOrderRequest) ItemsToItemCode() []string {
	itemCodes := []string{}

	for _, value := range o.Items {
		itemCodes = append(itemCodes, value.ItemCode)
	}

	return itemCodes
}

type NewOrderResponse struct {
	Message    string          `json:"message"`
	Data       NewOrderRequest `json:"data"`
	StatusCode int             `json:"statusCode"`
}

// type AllOrderResponse struct {
// 	Data []struct {
// 	} `json:"Data"`
// }

type GetOrderResponse struct {
	Code int           `json:"code"`
	Data OrderResponse `json:"data"`
}

type OrderResponse struct {
	Id           int            `json:"id"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	CustomerName string         `json:"customerName"`
	Items        []ItemResponse `json:"Item"`
}
