package controllers

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
	"strconv"
	"verifyMyTest/m/app/models"
	"verifyMyTest/m/app/repositories"
	"verifyMyTest/m/app/services"
	"verifyMyTest/m/util"
)

const RecordNotFound = "record not found"
const EmptyData = "data is empty"
const RouteParameterIsMissing = "route parameter is missing"
const ArgParameterIsEmpty = "arg parameter is empty"
const OkMessage = "ok"

var newDb, _ = util.NewDb()
var customerRepo = repositories.NewCustomerRepository(newDb)
var customerService = services.NewCustomerService(customerRepo)

type CustomerController struct{}

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (f *CustomerController) FindCustomerBy(c *gin.Context) {

	parameters := getQueryParams(c)

	field := c.Param("field")

	if field == "" {
		errorResponse("field route parameter is empty", c)
		c.Abort()
		return
	}

	arg := c.Param("value")

	if arg == "" {
		errorResponse(ArgParameterIsEmpty, c)
		c.Abort()
		return
	}

	findParameters := util.FindParameters{
		Field: field,
		Arg:   arg,
	}

	result := customerService.FindBy(findParameters, *parameters)

	successResponse(OkMessage, http.StatusOK, result, c)
}

func (f *CustomerController) FindCustomers(c *gin.Context) {
	parameters := getQueryParams(c)

	result := customerService.FindAll(*parameters)

	successResponse(OkMessage, http.StatusOK, result, c)

}

func (f *CustomerController) CreateCustomer(c *gin.Context) {

	var customer models.Customer

	v := validator.New()

	if err := c.ShouldBind(&customer); err != nil {
		errorResponse(err.Error(), c)
		c.Abort()
		return
	}

	err := v.Struct(&customer)

	if err != nil {
		errorResponse(err.Error(), c)
		c.Abort()
		return
	}

	result, err := customerService.Create(&customer)

	if err != nil {
		errorResponse(err.Error(), c)
		c.Abort()
		return
	}

	successResponse(OkMessage, http.StatusCreated, result, c)
}

func (f *CustomerController) UpdateCustomer(c *gin.Context) {

	var customer models.Customer

	if err := c.ShouldBind(&customer); err != nil {
		errorResponse(err.Error(), c)
		c.Abort()
		return
	}

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		errorResponse(err.Error(), c)
		c.Abort()
		return
	}

	result, rows := customerService.Update(&customer, id)

	if result != nil {
		errorResponse(err.Error(), c)
		c.Abort()
		return
	}

	if rows == 0 {
		errorResponse(EmptyData, c)
		c.Abort()
		return
	}

	successResponse(OkMessage, http.StatusOK, result, c)
}

func (f *CustomerController) DeleteCustomer(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		errorResponse(err.Error(), c)
		c.Abort()
		return
	}

	if id == 0 {
		errorResponse(RouteParameterIsMissing, c)
		c.Abort()
		return
	}

	result, rows := customerService.Delete(id)

	if result != nil {
		errorResponse(result.Error(), c)
		c.Abort()
		return
	}

	if rows == 0 {
		errorResponse(RecordNotFound, c)
		c.Abort()
		return
	}

	successResponse(OkMessage, http.StatusOK, result, c)
}

func errorResponse(message string, c *gin.Context) {
	util.Error(util.HTTPError, message)
	c.JSON(http.StatusBadRequest, Response{
		Success: false,
		Message: message,
	})
}

func successResponse(message string, status int, data interface{}, c *gin.Context) {
	c.JSON(status, Response{
		Success: true,
		Message: message,
		Data:    data,
	})
}

func getQueryParams(c *gin.Context) *util.Parameters {

	page, exists := c.GetQuery("page")

	if !exists {
		page = "1"
	}

	perPage, exists := c.GetQuery("perPage")

	if !exists {
		perPage = "10"
	}

	sortField, exists := c.GetQuery("field")

	if !exists {
		sortField = "name"
	}

	sortDirection, exists := c.GetQuery("direction")

	if !exists {
		sortDirection = "asc"
	}

	return &util.Parameters{
		Page:          page,
		PerPage:       perPage,
		SortField:     sortField,
		SortDirection: sortDirection,
	}

}
