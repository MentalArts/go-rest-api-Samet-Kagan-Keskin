package database

import (
	"fmt"
	"go-rest-api/internal/models"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

var DB *gorm.DB

func ConnectDatabase() {

	host := getEnv("DB_HOST", "localhost")
	user := getEnv("DB_USER", "postgres")
	password := getEnv("DB_PASSWORD", "postgres")
	dbname := getEnv("DB_NAME", "library")
	port := getEnv("DB_PORT", "5432")

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		host, user, password, dbname, port,
	)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)

	maxRetries := 5
	var err error

	for i := 0; i < maxRetries; i++ {
		log.Printf("Attempting to connect to the database (attempt %d/%d)...", i+1, maxRetries)

		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: newLogger,
		})

		if err != nil {
			log.Printf("Failed to connect to database: %v", err)
			if i < maxRetries-1 {
				retryDelay := time.Duration(2*(i+1)) * time.Second
				log.Printf("Retrying in %v...", retryDelay)
				time.Sleep(retryDelay)
				continue
			}
			log.Fatalf("Could not connect to database after %d attempts", maxRetries)
		} else {
			log.Println("Connected to database successfully!")
			break
		}
	}

	log.Println("Running database migrations...")
	err = DB.AutoMigrate(&models.Author{}, &models.Book{}, &models.Review{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
	log.Println("Database migration completed successfully")
}
