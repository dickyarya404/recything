package entity

type Article struct {
	ID          string
	Title       string
	Description string
	Category    []string
	Thumbnail   string
}

type ArticleTrashCategory struct {
	Categeory string
}
