package utils

import (
	"kelasbeta/finalproject/config"
	"kelasbeta/finalproject/models"
	"time"
)

func GetAnswerList() ([]models.Answer, error) {
	var answer models.Answer
	return answer.GetAll(config.Postgres.DB)
}

func InsertAnswer(data models.Answer) (models.Answer, error) {
	data.CreatedAt = time.Now()
	data.UpdatedAt = time.Now()
	err := data.Create(config.Postgres.DB)

	return data, err
}
