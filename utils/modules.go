package utils

import (
	"FINALPROJECT/config"
	"FINALPROJECT/model"
	"time"
)

func InsertModuleData(data model.Module) (model.Module, error){
	data.CreatedAt = time.Now()
	data.UpdatedAt = time.Now()
	err := data.Create(config.Postgres.DB)
	return data, err

}

func GetModulesList() ([]model.Module, error) {
	var modules model.Module
	return modules.GetAll(config.Postgres.DB)
}


func GetModulesByID(id uint) (model.Module, error) {
	module := model.Module{
		Model: model.Model{
			ID: id,
		},
	}
	return module.GetByID(config.Postgres.DB)
}

func UpdateModulesByID(id uint, moduleData model.Module) error {
	answer := model.Module{
		Model: model.Model{
			ID: id,
		},
		Identifier:       moduleData.Identifier,
		Name:    moduleData.Name,
		QuestionIDS:    moduleData.QuestionIDS,
	}

	return answer.UpdateOneByID(config.Postgres.DB)
}


func DeleteModulessByID(id uint) error {
	categori := model.Module{
		Model: model.Model{
			ID: id,
		},
	}
	return categori.DeleteByID(config.Postgres.DB)
}
