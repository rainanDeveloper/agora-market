package main

import (
	"log"

	"github.com/rainanDeveloper/agora-market/api"
)

func main() {
	server := api.Initialize();

	if err := server.Run(":8080"); err!= nil {
		log.Fatal(err)
	}
}