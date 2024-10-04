package main

import (
	"log"
	"modulux/api/routes"
	"modulux/config"
	"modulux/database"
)

func main() {

	cfg := config.LoadConfig()
	database.Connect(cfg)

	log.Println("Successfully connected to the database")

	r := routes.SetupRouter()
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to run server:", err)
	}
}
