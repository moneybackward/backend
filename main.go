package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	docs "github.com/moneybackward/backend/docs"
	"github.com/moneybackward/backend/middlewares"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/moneybackward/backend/models"
	"github.com/moneybackward/backend/routes"
)

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
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
	engine.Use(middlewares.ErrorMiddleware())

	docs.SwaggerInfo.BasePath = "/api/v1"

	models.ConnectDB()
	routes.RegisterRoutes(engine)

	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	engineErr := engine.Run(fmt.Sprintf(":%s", port))
	if engineErr != nil {
		log.Fatal(engineErr)
	}
}
