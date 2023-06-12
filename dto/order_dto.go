package dto

import "time"

type OrderResponse struct {
	OrderId      int            `json:"id"`
	CreatedAt    time.Time      `json:"createdAt"`
	UpdatedAt    time.Time      `json:"updatedAt"`
	OrderedAt    time.Time      `json:"orderedAt"`
	CustomerName string         `json:"customerName"`
	Items        []ItemResponse `json:"items"`
}

type NewOrderRequest struct {
	OrderedAt    time.Time        `json:"orderedAt" binding:"required"`
	CustomerName string           `json:"customerName" binding:"required"`
	Items        []NewItemRequest `json:"items" binding:"required"`
}

func (n *NewOrderRequest) ItemsToItemCode() []string {
	itemCodes := []string{}

	if len(n.Items) == 0 {
		return nil
	}

	for _, eachItem := range n.Items {
		itemCodes = append(itemCodes, eachItem.ItemCode)
	}

	return itemCodes
}

type NewOrderResponse struct {
	Message    string        `json:"message"`
	Data       OrderResponse `json:"data"`
	StatusCode int           `json:"code"`
}

type AllOrderResponse struct {
	Message    string          `json:"message"`
	Data       []OrderResponse `json:"data"`
	StatusCode int             `json:"code"`
}

type DeleteOrderResponse struct {
	Message    string `json:"message"`
	StatusCode int    `json:"code"`
}
