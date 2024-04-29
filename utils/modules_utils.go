package utils

import (
	"kelasbeta/finalproject/config"
	"kelasbeta/finalproject/mockstruct"
	"kelasbeta/finalproject/models"
)

func InsertModuleData(data models.Module) (models.Module, error) {
	err := data.Insert(config.Postgres.DB)
	return data, err
}

func GetModulesList() ([]models.Module, error) {
	var modules models.Module
	return modules.GetAll(config.Postgres.DB)
}

func GetModuleByID(id uint) (res mockstruct.Module, err error) {
	module := models.Module{
		Model: models.Model{
			ID: id,
		},
	}

	err = module.GetByID(config.Postgres.DB)

	if err != nil {
		return
	}

	res.CreatedAt = module.CreatedAt
	res.UpdatedAt = module.UpdatedAt
	res.ID = module.ID
	res.Identifier = module.Identifier
	res.Name = module.Name
	res.Questions = []models.Question{}

	for _, k := range module.QuestionIDS {
		qst := models.Question{
			Model: models.Model{
				ID: uint(k),
			},
		}
		err = module.GetByID(config.Postgres.DB)
		if err == nil {
			res.Questions = append(res.Questions, qst)
		}
	}
	return res, err
}

func UpdateModuleByID(id uint, moduleData models.Module) error {
	return moduleData.UpdateOneByID(config.Postgres.DB)
}

func DeleteModuleByID(id uint) error {
	module := models.Module{
		Model: models.Model{
			ID: id,
		},
	}
	return module.DeleteByID(config.Postgres.DB)
}
