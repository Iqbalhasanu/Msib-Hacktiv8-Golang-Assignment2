package order_pg

import (
	"database/sql"
	"time"

	"github.com/Iqbalhasanu/Msib-Hacktiv8-Golang-Assignment2.git/entity"
	"github.com/Iqbalhasanu/Msib-Hacktiv8-Golang-Assignment2.git/pkg/errs"
	"github.com/Iqbalhasanu/Msib-Hacktiv8-Golang-Assignment2.git/repository/order_repository"
)

const (
	createOrderQuery = `
		INSERT INTO "orders"
		(
			ordered_at,
			customer_name
		)
		VALUES($1, $2)
		RETURNING order_id, customer_name, ordered_at, created_at, updated_at
	`
	createItemQuery = `
		INSERT INTO "items"
		(
			item_code,
			description,
			quantity,
			order_id
		)
		VALUES($1, $2, $3, $4)
		RETURNING item_id, item_code, description, quantity, order_id, created_at, updated_at
	`
	getAllOrderQuery = `
		SELECT order_id, customer_name, ordered_at, created_at, updated_at
		FROM "orders";
	`
	getOrderByIdQuery = `
		SELECT order_id, customer_name, ordered_at, created_at, updated_at
		FROM "orders"
		WHERE order_id=$1;
	`
	getItemsByOrderIdQuery = `
		SELECT item_id, item_code, description, quantity, order_id, created_at, updated_at
		FROM "items"
		WHERE order_id=$1;
	`
	updateOrderQuery = `
		UPDATE "orders"
		SET ordered_at = $2,
		customer_name = $3,
		updated_at = $4
		WHERE order_id = $1
		RETURNING order_id, customer_name, ordered_at, created_at, updated_at
	`
	updateItemQuery = `
		UPDATE "items"
		SET description = $3,
		quantity = $4,
		updated_at = $5
		WHERE order_id = $1 AND item_code = $2
		RETURNING item_id, item_code, description, quantity, order_id, created_at, updated_at
	`

	deleteOrderQuery = `
		DELETE FROM "orders"
		WHERE order_id = $1;
	`
)

type orderPg struct {
	db *sql.DB
}

func NewOrderPg(db *sql.DB) order_repository.OrderRepository {
	return &orderPg{db: db}
}

func (o *orderPg) CreateOrder(orderPayload entity.Orders, itemsPayload []entity.Items) (*entity.Orders, errs.MessageErr) {
	tx, err := o.db.Begin()

	if err != nil {
		return nil, errs.InternalServerError("Something Went Wrong")
	}

	orderRow := tx.QueryRow(createOrderQuery, orderPayload.OrderedAt, orderPayload.CustomerName)

	var order entity.Orders

	err = orderRow.Scan(&order.OrderId, &order.CustomerName, &order.OrderedAt, &order.CreatedAt, &order.UpdatedAt)
	if err != nil {
		tx.Rollback()
		return nil, errs.InternalServerError("Something Went Wrong")
	}

	items := []entity.Items{}

	for _, item := range itemsPayload {
		itemRow := tx.QueryRow(createItemQuery, item.ItemCode, item.Description, item.Quantity, order.OrderId)

		var newItem entity.Items

		err = itemRow.Scan(&newItem.ItemId, &newItem.ItemCode, &newItem.Description, &newItem.Quantity, &newItem.OrderId, &newItem.CreatedAt, &newItem.UpdatedAt)
		if err != nil {
			tx.Rollback()
			return nil, errs.InternalServerError("Something Went Wrong")
		}

		items = append(items, newItem)
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return nil, errs.InternalServerError("Something Went Wrong")
	}

	order.Items = items

	return &order, nil
}

