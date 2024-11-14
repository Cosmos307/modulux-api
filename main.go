package main

import (
	"log"
	"modulux/api/controllers"
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
	err = r.Run(":6050")
	if err != nil {
		log.Fatal("Failed to run 6050 server:", err)
	}

}

func initialize() {
	cfg := config.LoadConfig()

	database.Connect(cfg)
	log.Println("Successfully connected to the database")

	controllers.InitializeCrossRef(cfg.CrossRefURL)
	middleware.InitializeJWT(cfg.JWTSecret)
}
