package main

import (
	"github.com/andreprisya/task-5-vix-btpns-AndrePrisyaLubis/database"
	"github.com/andreprisya/task-5-vix-btpns-AndrePrisyaLubis/models"
	"github.com/andreprisya/task-5-vix-btpns-AndrePrisyaLubis/router"
)

func main() {
	db := database.SetupDB()
	db.AutoMigrate(&models.User{})

	r := router.SetupRoutes(db)
	r.Run()
}
