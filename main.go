package main

import (
	"testgo/config"
	"testgo/models"
	"testgo/routes"
)

func main() {

	db := config.SetupDB()
	db.AutoMigrate(&models.Animal{})

	r := routes.SetupRoutes(db)
	r.Run()
}
