package models

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Module struct {
	Model
	Identifier   string `gorm:"not null" json:"identifier"`
	Name         string `gorm:"not null" json:"name"`
	QuestionIDS pq.Int64Array  `gorm:"type:string" json:"question_ids"`
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

func (cr *Module) GetByID(db *gorm.DB) (Module, error)  {
	res := Module{}

	err := db.
	Model(Module{}).
	Where("id = ?", cr.Model.ID).
	Take(&res).Error

	if err != nil {
		return Module{}, err
	}
	return res, nil

}

func (cr *Module) GetAll(db *gorm.DB) ([]Module,error)  {
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
	Select("Identifier", "Name", "QuestionIDS",).
	Where("id = ?", cr.Model.ID).
	Updates(map[string]any{
		"Identifier": cr.Identifier,
		"Name": cr.Name,
		"QuestionIDS": cr.QuestionIDS,
	}).
	Error
	if err != nil {
		return err
	}
	return nil
}

func (cr *Module) DeleteByID(db *gorm.DB)error  {
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