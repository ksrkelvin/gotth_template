package repository

import (
	"encontradev/internal/dto"
	"encontradev/internal/models"
)

func (r *Repository) CreateExternalUser(profile dto.ExternalAuthProfile) (models.User, error) {
	newUser := models.User{
		Name:         profile.Name,
		Email:        profile.Email,
		Avatar:       profile.Picture,
		AuthProvider: profile.Source,
	}

	if err := r.DB.Create(&newUser).Error; err != nil {
		return models.User{}, err
	}

	return newUser, nil
}
