package main

import (
	"log"
	"os"
	"time"
	"warteg-system-backend/controllers"
	"warteg-system-backend/database"
	"warteg-system-backend/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found, using system environment variables")
	}

	// Initialize Database
	database.ConnectDB()

	// Initialize Router
	r := gin.Default()

	// Global Middlewares
	r.Use(middleware.SecurityHeaders())
	
	// CORS configuration
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization", "X-Requested-With"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Rate Limiting
	limiter := middleware.NewIPRateLimiter(1, 5) // 1 request per second with burst of 5

	// Public Routes
	api := r.Group("/api")
	{
		auth := api.Group("/auth")
		auth.Use(middleware.RateLimitMiddleware(limiter))
		{
			auth.POST("/login", controllers.Login)
			auth.POST("/register", controllers.Register)
		}

		// Protected Routes
		protected := api.Group("/user")
		protected.Use(middleware.AuthMiddleware())
		{
			protected.GET("/profile", controllers.GetProfile)
		}
	}

	// Start server
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	r.Run(":" + port)
}
