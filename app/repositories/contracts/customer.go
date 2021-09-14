package contracts

import (
	"verifyMyTest/m/app/models"
	"verifyMyTest/m/util"
)

type Reader interface {
	FindBy(findParameters util.FindParameters, parameters util.Parameters) *util.Data
	FindAll(parameters util.Parameters) *util.Data
}

type Writer interface {
	Update(Customer *models.Customer, id int) (error, int64)
	Create(Customer *models.Customer) (models.CustomerResponse, error)
	Delete(id int) (error, int64)
}

type CustomerRepository interface {
	Reader
	Writer
}
