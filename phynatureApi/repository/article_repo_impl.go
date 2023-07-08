package repository

import (
	"phynatureApi/entity"

	"gorm.io/gorm"
)

type ArticleRepository struct {
	db *gorm.DB
}

func NewArticleRepository(db *gorm.DB) *ArticleRepository {
	return &ArticleRepository{db: db}
}

func (r *ArticleRepository) CreateArticle(article *entity.Article) error {
	err := r.db.Create(article).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *ArticleRepository) GetAllArticle() ([]entity.Article, error) {
	var articles []entity.Article
	err := r.db.Find(&articles).Error
	if err != nil {
		return nil, err
	}
	return articles, nil
}
