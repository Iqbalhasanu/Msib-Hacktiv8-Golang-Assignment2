package order_repository

import (
	"github.com/Iqbalhasanu/Msib-Hacktiv8-Golang-Assignment2.git/entity"
	"github.com/Iqbalhasanu/Msib-Hacktiv8-Golang-Assignment2.git/pkg/errs"
)

type OrderRepository interface {
	CreateOrder(orderPayload entity.Orders, itemsPayload []entity.Items) (*entity.Orders, errs.MessageErr)
	GetAllOrder() ([]entity.Orders, errs.MessageErr)
	GetOrderById(int) (*entity.Orders, errs.MessageErr)
	UpdateOrder(entity.Orders, []entity.Items) (*OrderItem, errs.MessageErr)
	DeleteOrder(int) errs.MessageErr
}
