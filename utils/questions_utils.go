package utils

import (
	"kelasbeta/finalproject/config"
	"kelasbeta/finalproject/models"
	"time"
)

func GetQuestionsByIDList(ids []uint) ([]models.Question, error) {
	var questions []models.Question
	err := config.Postgres.DB.Where("id IN ?", ids).Find(&questions).Error
	if err != nil {
		return nil, err
	}
	return questions, nil
}

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
	return question, question.GetByID(config.Postgres.DB)
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
