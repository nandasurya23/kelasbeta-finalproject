package models

import "gorm.io/gorm"

type Answer struct {
	Model
	Opsi       string `gorm:"not null" json:"opsi"`
	Jawaban    string `gorm:"not null" json:"jawaban"`
	Score      int    `gorm:"not null" json:"score"`
	QuestionID uint   `json:"question_id" gorm:"not null;foreignKey:QuestionID"`
}

func (cr *Answer) Create(db *gorm.DB) error {
	err := db.
		Model(Answer{}).
		Create(&cr).
		Error

	if err != nil {
		return err
	}

	return nil
}

func (cr *Answer) GetByID(db *gorm.DB) (Answer, error) {
	res := Answer{}

	err := db.
		Model(Answer{}).
		Where("id = ?", cr.Model.ID).
		Take(&res).
		Error

	if err != nil {
		return Answer{}, err
	}

	return res, nil
}

func (cr *Answer) GetAll(db *gorm.DB) ([]Answer, error) {
	res := []Answer{}

	err := db.
		Model(Answer{}).
		Find(&res).
		Error

	if err != nil {
		return []Answer{}, err
	}

	return res, nil
}

func (cr *Answer) UpdateByID(db *gorm.DB) error {
	err := db.
		Model(Answer{}).
		Select("opsi", "jawaban", "score", "question_id").
		Where("id = ?", cr.Model.ID).
		Updates(map[string]any{
			"opsi":        cr.Opsi,
			"jawaban":     cr.Jawaban,
			"score":       cr.Score,
			"question_id": cr.QuestionID,
		}).
		Error

	if err != nil {
		return err
	}

	return nil
}

func (cr *Answer) DeleteByID(db *gorm.DB) error {
	err := db.
		Model(Answer{}).
		Where("id = ?", cr.Model.ID).
		Delete(&cr).
		Error

	if err != nil {
		return err
	}

	return nil
}
