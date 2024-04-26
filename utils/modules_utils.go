package utils

import (
	"kelasbeta/finalproject/config"
	"kelasbeta/finalproject/models"
	"time"
)

func InsertModuleData(data models.Module) (models.Module, error) {
	data.CreatedAt = time.Now()
	data.UpdatedAt = time.Now()
	err := data.Create(config.Postgres.DB)
	return data, err

}

func GetModulesList() ([]models.Module, error) {
	var modules models.Module
	return modules.GetAll(config.Postgres.DB)
}

func GetModulesByID(id uint) (models.Module, error) {
	module := models.Module{
		Model: models.Model{
			ID: id,
		},
	}
	return module.GetByID(config.Postgres.DB)
}

func UpdateModulesByID(id uint, modulesData models.Module) error {
	answer := models.Module{
		Model: models.Model{
			ID: id,
		},
		Identifier:  modulesData.Identifier,
		Name:        modulesData.Name,
		QuestionIDS: modulesData.QuestionIDS,
	}

	return answer.UpdateOneByID(config.Postgres.DB)
}

func DeleteModulesByID(id uint) error {
	categori := models.Module{
		Model: models.Model{
			ID: id,
		},
	}
	return categori.DeleteByID(config.Postgres.DB)
}
