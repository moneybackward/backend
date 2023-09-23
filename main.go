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
	dotenvErr := godotenv.Load()
	if dotenvErr != nil {
		log.Fatal("Error loading .env file")
	}
	port, portExists := os.LookupEnv("BACKEND_PORT")
	if !portExists {
		port = "8080"
	}
	mode, modeExists := os.LookupEnv("BACKEND_MODE")
	if !modeExists {
		mode = "dev"
	}
	log.Printf("Running in %s mode", mode)

	engine := gin.Default()

	models.ConnectDB()
	routes.RegisterRoutes(engine)

	engineErr := engine.Run(fmt.Sprintf(":%s", port))
	if engineErr != nil {
		log.Fatal(engineErr)
	}
}
