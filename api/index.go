package handler

import (
	"net/http"
	"time"
	"warteg-system-backend/controllers"
	"warteg-system-backend/database"
	"warteg-system-backend/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var app *gin.Engine

func init() {
	_ = godotenv.Load()
	database.ConnectDB()

	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	// Control Panel Routes
	r.GET("/", func(c *gin.Context) {
		c.File("./static/index.html")
	})
	r.GET("/login", func(c *gin.Context) {
		c.File("./static/login.html")
	})

	r.Use(gin.Recovery())
	r.Use(middleware.SecurityHeaders())
	
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization", "X-Requested-With"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	limiter := middleware.NewIPRateLimiter(1, 5)

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
		}
	}

	app = r
}

func Handler(w http.ResponseWriter, r *http.Request) {
	if database.DB == nil {
		// Try to reconnect if DB is nil (serverless cold start might have failed init)
		database.ConnectDB()
	}

	if database.DB == nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Write([]byte("Database connection is not available. Check your environment variables and DB host access."))
		return
	}

	app.ServeHTTP(w, r)
}
