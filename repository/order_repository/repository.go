package order_repository

import (
	"github.com/Iqbalhasanu/Msib-Hacktiv8-Golang-Assignment2/entity"
	"github.com/Iqbalhasanu/Msib-Hacktiv8-Golang-Assignment2/pkg/errs"
)

type OrderRepository interface {
	CreateOrder(orderPayload entity.Order, itemsPayload []entity.Item) (*entity.Order, errs.MessageErr)
	GetAllOrder() (*[]OrderItem, errs.MessageErr)
	UpdateOrder(orderPayload entity.Order, itemsPayload []entity.Item) (*OrderItem, errs.MessageErr)
	DeleteOrder(orderId int) errs.MessageErr
}
