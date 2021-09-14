package repositories

import (
	"github.com/jinzhu/gorm"
	"verifyMyTest/m/app/models"
	"verifyMyTest/m/util"
)

type CustomerRepository struct {
	database *gorm.DB
}

func (r *CustomerRepository) Update(customer *models.Customer, id int) (error, int64) {
	result := r.database.Model(&customer).Where("id = ?", id).Update(&customer)

	err := result.Error

	if err != nil {
		util.Error(util.HTTPError, err.Error())
		return err, result.RowsAffected
	}

	return nil, result.RowsAffected
}

func (r *CustomerRepository) Create(customer *models.Customer) (models.CustomerResponse, error) {

	err := r.database.Create(&customer).Error

	response := models.CustomerResponse{
		Name:    customer.Name,
		Age:     customer.Age,
		Email:   customer.Email,
		Address: customer.Address,
	}

	if err != nil {
		util.Error(util.HTTPError, err.Error())
	}

	return response, err
}

func (r *CustomerRepository) Delete(id int) (error, int64) {

	customer := models.Customer{}

	result := r.database.Where("id = ?", id).Unscoped().Delete(&customer)
	err := result.Error
	if err != nil {
		util.Error(util.HTTPError, err.Error())
		return err, result.RowsAffected
	}

	return nil, result.RowsAffected
}

func NewCustomerRepository(database *gorm.DB) *CustomerRepository {
	return &CustomerRepository{
		database: database,
	}
}

func (r *CustomerRepository) FindAll(parameters util.Parameters) *util.Data {

	var model []models.CustomerResponse
	var findParameters util.FindParameters
	result := util.GetPage(r.database, &model, findParameters, parameters)
	return result
}

func (r *CustomerRepository) FindBy(findParameters util.FindParameters, parameters util.Parameters) *util.Data {
	var model []models.CustomerResponse
	result := util.GetPage(r.database, &model, findParameters, parameters)
	return result
}
