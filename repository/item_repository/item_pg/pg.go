package item_pg

import (
	"database/sql"
	"fmt"

	"github.com/Iqbalhasanu/Msib-Hacktiv8-Golang-Assignment2.git/entity"
	"github.com/Iqbalhasanu/Msib-Hacktiv8-Golang-Assignment2.git/pkg/errs"
	"github.com/Iqbalhasanu/Msib-Hacktiv8-Golang-Assignment2.git/item_repository"
)

type itemPg struct {
	db *sql.DB
}

func NewItemPg(db *sql.DB) item_repository.ItemRepository {
	return &itemPg{
		db: db,
	}
}

func (i *itemPg) generateQuery(length int) string {
	query := "("

	for i := 1; i <= length; i++ {
		if i == length {
			query += fmt.Sprintf("$%d)", i+1)
			continue
		}
		query += fmt.Sprintf("$%d,", i+1)
	}

	return query
}

func (i *itemPg) findItemsByItemCodeQuery(length int, orderId int) string {
	query := `
	SELECT item_id, item_code, quantity, description, order_id, created_at, updated_at
	FROM "items"
	WHERE order_id=$1 AND item_code IN 
	`

	param := i.generateQuery(length)

	return query + param
}

func (i *itemPg) FindItemsByItemCode(itemCodes []string, orderId int) ([]*entity.Items, errs.MessageErr) {
	query := i.findItemsByItemCodeQuery(len(itemCodes), orderId)

	args := []any{orderId}
	for _, value := range itemCodes {
		args = append(args, value)
	}

	rows, err := i.db.Query(query, args...)
	if err != nil {
		return nil, errs.InternalServerError("something went wrong")
	}

	defer rows.Close()

	items := []*entity.Items{}
	for rows.Next() {
		item := entity.Items{}

		err := rows.Scan(&item.ItemId, &item.ItemCode, &item.Quantity, &item.Description, &item.OrderId, &item.CreatedAt, &item.UpdatedAt)
		if err != nil {
			return nil, errs.InternalServerError("something went wrong")
		}

		items = append(items, &item)
	}

	return items, nil
}
