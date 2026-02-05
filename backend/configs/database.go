package configs

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDatabaseConfig() *gorm.DB {
	databaseUrl := "postgres://postgres:rayzen7@100.74.236.13:5432/sijaku_hebat?sslmode=disable"
	if databaseUrl == "" {
		log.Fatal("Database is empty")
	}

	db, err := gorm.Open(postgres.Open(databaseUrl), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	log.Println("Database connected successfully")
	return db
}
