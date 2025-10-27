package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email        string
	Password     string
	Name         string
	Avatar       string
	AuthProvider string
}
