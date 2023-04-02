package api

import (
	"net/http"

	"coffee-api-go/service"

	models "coffee-api-go/model"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	productService service.ProductService
}

func NewProductController(productService *service.ProductService) *ProductController {
	return &ProductController{
		productService: *productService,
	}
}

func (p *ProductController) GetAllProducts(c *gin.Context) {
	products, err := p.productService.GetAllProducts()

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{"data": products})
}

func (p *ProductController) GetProductByID(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"data": "Product with ID " + id})
}

func (p *ProductController) CreateProduct(c *gin.Context) {
	var newProduct models.Product
	if err := c.BindJSON(&newProduct); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	p.productService.CreateProduct(&newProduct)
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
