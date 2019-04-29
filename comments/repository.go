package comments

import "github.com/rahulvramesh/groot-comments/models"

//Repository -
type Repository interface {
	Store(comment models.Comment) (models.Comment, error)
}
