package order_pg

import (
	"database/sql"
	"fmt"

	"github.com/Iqbalhasanu/Msib-Hacktiv8-Golang-Assignment2/entity"
	"github.com/Iqbalhasanu/Msib-Hacktiv8-Golang-Assignment2/pkg/errs"
	"github.com/Iqbalhasanu/Msib-Hacktiv8-Golang-Assignment2/repository/order_repository"
)

const (
	createOrderQuery = `
		INSERT INTO "orders"
			(
				ordered_at,
				customer_name
			)
		VALUES($1, $2)
		RETURNING order_id, customer_name, ordered_at, created_at,updated_at
	`
	createItemQuery = `
		INSERT INTO "items"
			(
				item_code,
				quantity,
				description,
				order_id
			)
		VALUES($1, $2, $3, $4)
		RETURNING item_id
`
)

type orderPg struct {
	db *sql.DB
}

func NewOrderPG(db *sql.DB) order_repository.OrderRepository {
	return &orderPg{db: db}
}

func (o *orderPg) CreateOrder(orderPayload entity.Order, itemsPayload []entity.Item) (*entity.Order, errs.MessageErr) {
	tx, err := o.db.Begin()

	if err != nil {
		return nil, errs.NewInternalServerError("Something Went Wrorng")
	}

	orderRow := tx.QueryRow(createOrderQuery, orderPayload.OrderedAt, orderPayload.CustomerName)

	var order entity.Order

	err = orderRow.Scan(&order.OrderId, &order.CustomerName, &order.OrderedAt, &order.CreatedAt, &order.UpdatedAt)

	if err != nil {
		tx.Rollback()
		return nil, errs.NewInternalServerError("Something Went Wrorng")
	}

	for _, eachItem := range itemsPayload {
		itemRow := tx.QueryRow(createItemQuery, eachItem.ItemCode, eachItem.Quantity, eachItem.Description, order.OrderId)

		var itemId int

		err = itemRow.Scan(&itemId)

		if err != nil {
			tx.Rollback()
			return nil, errs.NewInternalServerError("Something Went Wrorng")
		}

	}

	err = tx.Commit()

	if err != nil {
		tx.Rollback()
		return nil, errs.NewInternalServerError("Something Went Wrorng")
	}

	return &order, nil

}

func (o *orderPg) GetAllOrder() (*[]order_repository.OrderItem, errs.MessageErr) {

	queryGetAllOrder := `SELECT order_id, customer_name, ordered_at, created_at, updated_at FROM "orders"`

	rows, err := o.db.Query(queryGetAllOrder)

	if err != nil {
		return nil, errs.NewInternalServerError("Something Went Wrorng dek")
	}

	defer rows.Close()

	var orders []order_repository.OrderItem

	for rows.Next() {
		var order order_repository.OrderItem

		err := rows.Scan(&order.OrderId, &order.CustomerName, &order.OrderedAt, &order.CreatedAt, &order.UpdatedAt)

		if err != nil {
			return nil, errs.NewInternalServerError("Something Went Wrorng mas")
		}

		queryGetItemByOrderId := `
		SELECT item_id, item_code, quantity, description, order_id, created_at, updated_at FROM "items"
		WHERE order_id = $1
		`

		itemrows, err := o.db.Query(queryGetItemByOrderId, order.OrderId)

		if err != nil {
			return nil, errs.NewInternalServerError("Something Went Wrorng mbak")
		}

		var items []entity.Item

		for itemrows.Next() {

			item := entity.Item{}

			err := itemrows.Scan(&item.ItemId, &item.ItemCode, &item.Quantity, &item.Description, &item.OrderId, &item.CreatedAt, &item.UpdatedAt)

			if err != nil {

				return nil, errs.NewInternalServerError("Something Went Wrorng asdasdasd")
			}

			items = append(items, item)

		}
		order.Items = append(order.Items, items...)

		orders = append(orders, order)

	}
	return &orders, nil
}

func (o *orderPg) UpdateOrder(orderPayload entity.Order, itemsPayload []entity.Item) (*order_repository.OrderItem, errs.MessageErr) {

	updateItemQuery := `
		UPDATE "items"
		SET description = $2,
		quantity = $3
		WHERE item_code = $1
		RETURNING item_id, item_code, quantity, description, updated_at, order_id, created_at
	`

	updateOrderQuery := `
		UPDATE "orders"
		SET ordered_at = $2,
		customer_name = $3
		WHERE order_id = $1
		RETURNING order_id, customer_name, created_at, updated_at
	`

	tx, err := o.db.Begin()

	if err != nil {

		return nil, errs.NewInternalServerError("Internal Server Error")
	}

	row := tx.QueryRow(updateOrderQuery, orderPayload.OrderId, orderPayload.OrderedAt, orderPayload.CustomerName)

	order := entity.Order{}

	err = row.Scan(&order.OrderId, &order.CustomerName, &order.CreatedAt, &order.UpdatedAt)

	if err != nil {
		tx.Rollback()
		return nil, errs.NewInternalServerError("Internal Server Error")
	}

	items := []entity.Item{}
	for _, eachItem := range itemsPayload {
		row = tx.QueryRow(updateItemQuery, eachItem.ItemCode, eachItem.Description, eachItem.Quantity)
		item := entity.Item{}
		err = row.Scan(&item.ItemId, &item.ItemCode, &item.Quantity, &item.Description, &item.UpdatedAt, &item.OrderId, &item.CreatedAt)

		if err != nil {
			tx.Rollback()
			return nil, errs.NewInternalServerError("Internal Server Error")
		}

		items = append(items, item)
	}

	err = tx.Commit()

	if err != nil {
		tx.Rollback()
		return nil, errs.NewInternalServerError("Internal Server Error")
	}

	result := order_repository.OrderItem{
		OrderId:      order.OrderId,
		CustomerName: order.CustomerName,
		OrderedAt:    order.OrderedAt,
		CreatedAt:    order.CreatedAt,
		UpdatedAt:    order.UpdatedAt,
		Items:        items,
	}

	return &result, nil
}

func (o *orderPg) DeleteOrder(orderId int) errs.MessageErr {

	checkOrderQuery := `
	SELECT * FROM "orders"
	WHERE order_id = $1
	`

	row := o.db.QueryRow(checkOrderQuery, orderId)

	var order entity.Order
	if err := row.Scan(&order.OrderId, &order.CustomerName, &order.OrderedAt, &order.CreatedAt, &order.UpdatedAt); err != nil {
		fmt.Println(err)
		return errs.NewNotFoundError("Order Not Found")
	}

	queryDelete := `
	DELETE FROM "orders" 
	WHERE order_id = $1
	`

	if _, err := o.db.Exec(queryDelete, orderId); err != nil {
		return errs.NewInternalServerError("Internal Server Error2")
	}

	return nil
}