func (o *orderPg) GetAllOrder() ([]entity.Orders, errs.MessageErr) {
	rows, err := o.db.Query(getAllOrderQuery)
	if err != nil {
		return nil, errs.InternalServerError("Something Went Wrong")
	}

	defer rows.Close()

	var orders []entity.Orders
	for rows.Next() {
		var order entity.Orders

		err := rows.Scan(&order.OrderId, &order.CustomerName, &order.OrderedAt, &order.CreatedAt, &order.UpdatedAt)
		if err != nil {
			return nil, errs.InternalServerError("Something Went Wrong")
		}

		itemsRows, err := o.db.Query(getItemsByOrderIdQuery, order.OrderId)
		if err != nil {
			return nil, errs.InternalServerError("Something Went Wrong")
		}

		defer itemsRows.Close()

		var items []entity.Items
		for itemsRows.Next() {
			var item entity.Items
			err := itemsRows.Scan(&item.ItemId, &item.ItemCode, &item.Description, &item.Quantity, &item.OrderId, &item.CreatedAt, &item.UpdatedAt)
			if err != nil {
				return nil, errs.InternalServerError("Something Went Wrong")
			}

			items = append(items, item)
		}

		order.Items = append(order.Items, items...)
		orders = append(orders, order)
	}

	return orders, nil
}

func (o *orderPg) GetOrderById(param int) (*entity.Orders, errs.MessageErr) {
	rows := o.db.QueryRow(getOrderByIdQuery, param)

	var order entity.Orders
	err := rows.Scan(&order.OrderId, &order.CustomerName, &order.OrderedAt, &order.CreatedAt, &order.UpdatedAt)
	switch err {
	case sql.ErrNoRows:
		return nil, errs.NotFound("order not found")
	case nil:

	default:
		return nil, errs.InternalServerError("something went wrong")
	}

	itemsRows, err := o.db.Query(getItemsByOrderIdQuery, order.OrderId)
	if err != nil {
		return nil, errs.InternalServerError("Something Went Wrong")
	}

	defer itemsRows.Close()

	var items []entity.Items
	for itemsRows.Next() {
		var item entity.Items
		// item_id, item_code, description, quantity, order_id
		err := itemsRows.Scan(&item.ItemId, &item.ItemCode, &item.Description, &item.Quantity, &item.OrderId, &item.CreatedAt, &item.UpdatedAt)
		if err != nil {
			return nil, errs.InternalServerError("Something Went Wrong")
		}

		items = append(items, item)
	}

	order.Items = append(order.Items, items...)

	return &order, nil
}

func (o *orderPg) UpdateOrder(orderPayload entity.Orders, itemsPayload []entity.Items) (*order_repository.OrderItem, errs.MessageErr) {
	tx, err := o.db.Begin()
	if err != nil {
		return nil, errs.InternalServerError("somethin went wrong")
	}

	row := tx.QueryRow(updateOrderQuery, orderPayload.OrderId, orderPayload.OrderedAt, orderPayload.CustomerName, time.Now().Format(time.RFC3339))

	order := entity.Orders{}

	err = row.Scan(&order.OrderId, &order.CustomerName, &order.OrderedAt, &order.CreatedAt, &order.UpdatedAt)
	if err != nil {
		tx.Rollback()
		return nil, errs.InternalServerError("order scanning")
	}
	// item_id, item_code, description, quantity, order_id, created_at, updated_at
	items := []entity.Items{}
	if len(itemsPayload) != 0 {
		for _, eachItem := range itemsPayload {
			row = tx.QueryRow(updateItemQuery, order.OrderId, eachItem.ItemCode, eachItem.Description, eachItem.Quantity, time.Now().Format(time.RFC3339))

			item := entity.Items{}

			err := row.Scan(&item.ItemId, &item.ItemCode, &item.Description, &item.Quantity, &item.OrderId, &item.CreatedAt, &item.UpdatedAt)
			if err != nil {
				tx.Rollback()
				return nil, errs.InternalServerError("item scanning")
			}

			items = append(items, item)
		}
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return nil, errs.InternalServerError("commit")
	}

	result := order_repository.OrderItem{
		Order: order,
		Items: items,
	}

	return &result, nil
}

func (o *orderPg) DeleteOrder(orderId int) errs.MessageErr {
	res, err := o.db.Exec(deleteOrderQuery, orderId)
	if err != nil {
		return errs.InternalServerError("something went wrong")
	}

	_, err = res.RowsAffected()
	if err != nil {
		return errs.InternalServerError("something went wrong")
	}

	return nil
}
