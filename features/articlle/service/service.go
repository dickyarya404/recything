package service

import "github.com/sawalreverr/recything/features/entity"

type articleService struct {
	ArticleRepository entiti.ArticleRepositoryInterface
}

func NewArticleServicea(article entity.ArticleRepositoryInterface) entity.ArticleServiceInterface {
	return &articleService{
		ArticleRepository: article,
	}
}
