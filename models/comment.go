package models

import (
	"github.com/jinzhu/gorm"
)

//Comment -  comment table
type Comment struct {
	gorm.Model
	Comment  string `json:"content" validate:"required"`
	AuthorID int    `json:"author_id"`
}
