package models

import "time"

//Comment - comment table
type Comment struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index" faker:"-"`

	Comment      string `json:"comment" validate:"store_comment:required" faker:"sentence"`
	AuthorID     int    `json:"author"`
	Organization string `json:"organization"`
}
