packages models

import (
	"github.com/kessix/rinha-de-backend-2024-q1/db"	
)

func CreateTransaction(transactions Transactions) (id int64, err error) {
	conn, err := db.OpenConnection()
	
	return
}