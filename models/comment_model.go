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

//GetComments - get all comments
func (cm *CommentModel) GetComments(orgName string) ([]Comment, error) {

	var (
		commets []Comment
	)

	db := db.GetSession()

	err := db.Where("organization = ?", orgName).Find(&commets).Error

	if err != nil {
		return nil, err
	}

	return commets, nil

}

//DeleteOrgComments -  delete all organization comments
func (cm *CommentModel) DeleteOrgComments(orgName string) error {

	db := db.GetSession()

	err := db.Delete(Comment{}, "organization = ?", orgName).Error

	if err != nil {
		return err
	}

	return nil
}
