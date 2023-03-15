package main

import (
	productApi "coffee-api-go/api/product"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create a new Gin router
	r := gin.Default()

	productController := &productApi.ProductController{}
	productApi.RegisterProductRoutes(r, productController)

	// Run the server on port 8080
	r.Run(":8080")
}
