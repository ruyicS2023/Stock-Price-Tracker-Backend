package main

import (
	"stock-price-tracker/api"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Register all API routes
	api.RegisterRoutes(r)

	// Start the server
	r.Run(":8080")
}
