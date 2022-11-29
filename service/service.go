package service

import "companies-crud/model"

type Store interface {
	Company(companyId string) (*model.Company, error)
	InsertCompany(company *model.Company) error
	UpdateCompany(company *model.Company) error
	DeleteCompany(companyId string) error

	CompanyType(companyType string) (string, error)
}
type Service struct {
	store Store
}

func New(s Store) *Service {
	return &Service{
		store: s,
	}
}
