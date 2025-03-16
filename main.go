package main

import (
	"log"
	"net/http"
	"os"

	"bringkad-arena-service-go/config"
	"bringkad-arena-service-go/databases"
	v1 "bringkad-arena-service-go/routes/v1"
	"bringkad-arena-service-go/validators"

	"github.com/gin-gonic/gin"
)

func main() {
	// Init gin router
	router := gin.Default()

	// Load ENV
	config.LoadENV()

	// Connect to DB
	databases.ConnectDB()

	// Register custom validator for request validation
	validators.RegisterRequestCustomValidator()

	// Health check API
	router.GET("/api/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "All system running successfully"})
	})

	// Swagger API
	// router.GET("/api/swagger")

	// Register API routes
	apiGroup := router.Group("/api")
	v1.RegisterRoutes(apiGroup.Group("v1"))

	// Start server
	port := os.Getenv("APP_PORT")
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start system at port %v, error: %v", port, err)
	}

}
