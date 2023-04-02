package main

import (
	productApi "coffee-api-go/api/product"
	"coffee-api-go/db"
	models "coffee-api-go/model"
	"coffee-api-go/repository"
	"coffee-api-go/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func main() {
	r := gin.Default()

	db := db.NewConnection()

	productRoutesAndDBMigration(r, db)

	r.Run(":8080")
}

func productRoutesAndDBMigration(router *gin.Engine, db *gorm.DB) {
	db.AutoMigrate(&models.Product{})

	productRepository := repository.NewProductRepository(db)
	productService := service.NewProductService(&productRepository)
	productController := productApi.NewProductController(productService)
	productApi.RegisterProductRoutes(router, productController)
}
