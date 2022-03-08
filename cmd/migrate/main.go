package main

import (
	"github.com/joho/godotenv"
	"openwt.com/boat-app-backend/pkg/database"
	"openwt.com/boat-app-backend/pkg/models"
)

func init() {
	godotenv.Load()
}

func main() {
	db, err := database.GetDatabase()

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.Boat{})
}
