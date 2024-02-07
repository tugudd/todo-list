package database

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
)

var Conn *pgx.Conn = DBinstance()

func DBinstance() *pgx.Conn {
	dbURL := "postgres://tugudd:password@localhost:5432/todo"
	conn, err := pgx.Connect(context.Background(), dbURL)

	if err != nil {
		log.Fatal(err.Error())
	}

	return conn
}
