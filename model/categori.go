package model

import (
	"gorm.io/gorm"
)

type Category struct {
	Model
	name  string `gorm:"not null"`
	order uint   `gorm:"not null"`
}

func (cs *Category) Create(db *gorm.DB) error {
	err := db.
		Model(Category{}).
		Create(&cs).
		Error

		if err != nil {
			return err
		}

		return nil
}

func (cs *Category) GetByID(db *gorm.DB) (Category, error) {
	res := Category{}

	err := db.
		Model(Category{}).
		Where("id = ?", cs.Model.ID).
		Take(&res).
		Error

	if err != nil {
		return Category{}, err
	}

	return res, nil
}

func (cs *Category) GetAll(db *gorm.DB) ([]Category, error) {
	res := []Category{}

	err := db.
		Model(Category{}).
		Find(&res).
		Error

	if err != nil {
		return []Category{}, err
	}

	return res, nil
}

func (cs *Category) UpdateOneByID(db *gorm.DB) error {
	err := db.
		Model(Category{}).
		Select("name", "order",).
		Where("id = ?", cs.Model.ID).
		Updates(map[string]any{
			"name":          cs.name,
			"order":          cs.order,
		}).
		Error

	if err != nil {
		return err
	}

	return nil
}

func (cs *Category) DeleteByID(db *gorm.DB) error {
	err := db.
		Model(Category{}).
		Where("id = ?", cs.Model.ID).
		Delete(&cs).
		Error

	if err != nil {
		return err
	}

	return nil
}
