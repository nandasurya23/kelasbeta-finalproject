package models

import "gorm.io/gorm"

type Question struct {
	Model
	Question     string `gorm:"not null" json:"question"`
	CategoriesID uint   `gorm:"not null" json:"categories_id"`
}

func (cr *Question) Create(db *gorm.DB) error {
	err := db.
		Model(Question{}).
		Create(&cr).
		Error

	if err != nil {
		return err
	}

	return nil
}

func (cr *Question) GetAll(db *gorm.DB) ([]Question,error) {
	res := []Question{}

	err := db.
		Model(Question{}).
		Find(&res).
		Error

	if err != nil {
		return []Question{}, err
	}

	return res, nil
}
