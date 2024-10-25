package main

import (
	"log"
	"modulux/api/middleware"
	"modulux/api/routes"
	"modulux/config"
	"modulux/database"
)

func main() {

	initialize()

	r := routes.SetupRouter()

	err := r.Run(":8080")
	if err != nil {
		log.Fatal("Failed to run server:", err)
	}

}

func initialize() {
	cfg := config.LoadConfig()

	database.Connect(cfg)
	log.Println("Successfully connected to the database")

	middleware.InitializeJWT(cfg.JWTSecret)
}
