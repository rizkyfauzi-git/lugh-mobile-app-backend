package handler

import (
	"net/http"
	"warteg-system-backend/controllers"
	"warteg-system-backend/database"
	"warteg-system-backend/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"time"
)

var app *gin.Engine

func init() {
	// In Vercel, we don't strictly need .env because we set them in the dashboard
	_ = godotenv.Load()

	// Initialize Database
	database.ConnectDB()

	// Initialize Router
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Recovery())

	// Middlewares
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

		protected := api.Group("/user")
		protected.Use(middleware.AuthMiddleware())
		{
			protected.GET("/profile", controllers.GetProfile)
		}
	}

	app = r
}

// Handler is the entry point for Vercel serverless function
func Handler(w http.ResponseWriter, r *http.Request) {
	app.ServeHTTP(w, r)
}
