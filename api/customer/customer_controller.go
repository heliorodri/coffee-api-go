package api

import (
	model "coffee-api-go/model/customer"
	service "coffee-api-go/service/customer"
	"coffee-api-go/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CustomerController struct {
	customerService service.CustomerService
}

func NewCustomerController(customerService *service.CustomerService) *CustomerController {
	return &CustomerController{
		customerService: *customerService,
	}
}

func (controller *CustomerController) GetAll(c *gin.Context) {
	customers, err := controller.customerService.GetAll()

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": customers})
}

func (controller *CustomerController) GetByID(c *gin.Context) {
	id := utils.ExtractId(c)

	customer, err := controller.customerService.GetByID(uint(id))

	if customer == nil && err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": customer})
}

func (controller *CustomerController) Create(c *gin.Context) {
	var newCustomer model.Customer

	if err := c.BindJSON(&newCustomer); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	customer, err := controller.customerService.Create(&newCustomer)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": customer})
}

func (controller *CustomerController) Update(c *gin.Context) {
	var newCustomer model.Customer

	if err := c.BindJSON(&newCustomer); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	customer, err := controller.customerService.Update(&newCustomer)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": customer})
}

func (controller *CustomerController) Delete(c *gin.Context) {
	id := utils.ExtractId(c)

	err := controller.customerService.Delete(uint(id))

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "Customer deleted"})
}

func RegisterCustomerRoutes(router *gin.Engine, customerController *CustomerController) {
	router.GET("/customers", customerController.GetAll)
	router.GET("/customers/:id", customerController.GetByID)
	router.POST("/customers", customerController.Create)
	router.PUT("/customers", customerController.Update)
	router.DELETE("/customers/:id", customerController.Delete)
}
