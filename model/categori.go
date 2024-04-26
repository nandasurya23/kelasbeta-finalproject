package model

import (
	"gorm.io/gorm"
)

type Category struct {
	Model
	Name  string `gorm:"not null" json:"name"`
	Order int   `gorm:"not null" json:"order"`
}

func (cr *Category) Create(db *gorm.DB) error {
	err := db.
		Model(Category{}).
		Create(&cr).
		Error

	if err != nil {
		return err
	}

	return nil
}

func (cr *Category) GetByID(db *gorm.DB) (Category, error) {
	res := Category{}

	err := db.
		Model(Category{}).
		Where("id = ?", cr.Model.ID).
		Take(&res).
		Error

	if err != nil {
		return Category{}, err
	}

	return res, nil
}

func (cr *Category) GetAll(db *gorm.DB) ([]Category, error) {
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

func (cr *Category) UpdateOneByID(db *gorm.DB) error {
	err := db.
		Model(Category{}).
		Select("Name", "Order").
		Where("id = ?", cr.Model.ID).
		Updates(map[string]any{
			"Name":  cr.Name,
			"Order": cr.Order,
		}).
		Error

	if err != nil {
		return err
	}

	return nil
}

func (cr *Category) DeleteByID(db *gorm.DB) error {
	err := db.
		Model(Category{}).
		Where("id = ?", cr.Model.ID).
		Delete(&cr).
		Error

	if err != nil {
		return err
	}

	return nil
}
