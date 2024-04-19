package models

type Question struct {
	ID           uint   `db:"id"`
	Question     string `db:"question"`
	CategoriesID uint   `db:"categories_id"`
}
