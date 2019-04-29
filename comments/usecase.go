package comments

import "github.com/rahulvramesh/groot-comments/models"

//Usecase -
type Usecase interface {
	Store(comment models.Comment) (models.Comment, error)
}
