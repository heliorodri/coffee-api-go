package main

import "github.com/gin-gonic/gin"

func main() {
	// Create a new Gin router
	r := gin.Default()

	// Define your API routes here
	r.GET("/coffee-shop/v1/hello/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

	// Run the server on port 8080
	r.Run(":8080")
}
