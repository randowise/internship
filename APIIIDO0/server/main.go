package main

import (
	"example/APIIIDO0/database"
	"example/APIIIDO0/server"
)

func main() {
	database.StartDB()

	server := server.NewServer()

	server.Run()
}
