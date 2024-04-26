package models

import "gorm.io/gorm"

type Question struct {
	Model
	Question   string   `gorm:"not null" json:"question"`
	CategoryID uint     `gorm:"not null" json:"category_id"`
	Category   Category `json:"category,omitempty"`
	Answers    []Answer `json:"answers,omitempty"`
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

func (cr *Question) GetByID(db *gorm.DB) (Question, error) {
	res := Question{}

	err := db.
		Model(Question{}).
		Preload("Category").
		Preload("Answers").
		Where("id = ?", cr.Model.ID).
		Take(&res).
		Error

	if err != nil {
		return Question{}, err
	}

	return res, nil
}

func (cr *Question) GetAll(db *gorm.DB) ([]Question, error) {
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

func (cr *Question) UpdateByID(db *gorm.DB) error {
	err := db.
		Model(Question{}).
		Select("question", "category_id").
		Where("id = ?", cr.Model.ID).
		Updates(map[string]any{
			"question":      cr.Question,
			"category_id": cr.CategoryID,
		}).
		Error

	if err != nil {
		return err
	}

	return nil
}

func (cr *Question) DeleteByID(db *gorm.DB) error {
	err := db.
		Model(Question{}).
		Where("id = ?", cr.Model.ID).
		Delete(&cr).
		Error

	if err != nil {
		return err
	}

	return nil
}
