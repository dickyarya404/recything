package entity

func CategoryModelToCategory(category CategoryModel) Category {
	return ArticleTrashCategory{
		Category: category.Trashtype,
	}
}
