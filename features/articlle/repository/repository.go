package repository

import "gorm.io/gorm"

type ArticleRepository struct {
	db            *gorm.DB
	trashcategory trashcategory.TrashCategoryInterface
}
