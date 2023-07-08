package service

import (
	"phynatureApi/entity"
	"phynatureApi/repository"
)

type articleService struct {
	articleRepo repository.ArticleRepo
}

func NewArticleService(articleRepo repository.ArticleRepo) ArticleService {
	return &articleService{articleRepo: articleRepo}
}

func (s *articleService) CreateArticle(article *entity.Article) error {
	err := s.articleRepo.CreateArticle(article)
	if err != nil {
		return err
	}
	return nil
}

func (s *articleService) GetAllArticle() ([]entity.Article, error) {
	article, err := s.articleRepo.GetAllArticle()
	if err != nil {
		return nil, err
	}
	return article, nil
}
