//Package models - model functions and entities
package models

import (
	"github.com/rahulvramesh/groot-comments/db"
	log "github.com/sirupsen/logrus"
)

//CommentModel - comment model
type CommentModel struct {
}

//StoreComment -  store comments to db
//returns stored comment
func (cm *CommentModel) StoreComment(comment *Comment) (*Comment, error) {
	db := db.GetSession()

	err := db.Create(&comment).Error

	if err != nil {
		return nil, err
	}

	return comment, nil
}

//GetComments - get all comments by organization name
//returns array commets
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

//DeleteOrgComments -  delete all organization comments / soft delete
func (cm *CommentModel) DeleteOrgComments(orgName string) error {

	db := db.GetSession()

	err := db.Delete(Comment{}, "organization = ?", orgName).Error

	if err != nil {
		return err
	}

	return nil
}

//GetAllMembersByOrg - get all members by organization name
func (cm *CommentModel) GetAllMembersByOrg(orgName string) ([]Member, error) {

	var (
		members []Member
	)

	db := db.GetSession()

	//used raw query for joining
	rows, err := db.Select(`
			DISTINCT(members.ID),
			members.LOGIN,
			members.url,
			members.followers,
			members.FOLLOWING 
	`).
		Table("members").
		Joins("LEFT JOIN comments ON comments.author_id = members.ID").
		Where("comments.organization = ? AND comments.deleted_at IS NULL", orgName).
		Order("members.followers desc").
		Rows()
	defer rows.Close()
	for rows.Next() {

		memberTemp := Member{}

		if err = rows.Scan(
			&memberTemp.ID,
			&memberTemp.Login,
			&memberTemp.URL,
			&memberTemp.Followers,
			&memberTemp.Following,
		); err != nil {
			log.Error("failed to scan result")
			return []Member{}, err
		}

		members = append(members, memberTemp)
	}

	return members, nil

}
