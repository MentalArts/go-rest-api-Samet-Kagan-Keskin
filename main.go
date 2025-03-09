package main

import (
	"log"
	"os"

	// Import the docs package for Swagger
	_ "go-rest-api/docs"
	"go-rest-api/internal/database"
	"go-rest-api/internal/handlers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Book Library Management API
// @version 1.0
// @description A RESTful API for managing a book library system
// @termsOfService http://example.com/terms/

// @contact.name API Support
// @contact.url http://example.com/support
// @contact.email support@example.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
// @schemes http
func main() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found or error loading: %v", err)
	}

	// Connect to database
	database.ConnectDatabase()

	// Set Gin to release mode in production
	if os.Getenv("GIN_MODE") == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Initialize router
	r := gin.Default()

	// CORS middleware
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	// Health check endpoint
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "UP",
		})
	})

	// API v1 routes
	v1 := r.Group("/api/v1")
	{
		// Book routes
		books := v1.Group("/books")
		{
			books.GET("", handlers.GetBooks)
			books.GET("/:id", handlers.GetBook)
			books.POST("", handlers.CreateBook)
			books.PUT("/:id", handlers.UpdateBook)
			books.DELETE("/:id", handlers.DeleteBook)

			// Review routes related to books
			books.GET("/:id/reviews", handlers.GetBookReviews)
			books.POST("/:id/reviews", handlers.AddReview)
		}

		// Author routes
		authors := v1.Group("/authors")
		{
			authors.GET("", handlers.GetAuthors)
			authors.GET("/:id", handlers.GetAuthor)
			authors.POST("", handlers.CreateAuthor)
			authors.PUT("/:id", handlers.UpdateAuthor)
			authors.DELETE("/:id", handlers.DeleteAuthor)
		}

		// Review routes (for update and delete)
		reviews := v1.Group("/reviews")
		{
			reviews.PUT("/:id", handlers.UpdateReview)
			reviews.DELETE("/:id", handlers.DeleteReview)
		}
	}

	// Swagger documentation route
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Get port from environment variable or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server started on :%s", port)
	log.Printf("Swagger UI available at http://localhost:%s/swagger/index.html", port)

	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
