package main

import (
	"log"

	"IkonKutz.API/initializers"
	"IkonKutz.API/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	// Load environment variables first
	initializers.LoadEnvVariables()

	//
	initializers.LoadConfig()

	// Connect to the database
	initializers.ConnectToDB()

	// Migrate the database
	initializers.SyncDatabase()

	// Seed the admin user if credentials are provided
	initializers.SeedAdminUser()
}

func main() {
	// Set Gin mode based on configuration
	if initializers.AppConfig.GinMode != "" {
		gin.SetMode(initializers.AppConfig.GinMode)
	}

	router := gin.Default()

	// Minimal production CORS setup.
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{initializers.AppConfig.ClientOrigin},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Register routes
	routes.RegisterRoutes(router)

	port := initializers.GetPort()

	// Start the server
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
