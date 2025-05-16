package main

import (
	"flag"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "online-task/docs"
	"online-task/internal/auth"
	"online-task/internal/booking"
	"online-task/internal/event"
	"online-task/internal/tag"
	"online-task/internal/upload"
	"online-task/pkg/database"
	"online-task/pkg/seed"
)

// @title           Event Booking API
// @version         1.0
// @description     A REST API for an event booking system

// @contact.name   API Support
// @contact.email  support@example.com

// @license.name  MIT
// @license.url   https://opensource.org/licenses/MIT

// @host      areeb.ddns.net:2025
// @BasePath  /api

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
func main() {
	// Parse command line flags
	seedAdmin := flag.Bool("seed-admin", false, "Seed admin user")
	flag.Parse()

	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Printf("Error loading .env file: %v", err)
	}

	// Initialize database
	database.Init()

	// Seed admin user if flag is set
	if *seedAdmin {
		if err := seed.SeedAdminUser(); err != nil {
			log.Fatal("Failed to seed admin user:", err)
		}
		// Exit after seeding if that's the only operation requested
		if flag.NFlag() == 1 {
			return
		}
	}

	// Initialize router
	r := gin.Default()

	// Serve static files for uploads
	r.Static("/uploads", "./uploads")

	// Configure CORS
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"https://areeb.ddns.net", "http://areeb.ddns.net", "http://localhost:2025"}
	config.AllowCredentials = true
	config.AllowHeaders = append(config.AllowHeaders, "Authorization")
	r.Use(cors.New(config))

	// Swagger documentation
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// API routes
	api := r.Group("/api")
	{
		// Auth routes
		authGroup := api.Group("/auth")
		{
			authGroup.POST("/register", auth.RegisterHandler)
			authGroup.POST("/login", auth.LoginHandler)
		}

		// Events routes
		eventsGroup := api.Group("/events")
		{
			eventsGroup.GET("", event.GetAllEventsHandler)
			eventsGroup.GET("/:id", event.GetEventHandler)
			eventsGroup.POST("", auth.AuthMiddleware(), auth.AdminMiddleware(), event.CreateEventHandler)
			eventsGroup.PUT("/:id", auth.AuthMiddleware(), auth.AdminMiddleware(), event.UpdateEventHandler)
			eventsGroup.DELETE("/:id", auth.AuthMiddleware(), auth.AdminMiddleware(), event.DeleteEventHandler)
		}

		// Tags routes
		tagsGroup := api.Group("/tags")
		{
			tagsGroup.GET("", tag.GetAllTagsHandler)
			tagsGroup.GET("/:id", tag.GetTagHandler)
			tagsGroup.POST("", auth.AuthMiddleware(), auth.AdminMiddleware(), tag.CreateTagHandler)
			tagsGroup.PUT("/:id", auth.AuthMiddleware(), auth.AdminMiddleware(), tag.UpdateTagHandler)
			tagsGroup.DELETE("/:id", auth.AuthMiddleware(), auth.AdminMiddleware(), tag.DeleteTagHandler)
		}

		// Bookings routes
		bookingsGroup := api.Group("/bookings")
		{
			bookingsGroup.Use(auth.AuthMiddleware())
			bookingsGroup.POST("", booking.CreateBookingHandler)
			bookingsGroup.GET("/user", booking.GetUserBookingsHandler)
		}

		// Upload routes
		uploadGroup := api.Group("/upload")
		{
			uploadGroup.Use(auth.AuthMiddleware(), auth.AdminMiddleware())
			uploadGroup.POST("/image", upload.UploadImageHandler)
		}
	}

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
