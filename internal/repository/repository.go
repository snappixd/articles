package repository

import (
	"articles_psql/internal/models"
	"context"
	"database/sql"
)

type Articles interface {
	Create(ctx context.Context, article models.Article) error
	GetAll(ctx context.Context) ([]models.Article, error)
	GetByID(ctx context.Context, id int) (models.Article, error)
	Delete(ctx context.Context, id int) error
}

type Repositories struct {
	Articles Articles
}

func NewRepositories(db *sql.DB) *Repositories {
	return &Repositories{
		Articles: NewArticlesRepo(db),
	}
}
