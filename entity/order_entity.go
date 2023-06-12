package entity

import "time"

type Orders struct {
	OrderId      int
	CustomerName string
	OrderedAt    time.Time
	Items        []Items
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
