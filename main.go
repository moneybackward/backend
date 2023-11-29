package main

import (
	"fmt"
	"os"

	"github.com/gin-contrib/cors"
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

	allowedOrigins := []string{
		"https://moneybackward.jeremiaaxel.my.id",
		"http://localhost:9000",
	}

	engine := gin.Default()
	engine.Use(middlewares.ErrorMiddleware())
	engine.Use(cors.New(cors.Config{
		AllowOrigins:     allowedOrigins,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	docs.SwaggerInfo.BasePath = "/api/v1"

	models.ConnectDB()
	engine.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})
	routes.RegisterRoutes(engine)

	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	engineErr := engine.Run(fmt.Sprintf(":%s", port))
	if engineErr != nil {
		log.Fatal().Msg("Error running engine" + engineErr.Error())
	}
}
