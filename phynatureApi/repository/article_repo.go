package repository

import "phynatureApi/entity"

type ArticleRepo interface {
	CreateArticle(article *entity.Article) error
	GetAllArticle() ([]entity.Article, error)
}
