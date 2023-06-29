package db_test

import (
	"album/db"
	"log"
	"testing"
)

func TestNewDB(t *testing.T) {
	db, err := db.NewDB()
	if err != nil {
		log.Fatal("cannot connect to db")
	}
	defer db.Close()
}
