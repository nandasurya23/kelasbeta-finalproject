package models_test

import (
	"fmt"
	"kelasbeta/finalproject/config"
	"kelasbeta/finalproject/models"
	"kelasbeta/finalproject/utils"
	"testing"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func InitQuestion() {
	err := godotenv.Load("../.env")
	if err != nil {
		logrus.Println(".env not found, using global variable")
	}
}

func TestCreateDataQuestionSuccess( *testing.T)  {
	InitQuestion()

	config.OpenDB()
	
	data := models.Question{
		Question: "pertanyaan pertama",
		CategoryID: 1,

	}
	utils.InsertQuestionData(data)
}

func TestGetByIDQuestionSuccess(t *testing.T) {
	InitQuestion()

	questionData := models.Question{
		Question: "pertanyaan 2",
		CategoryID: 2,

		
	}

	err := questionData.Create(config.Postgres.DB)
	assert.Nil(t, err)

	Data := models.Answer{
		Model: models.Model{
			ID: 12,
		},
	}

	data, err := Data.GetByID(config.Postgres.DB)
	assert.Nil(t, err)

	fmt.Println(data)
}



func TestGetAllQuestionSuccess(t *testing.T) {
	InitQuestion()

	questionData := models.Question{
		Question: "pertanyaan 2",
		CategoryID: 2,

		
	}

	err := questionData.Create(config.Postgres.DB)
	assert.Nil(t, err)

	data, err := questionData.GetAll(config.Postgres.DB)
	assert.Nil(t, err)
	assert.GreaterOrEqual(t, len(data), 1)

	fmt.Println(data)
	
}

func TestUpdateDataQuestionSuccess( *testing.T)  {
	InitQuestion()

	config.OpenDB()
	
	data := models.Question{
		Question: "pertanyaan 2 update",
		CategoryID: 2,
	}
	utils.UpdateQuestionByID(1, data)
}

func TestDeleteDataQuestionSuccess( *testing.T)  {
	InitQuestion()

	config.OpenDB()

	utils.DeleteQuestionByID(1)
}