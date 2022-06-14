package service

import (
	"articles_psql/internal/models"
	"articles_psql/internal/repository"
	"context"
)

type Articles interface {
	Create(ctx context.Context, article models.Article) error
	GetAll(ctx context.Context) ([]models.Article, error)
	Edit(ctx context.Context, article models.Article) error
	GetByID(ctx context.Context, id int) (models.Article, error)
	Delete(ctx context.Context, id int) error
}

type Services struct {
	Articles Articles
}

type Deps struct {
	Repos *repository.Repositories
}

func NewService(deps Deps) *Services {
	articlesService := NewArticlesService(deps.Repos.Articles)

	return &Services{
		Articles: articlesService,
	}
}
