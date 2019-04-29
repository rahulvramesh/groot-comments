package models

import (
	"github.com/jinzhu/gorm"
)

//Author table
type Author struct {
	gorm.Model
	Email string `json:"email" validate:"required"`
}
