package models

import "github.com/jinzhu/gorm"

//Comment - comment table
type Comment struct {
	gorm.Model

	Comment      string `json:"comment" validate:"store_comment:required"`
	AuthorID     int    `json:"author"`
	Organization string `json:"organization"`
}
