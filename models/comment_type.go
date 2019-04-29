package models

import "github.com/jinzhu/gorm"

//Comment - comment table
type Comment struct {
	gorm.Model

	Comment  string `json:"comment"`
	AuthorID int    `json:"author"`
}
