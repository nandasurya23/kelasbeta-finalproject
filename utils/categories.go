package utils

import (
	"FINALPROJECT/config"
	"FINALPROJECT/model"
	"time"
)


func InsertCategoryData(data model.category) (model.category, error) {
	data.CreatedAt = time.Now()
	data.UpdatedAt = time.Now()
	err := data.Create(config.)
		//bagaimana di postgres DB???
	return data, err
}

func GetCategoriesList() ([]model.Category, error) {
	var Categories model.category
	return Categories.GetAll(config.)
	//bagaimana di postgres DB???
}

func GetCategoriesByID(id uint) (model.Car, error) {
	categories := model.Category{
		Model: model.Model{
			ID: id,
		},
	}
	return categories.GetByID(config.)
	//bagaimana di postgres DB???
}