package models

import "gorm.io/gorm"

type Name struct {
	gorm.Model
	Name        string
	Surname     string
	Patronymic  string
	Age         int
	Gender      string
	Nationality string
}
