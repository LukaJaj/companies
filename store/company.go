package store

import (
	"companies-crud/model"

	"gorm.io/gorm"
)

func (s Store) Company(companyId string) (*model.Company, error) {
	var company *model.Company
	err := s.db.Transaction(func(tx *gorm.DB) error {
		tx = s.db.First(&company, "id = ?", companyId)
		return tx.Error
	})

	return company, err
}

func (s Store) InsertCompany(company *model.Company) error {

	err := s.db.Transaction(func(tx *gorm.DB) error {
		tx = s.db.First(&company, "id = ?", company.Id)
		if tx.Error != gorm.ErrRecordNotFound {
			return model.ErrCompanyAlreadyExists
		}
		return tx.Error
	})

	err = s.db.Transaction(func(tx *gorm.DB) error {
		tx = s.db.Create(&company)
		return tx.Error
	})

	return err
}

func (s Store) UpdateCompany(company *model.Company) error {
	err := s.db.Transaction(func(tx *gorm.DB) error {
		tx = s.db.Save(&company)
		return tx.Error
	})
	return err

}

func (s Store) DeleteCompany(companyId string) error {
	err := s.db.Transaction(func(tx *gorm.DB) error {
		var company model.Company
		tx = s.db.Where("id = ?", companyId).Delete(&company)
		return tx.Error
	})
	return err
}

func (s Store) CompanyType(companyType string) (string, error) {
	var t *model.CompanyType

	err := s.db.Transaction(func(tx *gorm.DB) error {
		tx = s.db.Table("company_type").First(&t, "name = ?", companyType)
		return tx.Error
	})

	return t.Name, err

}
