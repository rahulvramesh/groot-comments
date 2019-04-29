package models

import "time"

//Member - member type
type Member struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index" faker:"-"`

	Login     string `json:"login" faker:"email"`
	URL       string `json:"profile_picture" faker:"url"`
	Followers int    `json:"followers"`
	Following int    `json:"following"`
}
