package main

import (
	productApi "coffee-api-go/api/product"
	productModel "coffee-api-go/model/product"
	productRepo "coffee-api-go/repository/product"
	productService "coffee-api-go/service/product"

	customerApi "coffee-api-go/api/customer"
	customerModel "coffee-api-go/model/customer"
	customerRepo "coffee-api-go/repository/customer"
	customerService "coffee-api-go/service/customer"

	"coffee-api-go/db"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func main() {
	r := gin.Default()

	db := db.NewConnection()

	productRoutesAndDBMigration(r, db)
	customerRoutesAndDBMigration(r, db)

	r.Run(":8080")
}

func productRoutesAndDBMigration(router *gin.Engine, db *gorm.DB) {
	db.AutoMigrate(&productModel.Product{})

	productRepository := productRepo.NewProductRepository(db)
	productService := productService.NewProductService(&productRepository)
	productController := productApi.NewProductController(productService)
	productApi.RegisterProductRoutes(router, productController)
}

func customerRoutesAndDBMigration(router *gin.Engine, db *gorm.DB) {
	db.AutoMigrate(&customerModel.Customer{})

	customerRepository := customerRepo.NewCustomerRepository(db)
	customerService := customerService.NewCustomerService(&customerRepository)
	customerController := customerApi.NewCustomerController(customerService)
	customerApi.RegisterCustomerRoutes(router, customerController)
}
