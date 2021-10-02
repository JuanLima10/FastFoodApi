package main

import (
	"go/database"
	"go/server"
)

func main() {
	database.StartDB()

	server := server.NewServer()
	server.Run()
}
