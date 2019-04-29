package models

import (
	"github.com/rahulvramesh/groot-comments/db"
)

//StoreComment -  store comment to db
func StoreComment(comment *Comment) {
	db := db.GetSession()

	db.Create(&comment)

}
