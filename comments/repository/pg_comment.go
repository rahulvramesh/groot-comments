package repository

import (
	"database/sql"

	"github.com/rahulvramesh/groot-comments/comments"
)

type pgCommentRepository struct {
	Conn *sql.DB
}

func NewPgArticleRepository(Conn *sql.DB) comments.Repository {
	return &pgArticleRepository{Conn}
}
