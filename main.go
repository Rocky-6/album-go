package main

import (
	"album/controller/router"
	"album/db"
	"log"
)

func main() {
	albumDB, err := db.NewDB()
	if err != nil {
		log.Fatal(err)
	}
	defer albumDB.Close()

	router := router.NewRouter(albumDB)
	router.Run(":8080")
}
