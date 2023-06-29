package main

import (
	"album/db"
	"album/handler/router"
	"log"
)

func main() {
	albumDB, err := db.NewDB()
	if err != nil {
		log.Fatal(err)
	}
	defer albumDB.Close()

	router := router.NewRouter(albumDB)

	router.Run("localhost:8080")
}
