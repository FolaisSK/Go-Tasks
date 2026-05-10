package main

import (
	"TaskManager/config"
	"TaskManager/db"
	"TaskManager/middleware"
	"TaskManager/routes"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	cfg := config.Load()

	db.Connect(cfg)

	r := gin.Default()
	r.Use(middleware.Logger())
	routes.SetupRoutes(r)
	r.Run(":8080")
}
