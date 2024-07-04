package main

import (
	"log"

	"github.com/DBrange/didis-comp-bk/cmd/api/routes"
)

func main() {
	router := routes.NewRouter()

	log.Println("Server listening on 8080")

	if err := router.Run(":8080"); err != nil {
		log.Fatal("Server error:", err)
	}
}
