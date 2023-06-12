package order_repository

import (
	"time"

	"github.com/Iqbalhasanu/Msib-Hacktiv8-Golang-Assignment2/entity"
)

type OrderItem struct {
	OrderId      int
	CustomerName string
	OrderedAt    time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Items        []entity.Item
}

// func (o *OrderItem) OrderItemToOrderResponse() dto.AllOrderResponse {
// 	orderResponse := dto.AllOrderResponse{}

// 	for _, eachOrder := range o. {

// 		items := dto.ItemResponse{}
// 		for _, eachItem := range eachOrder {
// 			items.ItemCode = eachItem.ItemCode
// 			items.Quantity = eachItem.Quantity
// 			items.Description = eachItem.Description
// 		}

// 		order := dto.OrderResponse{
// 			Id:           o.OrderId,
// 			CreatedAt:    o.CreatedAt,
// 			UpdatedAt:    o.UpdatedAt,
// 			CustomerName: o.CustomerName,
// 			Items:        []dto.ItemResponse{items},
// 		}

// 	}

// 	return orderResponse
// }
