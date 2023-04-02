package main

import (
	productApi "coffee-api-go/api/product"
	"coffee-api-go/db"
	models "coffee-api-go/model"
	"coffee-api-go/repository"
	"coffee-api-go/service"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create a new Gin router
	r := gin.Default()

	db := db.NewConnection()
	db.AutoMigrate(&models.Product{})

	ProductRepository := repository.NewProductRepository(db)
	productService := service.NewProductService(&ProductRepository)
	productController := productApi.NewProductController(productService)
	productApi.RegisterProductRoutes(r, productController)

	r.Run(":8080")

	// fmt.Print("teste")
}
