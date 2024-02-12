package models

import (
	"github.com/kessix/rinha-de-backend-2024-q1/db"
)

func Get(id int64) (transactions Transactions, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	row := conn.QueryRow(`SELECT * FROM transcations WHERE id=$1`, id)

	err = row.Scan(&transactions.ID, &transactions.Value, &transactions.Type, &transactions.Description, &transactions.CreatedAt)

	return
}
