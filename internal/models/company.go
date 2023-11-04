package models

import "gorm.io/gorm"

type Company struct {
	gorm.Model
	Name     string `json:"name" gorm:"unique" validate:"required"`
	Location string `json:"location" validate:"required"`
	Type     string `json:"type" validate:"required"`
}

type Jobs struct {
	gorm.Model
	Company     Company `json:"-" gorm:"ForeignKey:company_id"`
	CompanyID   uint    `json:"company_id"`
	Name        string  `json:"name"`
	Salary      string  `json:"salary"`
	Description string  `json:"description"`
}
