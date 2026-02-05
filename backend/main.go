package main

import (
	"log"
	"sijaku-hebat/configs"
	"sijaku-hebat/models"
	"sijaku-hebat/routes"
)

func main() {
	db := configs.InitDatabaseConfig()
	if err := db.AutoMigrate(
		&models.Project{},
		&models.Itc{},
		&models.Module{},
	); err != nil {
		log.Fatal("migration failed", err)
	}

	r := routes.InitRouter()
	log.Println("server running on port 8000")
	if err := r.Run(":8000"); err != nil {
		log.Fatal("failed to run server:", err)
	}
}
