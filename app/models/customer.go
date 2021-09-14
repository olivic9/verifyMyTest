package models

import (
	"time"
)

type Customer struct {
	Name      string    `json:"name" validate:"required,min=5,max=100"`
	Age       uint      `json:"age" validate:"required,min=2"`
	Email     string    `json:"email" validate:"required,min=8,max=255"`
	Password  string    `json:"password" validate:"required,min=5,max=30"`
	Address   string    `json:"address" validate:"required,min=5,max=255"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CustomerModel struct{}

type CustomerResponse struct {
	Name    string `json:"name" binding:"required"`
	Age     uint   `json:"age" binding:"required"`
	Email   string `json:"email" binding:"required"`
	Address string `json:"address" binding:"required"`
}

type Tabler interface {
	TableName() string
}

func (CustomerResponse) TableName() string {
	return "customers"
}
