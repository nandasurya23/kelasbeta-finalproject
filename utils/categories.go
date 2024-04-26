package utils

import (
	"FINALPROJECT/config"
	"FINALPROJECT/model"
	"time"
)


func InsertCategoryData(data model.Category) (model.Category, error) {
	data.CreatedAt = time.Now()
	data.UpdatedAt = time.Now()
	err := data.Create(config.Postgres.DB)
	return data, err
}

func GetCategoriesList() ([]model.Category, error) {
	var Categories model.Category
	return Categories.GetAll(config.Postgres.DB)
}

func GetCategoriesByID(id uint) (model.Category, error) {
	categories := model.Category{
		Model: model.Model{
			ID: id,
		},
	}
	return categories.GetByID(config.Postgres.DB)
}

func UpdateCategoriesByID(id uint, categoriesData model.Category) error {
	answer := model.Category{
		Model: model.Model{
			ID: id,
		},
		Name:       categoriesData.Name,
		Order:    categoriesData.Order,
	}

	return answer.UpdateOneByID(config.Postgres.DB)
}


func DeleteCategoriesByID(id uint) error {
	categori := model.Category{
		Model: model.Model{
			ID: id,
		},
	}
	return categori.DeleteByID(config.Postgres.DB)
}
