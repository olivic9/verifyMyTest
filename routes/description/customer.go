package description

import (
	"verifyMyTest/m/app/models"
	"verifyMyTest/m/util"
)

// CustomersResponse Customers Success Response
// swagger:response CustomersResponse
type CustomersResponse struct {
	// in: body
	Body struct {
		Success bool
		Message string
		Data    *util.Data
	}
}

// CustomerResponse Customer Success Response
// swagger:response CustomerResponse
type CustomerResponse struct {
	// in: body
	Body struct {
		Success bool
		Message string
		Data    *models.CustomerResponse
	}
}

// CustomerErrorResponse Customer Error Response
// swagger:response CustomerErrorResponse
type CustomerErrorResponse struct {
	// in: body
	Body struct {
		Success bool
		Message string
		Data    interface{}
	}
}

// swagger:parameters CreateCustomer UpdateCustomer
type CustomerParameters struct {
	// in: body
	Body *models.Customer
}

//swagger:parameters PaginateParameters
type PaginateParameters struct {
	// in: query
	*util.Parameters
}

//swagger:parameters FindParameters
type FindParameters struct {
	// in: query
	util.FindParameters
	// in: query
	util.Parameters
}
