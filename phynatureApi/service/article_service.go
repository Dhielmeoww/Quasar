package service

import "phynatureApi/entity"

type ArticleService interface {
	CreateArticle(article *entity.Article) error
	GetAllArticle() ([]entity.Article, error)
}
