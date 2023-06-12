package entity

import (
	"time"

	"github.com/Iqbalhasanu/Msib-Hacktiv8-Golang-Assignment2.git/dto"
)

type Items struct {
	ItemId      int
	ItemCode    string
	Description string
	Quantity    int
	OrderId     int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (i *Items) ItemToItemResponse() dto.ItemResponse {
	return dto.ItemResponse{
		ItemId:      i.ItemId,
		ItemCode:    i.ItemCode,
		Description: i.Description,
		Quantity:    i.Quantity,
		OrderId:     i.OrderId,
		CreatedAt:   i.CreatedAt,
		UpdatedAt:   i.UpdatedAt,
	}
}
