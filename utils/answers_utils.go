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

func GetAnswerByID(id uint) (models.Answer, error) {
	answer := models.Answer{
		Model: models.Model{
			ID: id,
		},
	}
	return answer.GetByID(config.Postgres.DB)
}

func UpdateAnswerByID(id uint, answerData models.Answer) error {
	answer := models.Answer{
		Model: models.Model{
			ID: id,
		},
		Opsi:       answerData.Opsi,
		Jawaban:    answerData.Jawaban,
		Score:      answerData.Score,
		QuestionID: answerData.QuestionID,
	}

	return answer.UpdateByID(config.Postgres.DB)
}

func DeleteAnswerByID(id uint) error {
	question := models.Answer{
		Model: models.Model{
			ID: id,
		},
	}
	return question.DeleteByID(config.Postgres.DB)
}
