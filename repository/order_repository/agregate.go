package order_repository

import "github.com/Iqbalhasanu/Msib-Hacktiv8-Golang-Assignment2.git/entity"

type OrderItem struct {
	Order entity.Orders
	Items []entity.Items
}
