package models

import (
	"github.com/rahulvramesh/groot-comments/db"
)

//CommentModel - comment model
type CommentModel struct {
}

//StoreComment -  store comment to db
func (cm *CommentModel) StoreComment(comment *Comment) (*Comment, error) {
	db := db.GetSession()

	err := db.Create(&comment).Error

	if err != nil {
		return nil, err
	}

	return comment, nil
}
