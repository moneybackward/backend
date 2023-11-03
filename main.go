package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	docs "github.com/moneybackward/backend/docs"
	"github.com/moneybackward/backend/middlewares"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
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
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	if dotenvErr != nil {
		log.Fatal().Msg("Error loading .env file")
	}
	port, portExists := os.LookupEnv("BACKEND_PORT")
	if !portExists {
		port = "8080"
	}
	mode, modeExists := os.LookupEnv("BACKEND_MODE")
	if !modeExists {
		mode = "dev"
	}
	log.Info().Msg("Running in " + mode + " mode")

	engine := gin.Default()
	engine.Use(middlewares.ErrorMiddleware())

	docs.SwaggerInfo.BasePath = "/api/v1"

	models.ConnectDB()
	routes.RegisterRoutes(engine)

	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	engineErr := engine.Run(fmt.Sprintf(":%s", port))
	if engineErr != nil {
		log.Fatal().Msg("Error running engine" + engineErr.Error())
	}
}
