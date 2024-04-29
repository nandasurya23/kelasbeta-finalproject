package models

import (
	"fmt"
	"time"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Module struct {
	Model
	Identifier  string        `gorm:"not null" json:"identifier"`
	Name        string        `gorm:"not null" json:"name"`
	QuestionIDS pq.Int64Array `gorm:"type:string" json:"question_ids"`
}

func (cr *Module) Create(db *gorm.DB) error {
	err := db.
		Model(Module{}).
		Create(&cr).
		Error

	if err != nil {
		return err
	}
	return nil
}

func (cr *Module) GetByID(db *gorm.DB) error {
	return db.Model(Module{}).Where("id = ?", cr.ID).Take(&cr).Error
}

func (cr *Module) GetByIdentifier(db *gorm.DB) error {
	return db.Model(Module{}).Where("identifier = ?", cr.Identifier).Take(&cr).Error
}

func (cr *Module) Insert(db *gorm.DB) error {
	cr.Identifier = fmt.Sprintf("MDL-%d", time.Now().Unix())
	cr.CreatedAt = time.Now()
	cr.UpdatedAt = time.Now()
	return db.Model(Module{}).Create(&cr).Error
}

func (cr *Module) GetAll(db *gorm.DB) ([]Module, error) {
	res := []Module{}

	err := db.
		Model(Module{}).
		Find(&res).
		Error

	if err != nil {
		return []Module{}, err
	}
	return res, nil
}

func (cr *Module) UpdateOneByID(db *gorm.DB) error {
	err := db.
		Model(Module{}).
		Select("Identifier", "Name", "QuestionIDS").
		Where("id = ?", cr.Model.ID).
		Updates(map[string]interface{}{
			"Identifier":  cr.Identifier,
			"Name":        cr.Name,
			"QuestionIDS": cr.QuestionIDS,
		}).
		Error
	if err != nil {
		return err
	}
	return nil
}

func (cr *Module) DeleteByID(db *gorm.DB) error {
	err := db.
		Model(Module{}).
		Where("id = ?", cr.Model.ID).
		Delete(&cr).
		Error

	if err != nil {
		return err
	}

	return nil
}
