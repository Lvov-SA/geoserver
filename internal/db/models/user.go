package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	login    string
	Passowrd string
}
