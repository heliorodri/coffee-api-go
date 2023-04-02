package entity

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	Name  string `json:"name"`
	Email string `json:"email"`
}
