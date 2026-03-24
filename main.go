package main

import (
	"log"

	"IkonKutz.API/initializers"
	"IkonKutz.API/routes"
	"github.com/gin-gonic/gin"
)

func init() {
	// Load environment variables first
	initializers.LoadEnvVariables()

	// Connect to the database
	initializers.ConnectToDB()

	// Migrate the database
	initializers.SyncDatabase()
}

func main() {
	router := gin.Default()

	// Register routes
	routes.RegisterRoutes(router)

	port := initializers.GetPort()

	// Start the server
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
