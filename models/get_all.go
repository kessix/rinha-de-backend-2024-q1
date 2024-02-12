package models

import (
	"github.com/kessix/rinha-de-backend-2024-q1/db"
)

func GetAll() (transactions []Transactions, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	rows, err := conn.Query(`SELECT * FROM transcations`)

	if err != nil {
		return
	}

	for rows.Next() {
		var transaction Transactions
		err = rows.Scan(&transactions.ID, &transactions.Value, &transactions.Type, &transactions.Description, &transactions.CreatedAt)
		if err != nil {
			continue
		}

		transactions = append(transactions, transaction)
	}

	return
}
