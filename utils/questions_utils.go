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
