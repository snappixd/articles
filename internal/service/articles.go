package service

import (
	"articles_psql/internal/models"
	"articles_psql/internal/repository"
	"context"
)

type ArticlesService struct {
	repo repository.Articles
}

func NewArticlesService(repo repository.Articles) *ArticlesService {
	return &ArticlesService{
		repo: repo,
	}
}

func (s *ArticlesService) Create(ctx context.Context, article models.Article) error {
	return s.repo.Create(ctx, article)
}

func (s *ArticlesService) Edit(ctx context.Context, article models.Article) error {
	return s.repo.Update(ctx, article)
}

func (s *ArticlesService) GetAll(ctx context.Context) ([]models.Article, error) {
	return s.repo.GetAll(ctx)
}

func (s *ArticlesService) GetByID(ctx context.Context, id int) (models.Article, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *ArticlesService) Delete(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}
