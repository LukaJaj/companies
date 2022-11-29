package mock

import "gorm.io/gorm"

type Store struct {
}

func New(db *gorm.DB) Store {
	return Store{}
}
