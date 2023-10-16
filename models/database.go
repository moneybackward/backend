package models

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func getDbConfig() map[string]string {
	result := make(map[string]string)
	defaults := map[string]string{
		"DB_HOST":     "localhost",
		"DB_NAME":     "moneybackward",
		"DB_USER":     "postgres",
		"DB_PASSWORD": "postgres",
		"DB_PORT":     "5432",
	}

	for key, value := range defaults {
		env, exists := os.LookupEnv(key)
		if exists {
			result[key] = env
		} else {
			log.Default().Printf("Environment variable %s not set. Defaulting to %s", key, value)
			result[key] = value
		}
	}

	return result
}

func ConnectDB() *gorm.DB {
	dbConfigs := getDbConfig()

	dbConnection := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Bangkok",
		dbConfigs["DB_HOST"],
		dbConfigs["DB_USER"],
		dbConfigs["DB_PASSWORD"],
		dbConfigs["DB_NAME"],
		dbConfigs["DB_PORT"],
	)
	postgresConfig := postgres.Config{
		DSN:                  dbConnection,
		PreferSimpleProtocol: true,
	}

	db, err := gorm.Open(postgres.New(postgresConfig), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	err = db.AutoMigrate(
		&User{},
		&Note{},
		&Category{},
		&Transaction{},
	)
	if err != nil {
		panic("failed to migrate database")
	}

	DB = db
	return DB
}
