package main

import (
	productApi "coffee-api-go/api/product"
	"coffee-api-go/db"
	productModel "coffee-api-go/model/product"
	productRepo "coffee-api-go/repository/product"
	productService "coffee-api-go/service/product"

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
	db.AutoMigrate(&productModel.Product{})

	productRepository := productRepo.NewProductRepository(db)
	productService := productService.NewProductService(&productRepository)
	productController := productApi.NewProductController(productService)
	productApi.RegisterProductRoutes(router, productController)
}
