package entity

import "gorm.io/gorm"

type Actor struct {
	gorm.Model
	FirstName string
	LastName  string
}
