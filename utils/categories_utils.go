package utils

import (
	"kelasbeta/finalproject/config"
	"kelasbeta/finalproject/models"
	"time"
)

func InsertCategoryData(data models.Category) (models.Category, error) {
	data.CreatedAt = time.Now()
	data.UpdatedAt = time.Now()
	err := data.Create(config.Postgres.DB)
	return data, err
}

func GetCategoriesList() ([]models.Category, error) {
	var Categories models.Category
	return Categories.GetAll(config.Postgres.DB)
}

func GetCategoriesByID(id uint) (models.Category, error) {
	categories := models.Category{
		Model: models.Model{
			ID: id,
		},
	}
	return categories.GetByID(config.Postgres.DB)
}

func UpdateCategoriesByID(id uint, categoriesData models.Category) error {
	answer := models.Category{
		Model: models.Model{
			ID: id,
		},
		Name:       categoriesData.Name,
		Order:    categoriesData.Order,
	}

	return answer.UpdateOneByID(config.Postgres.DB)
}


func DeleteCategoriesByID(id uint) error {
	categori := models.Category{
		Model: models.Model{
			ID: id,
		},
	}
	return categori.DeleteByID(config.Postgres.DB)
}