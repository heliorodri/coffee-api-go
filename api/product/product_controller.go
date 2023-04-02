package api

import (
	"net/http"
	"strconv"

	"coffee-api-go/service"

	models "coffee-api-go/model"
	utils "coffee-api-go/utils"

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
	id := utils.ExtractId(c)

	product, err := p.productService.GetProductByID(uint(id))

	if product == nil && err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": product})
}

func (p *ProductController) CreateProduct(c *gin.Context) {
	var newProduct models.Product

	if err := c.BindJSON(&newProduct); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product, err := p.productService.CreateProduct(&newProduct)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": product})
}

func (p *ProductController) UpdateProduct(c *gin.Context) {
	var newValues models.Product
	id := utils.ExtractId(c)

	if err := c.BindJSON(&newValues); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if newValues.ID <= 0 {
		newValues.ID = id
	}

	updatedProduct, err := p.productService.UpdateProduct(&newValues)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": updatedProduct})
}

func (p *ProductController) DeleteProduct(c *gin.Context) {
	id := utils.ExtractId(c)

	if err := p.productService.DeleteProduct(id); err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "Product with ID " + strconv.FormatUint(uint64(id), 10) + " deleted"})
}

func RegisterProductRoutes(router *gin.Engine, productController *ProductController) {
	router.GET("/products", productController.GetAllProducts)
	router.GET("/products/:id", productController.GetProductByID)
	router.POST("/products", productController.CreateProduct)
	router.PUT("/products/:id", productController.UpdateProduct)
	router.DELETE("/products/:id", productController.DeleteProduct)
}
