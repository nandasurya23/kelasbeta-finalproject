package utils

import (
	"kelasbeta/finalproject/config"
	"kelasbeta/finalproject/models"
	"time"
)

func GetQuestionList() ([]models.Question, error) {
	var question models.Question
	return question.GetAll(config.Postgres.DB)
}

func InsertQuestionData(data models.Question) (models.Question, error) {
	data.CreatedAt = time.Now()
	data.UpdatedAt = time.Now()
	err := data.Create(config.Postgres.DB)

	return data, err
}

func GetQuestionByID(id uint) (models.Question, error) {
	question := models.Question{
		Model: models.Model{
			ID: id,
		},
	}
	return question.GetByID(config.Postgres.DB)
}

func UpdateQuestionByID(id uint, questionData models.Question) error {
	question := models.Question{
		Model: models.Model{
			ID: id,
		},
		Question:   questionData.Question,
		CategoryID: questionData.CategoryID,
	}

	return question.UpdateByID(config.Postgres.DB)
}

func DeleteQuestionByID(id uint) error {
	question := models.Question{
		Model: models.Model{
			ID: id,
		},
	}
	return question.DeleteByID(config.Postgres.DB)
}
