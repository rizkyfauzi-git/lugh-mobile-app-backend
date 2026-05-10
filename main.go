package main

import (
	"io"
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
	// Initialize Log File
	logFile, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		mw := io.MultiWriter(os.Stdout, logFile)
		log.SetOutput(mw)
		gin.DefaultWriter = mw
	}

	// Load environment variables
	err = godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found, using system environment variables")
	}

	// Initialize Database
	database.ConnectDB()

	// Initialize Router
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	// Serve Static Files
	r.Static("/static", "./static")
	r.GET("/monitor", func(c *gin.Context) {
		c.File("./static/monitor.html")
	})

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
		protected := api.Group("/")
		protected.Use(middleware.AuthMiddleware())
		{
			protected.GET("/user/profile", controllers.GetProfile)

			// Wallets
			protected.GET("/wallets", controllers.GetWallets)
			protected.POST("/wallets", controllers.CreateWallet)

			// Categories
			protected.GET("/categories", controllers.GetCategories)
			protected.POST("/categories", controllers.CreateCategory)

			// Transactions
			protected.GET("/transactions", controllers.GetTransactions)
			protected.POST("/transactions", controllers.CreateTransaction)
			protected.GET("/transactions/summary", controllers.GetSummary)

			// Monitoring
			protected.GET("/monitor/logs", controllers.GetLogs)
			protected.GET("/monitor/docs", controllers.GetDocs)
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
