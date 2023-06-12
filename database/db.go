package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var (
	db  *sql.DB
	err error
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbName   = "assignment-2"
)

func Initialize() {
	createConnection()
	createTable()
}

func createConnection() {
	psqlcon := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)

	db, err = sql.Open("postgres", psqlcon)
	if err != nil {
		log.Panic("error occured while trying to validate database arguments:", err)
	}

	err = db.Ping()
	if err != nil {
		log.Panic("error occured while trying to connect to database:", err)
	}

	fmt.Println("Connected")
}

func createTable() {
	orderTable := `
		CREATE TABLE IF NOT EXISTS "orders" (
			order_id SERIAL PRIMARY KEY,
			customer_name VARCHAR(255) NOT NULL,
			ordered_at timestamptz DEFAULT now(),
			created_at timestamptz DEFAULT now(),
			updated_at timestamptz DEFAULT now()
		);
	`

	itemTable := `
		CREATE TABLE IF NOT EXISTS "items" (
			item_id SERIAL PRIMARY KEY,
			item_code VARCHAR(191) NOT NULL,
			quantity int NOT NULL,
			description TEXT NOT NULL,
			order_id int NOT NULL,
			created_at timestamptz DEFAULT now(),
			updated_at timestamptz DEFAULT now(),
			CONSTRAINT items_order_id_fk
				FOREIGN KEY(order_id)
					REFERENCES orders(order_id)
						ON DELETE CASCADE
		);
	`

	_, err = db.Exec(orderTable)
	if err != nil {
		log.Panic("error occured while trying create orders table:", err)
	}

	_, err = db.Exec(itemTable)
	if err != nil {
		log.Panic("error occured while trying create items table:", err)
	}
}

func GetDatabaseInstance() *sql.DB {
	return db
}
