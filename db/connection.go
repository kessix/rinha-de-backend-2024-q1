package db

import (
	"database/sql"
	"fmt"

	"github.com/kessix/rinha-de-backend-2024-q1/configs"
	_ "github.com/lib/pq"
)

func OpenConnection() {
	conf := configs.GetDB(*sql.DB, error)

	sc := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		conf.Host, conf.Port, conf.User, conf.Pass, conf.Database)

	conn, err := sql.Open("postgres", sc)
	if err != nil {
		panic(err)
	}

	err = conn.Ping()

	return conn, err
}
