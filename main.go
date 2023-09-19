package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/moneybackward/backend/models"
	"github.com/moneybackward/backend/routes"
)

func main() {
	err := godotenv.Load()
	port, portExists := os.LookupEnv("BACKEND_PORT")
	if !portExists {
		port = "8080"
	}
	mode, modeExists := os.LookupEnv("BACKEND_MODE")
	if !modeExists {
		mode = "dev"
	}
	log.Printf("Running in %s mode", mode)

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	engine := gin.Default()

	db := models.ConnectDB()
	routes.RegisterRoutes(engine, db)

	engine.Run(fmt.Sprintf(":%s", port))
}
