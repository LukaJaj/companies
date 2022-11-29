package mock

import (
	"companies-crud/model"
	"errors"

	"github.com/google/uuid"
)

func (s Store) Company(companyId string) (*model.Company, error) {
	if companyId == "0104c7e9-03b6-4d2f-a725-03730244d8c9" {
		return &model.Company{
			Id:                uuid.MustParse("0104c7e9-03b6-4d2f-a725-03730244d8c9"),
			Name:              "facebook",
			Description:       "social media company",
			AmountOfEmployees: 5000,
			Registered:        true,
			Type:              "corporations",
		}, nil

	} else if companyId == "error_company" {
		return nil, errors.New("failed to get company")
	} else if companyId == "company_not_found" {
		return nil, model.ErrCompanyNotFound
	} else if companyId == "5504c7e9-45yu-4d2f-a725-03730244dvc8" {
		return nil, errors.New("company does not exists")
	} else if companyId == "9004c7e9-03b6-4d2f-a725-03730244d8d7" {
		return &model.Company{
			Id:                uuid.MustParse("9004c7e9-03b6-4d2f-a725-03730244d8d7"),
			Name:              "twitter",
			Description:       "social media company?",
			AmountOfEmployees: 1000,
			Registered:        true,
			Type:              "startup",
		}, nil
	} else if companyId == "0204c7e9-03b6-4d2f-a725-03730244d8c9" {
		return &model.Company{
			Id:                uuid.MustParse("0204c7e9-03b6-4d2f-a725-03730244d8c9"),
			Name:              "amazon",
			Description:       "retail company",
			AmountOfEmployees: 10000,
			Registered:        true,
			Type:              "corporations",
		}, nil
	}

	return nil, nil
}

func (s Store) InsertCompany(company *model.Company) error {
	if company.Id.String() == "4304c7e9-03b6-4d2f-a725-03730244d800" {
		return model.ErrCompanyAlreadyExists
	}
	if company.Id.String() == "9104c7e9-03b6-4d2f-a725-03730244d832" {
		return errors.New("unexpected error")
	}
	return nil
}

func (s Store) UpdateCompany(company *model.Company) error {
	if company.Id.String() == "9004c7e9-03b6-4d2f-a725-03730244d8d7" {
		return errors.New("error while updating")
	}
	return nil
}

func (s Store) DeleteCompany(companyId string) error {
	if companyId == "2704c7e9-03b6-4d2f-a725-03730244d8k1" {
		return nil
	}
	return errors.New("something went wrong on db")
}

func (s Store) CompanyType(companyType string) (string, error) {
	return "", nil
}
