package service

import (
	"companies-crud/model"
	"errors"
)

func (s Service) Company(companyId string) (*model.Company, error) {
	company, err := s.store.Company(companyId)
	if err != nil {
		return nil, err
	}

	return company, nil
}

func (s Service) CreateCompany(companyDto model.CompanyDto) error {
	company, err := s.dtoToCompany(&companyDto)
	if err != nil {
		return err
	}

	err = s.store.InsertCompany(company)
	if errors.Is(err, model.ErrCompanyAlreadyExists) {
		return model.ErrCompanyAlreadyExists
	}

	return err
}

func (s Service) UpdateCompany(company *model.Company) error {
	err := s.store.UpdateCompany(company)
	return err
}

func (s Service) DeleteCompany(companyId string) error {
	err := s.store.DeleteCompany(companyId)
	return err
}

func (s Service) dtoToCompany(company *model.CompanyDto) (*model.Company, error) {

	companyType, err := s.store.CompanyType(company.Type)
	if err != nil {
		return nil, err
	}

	return &model.Company{
		Id:                company.Id,
		Name:              company.Name,
		Description:       company.Description,
		AmountOfEmployees: company.AmountOfEmployees,
		Registered:        company.Registered,
		Type:              companyType,
	}, nil
}
