package entity

type ArticleRepositoryInterface interface {
	FindAll() []Article
	FindByID(ID string) (Article, error)
}

type ArticleServiceInterface interface {
	FindAll() []Article
	FindByID(ID string) (Article, error)
}
