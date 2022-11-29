package model

import (
	"github.com/google/uuid"
)

type Company struct {
	Id                uuid.UUID `gorm:"primaryKey; type:uuid not null" json:"id"`
	Name              string    `gorm:"index; type:varchar(15) unique not null" json:"name"`
	Description       string    `gorm:"type:varchar(3000)" json:"description"`
	AmountOfEmployees int       `gorm:"type:integer  not null" json:"amount_of_employees"`
	Registered        bool      `gorm:"type:boolean  not null" json:"registered"`
	Type              string    `gorm:"type:text  not null" json:"type"`
}

type CompanyType struct {
	Id   int    `gorm:"primaryKey; type:integer not null" json:"id"`
	Name string `gorm:"index; primaryKey; type:varchar(30) unique not null" json:"name"`
}

type CompanyDto struct {
	Id                uuid.UUID `json:"id"`
	Name              string    `json:"name"`
	Description       string    `json:"description,omitempty"`
	AmountOfEmployees int       `json:"amount_of_employees"`
	Registered        bool      `json:"registered"`
	Type              string    `json:"type"`
}

type UpdateCompanyDto struct {
	Name              string `json:"name,omitempty"`
	Description       string `json:"description,omitempty"`
	AmountOfEmployees int    `json:"amount_of_employees,omitempty"`
	Registered        bool   `json:"registered,omitempty"`
	Type              string `json:"type,omitempty"`
}
