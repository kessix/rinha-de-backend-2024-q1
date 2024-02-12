package models

import "github.com/kessix/rinha-de-backend-2024-q1/db"

func Update(id int64, transactions Transactions) (int64, error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
}
