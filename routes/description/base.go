package description

import (
	"verifyMyTest/m/app/models"
)

type BaseResponseBody struct {
	// Response status.
	// One of next values:
	// - ok,
	// - error
	Status string
}

type customerBaseResponse struct {
	TotalRecords int
	CurrentPage  string
	TotalPages   int64
	Records      *models.CustomerResponse
}
