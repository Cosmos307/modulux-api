package database

import (
	"context"
	"fmt"
	"log"
	"modulux/config"

	"github.com/jackc/pgx/v5"
)

var DB *pgx.Conn

// Connect initializes the database connection
func Connect(cfg *config.Config) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:5432/%s",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBName)
	var err error
	DB, err = pgx.Connect(context.Background(), dsn)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Verify the connection
	err = DB.Ping(context.Background())
	if err != nil {
		log.Fatal("Failed to ping database:", err)
	}
}
