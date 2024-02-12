package models

import (
	"github.com/kessix/rinha-de-backend-2024-q1/db"
)

func CreateTransaction(transactions Transactions) (id int64, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}

	defer conn.Close()

	sql := `INSERT INTO transactions (value, type, description, createdat) VALUES ($1, $2, $3, $4) RETURNING id`

	err = conn.QueryRow(sql, transactions.Value, transactions.Type, transactions.Description, transactions.CreatedAt).Scan(&id)

	return
}

// ID int64 `json: "id"`
// Value int64 `json: "value"`
// Type string `json: "type"`
// Description string `json: "description"`
// CreatedAt int64 `json: "createdat"`
