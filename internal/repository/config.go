package repository

import (
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func RegisterRepository(DB *gorm.DB) (repository *Repository, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
	}()
	repo := &Repository{
		DB: DB,
	}

	return repo, err
}
