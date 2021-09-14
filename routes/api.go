package routes

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"
	"verifyMyTest/m/app/http/controllers"
)

func Setup() *gin.Engine {

	gin.SetMode("release")

	r := gin.New()
	r.Use(gin.Recovery())

	// Handle error response when a route is not defined
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"message": "Not found"})
	})

	// Display Swagger documentation
	r.StaticFile("doc/swagger.json", "doc/swagger.json")
	config := &ginSwagger.Config{
		URL: "/doc/swagger.json", //The url pointing to API definition
	}
	// use ginSwagger middleware to
	r.GET("/swagger/*any", ginSwagger.CustomWrapHandler(config, swaggerFiles.Handler))

	customerController := new(controllers.CustomerController)

	// swagger:route GET /ping common getPing
	//
	// Ping
	//
	// Get Ping and reply Pong
	//
	//     Responses:
	//       200:
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// swagger:route GET /customer/:field customer FindParameters
	//
	// Customer
	//
	// Get customer data by field
	//
	//     Responses:
	//       200: CustomersResponse
	//       400: CustomerErrorResponse

	r.GET("/customer/:field/:value", customerController.FindCustomerBy)

	//swagger:route GET /customers customer PaginateParameters
	//
	//Customers paginated list
	//
	//Get customers list data
	//
	//    Responses:
	//      200: CustomersResponse
	r.GET("/customers", customerController.FindCustomers)

	// swagger:route POST /customer customer CreateCustomer
	//
	// New customer
	//
	// Create new customer
	//
	//     Responses:
	//       201: CustomerResponse
	//       400: CustomerErrorResponse
	r.POST("/customer", customerController.CreateCustomer)

	// swagger:route PUT /customer/:id customer UpdateCustomer
	//
	// Update customer
	//
	// Update existing customer
	//
	//     Responses:
	//       200: CustomerResponse
	//       400: CustomerErrorResponse
	r.PUT("/customer/:id", customerController.UpdateCustomer)

	// swagger:route DELETE /customer/:id customer DeleteCustomer
	//
	// Delete customer
	//
	// Delete existing customer
	//
	//     Responses:
	//       200:
	r.DELETE("/customer/:id", customerController.DeleteCustomer)

	return r
}
