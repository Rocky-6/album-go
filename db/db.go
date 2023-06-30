package db

import (
	"database/sql"
	_ "embed"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

//go:embed schema.sql
var schema string

func NewDB() (*sql.DB, error) {
	var (
		host     string = os.Getenv("POSTGRES_HOSTNAME")
		user     string = os.Getenv("POSTGRES_USER")
		password string = os.Getenv("POSTGRES_PASSWORD")
		dbname   string = os.Getenv("POSTGRES_DB")
	)

	fmt.Println(schema)

	psqlconn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", host, user, password, dbname)

	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	fmt.Println("Connected!")

	if _, err = db.Exec(schema); err != nil {
		return nil, err
	}

	return db, nil
}
