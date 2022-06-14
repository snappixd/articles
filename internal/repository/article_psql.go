package repository

import (
	"articles_psql/internal/models"
	"context"
	"database/sql"
	"errors"
)

type ArticlesRepo struct {
	db *sql.DB
}

func NewArticlesRepo(db *sql.DB) *ArticlesRepo {
	return &ArticlesRepo{
		db: db,
	}
}

func (r *ArticlesRepo) Create(ctx context.Context, article models.Article) error {
	_, err := r.db.Exec("INSERT INTO `articles` (`author`, `title`, `anons`, `text`) VALUES (?, ?, ?, ?)",
		article.Author, article.Title, article.Anons, article.Text)
	return err
}

func (r *ArticlesRepo) Update(ctx context.Context, article models.Article) error {
	_, err := r.db.Exec("UPDATE articles SET author=?, title=?, anons=?, text=? WHERE id=?",
		article.Author, article.Title, article.Anons, article.Text, article.ID)

	return err
}

func (r *ArticlesRepo) GetAll(ctx context.Context) ([]models.Article, error) {
	rows, err := r.db.Query("SELECT * FROM articles")
	if err != nil {
		return nil, err
	}

	articles := make([]models.Article, 0)

	for rows.Next() {
		var article models.Article
		if err := rows.Scan(&article.ID, &article.Author, &article.Title, &article.Anons, &article.Text, &article.Photo); err != nil {
			return nil, err
		}

		articles = append(articles, article)
	}

	return articles, rows.Err()
}

func (r *ArticlesRepo) GetByID(ctx context.Context, id int) (models.Article, error) {
	var article models.Article

	err := r.db.QueryRow("SELECT id, author, title, anons, text, photo FROM articles WHERE id=?", id).
		Scan(&article.ID, &article.Author, &article.Title, &article.Anons, &article.Text, &article.Photo)
	if err == sql.ErrNoRows {
		return article, errors.New("No Article with id was found")
	}

	return article, err
}

func (r *ArticlesRepo) Delete(ctx context.Context, id int) error {
	_, err := r.db.Exec("DELETE FROM articles WHERE id=?", id)

	return err
}
