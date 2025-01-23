package database

import (
	"gmxBackend/models"
	"log"
)

func InitTable() {
	ConnectDB()
	db := GetDB()
	err := db.AutoMigrate(
		models.Order{},
	)
	if err != nil {
		log.Fatalf("Failed to auto migrate: %v", err)
	}
	log.Println("Database tables created successfully", db)
}
