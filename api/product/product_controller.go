package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductController struct{}

func (p *ProductController) GetAllProducts(c *gin.Context) {
	// your logic for fetching all products goes here
	c.JSON(http.StatusOK, gin.H{"data": "All products"})
}

func (p *ProductController) GetProductByID(c *gin.Context) {
	// your logic for fetching a product by ID goes here
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"data": "Product with ID " + id})
}

func (p *ProductController) CreateProduct(c *gin.Context) {
	// your logic for creating a new product goes here
	c.JSON(http.StatusOK, gin.H{"data": "Product created"})
}

func (p *ProductController) UpdateProduct(c *gin.Context) {
	// your logic for updating a product goes here
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"data": "Product with ID " + id + " updated"})
}

func (p *ProductController) DeleteProduct(c *gin.Context) {
	// your logic for deleting a product goes here
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"data": "Product with ID " + id + " deleted"})
}

func RegisterProductRoutes(router *gin.Engine, productController *ProductController) {
	router.GET("/products", productController.GetAllProducts)
	router.GET("/products/:id", productController.GetProductByID)
	router.POST("/products", productController.CreateProduct)
	router.PUT("/products/:id", productController.UpdateProduct)
	router.DELETE("/products/:id", productController.DeleteProduct)
}
