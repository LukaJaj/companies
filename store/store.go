package store

import (
	"companies-crud/model"
	"errors"

	"gorm.io/gorm"
)

type Store struct {
	db gorm.DB
}

func NewStore(db gorm.DB) Store {
	return Store{
		db: db,
	}
}

func Migrate(db *gorm.DB) error {
	if err := db.AutoMigrate(&model.Company{}, &model.CompanyType{}); err == nil && db.Migrator().HasTable(&model.CompanyType{}) {
		if err := db.First(&model.CompanyType{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			db.Transaction(func(tx *gorm.DB) error {
				var companyTypes = []model.CompanyType{
					{Name: "corporations"},
					{Name: "nonprofit"},
					{Name: "cooperative"},
					{Name: "sole_proprietorship"},
				}
				db.Create(companyTypes)
				return tx.Error
			})

		}
	}
	return nil
}
