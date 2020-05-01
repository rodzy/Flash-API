package main

import (
	"log"

	"github.com/rodzy/flash/db"
	"github.com/rodzy/flash/handlers"
)

func main() {
	if db.CheckConnection() == 0 {
		log.Fatal("Connection down")
		return
	}
	handlers.DirectDrivers()
}
