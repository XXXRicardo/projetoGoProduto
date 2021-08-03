package main

import (
	"api/database"
	"api/server"
)

func main() {

	database.StartDB()

	server := server.NewServer()

	server.Run()

}
