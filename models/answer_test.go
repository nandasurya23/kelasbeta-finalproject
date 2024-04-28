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

func InitAnswer() {
	err := godotenv.Load("../.env")
	if err != nil {
		logrus.Println(".env not found, using global variable")
	}
}

func TestCreateDataAnswerSuccess( *testing.T)  {
	InitAnswer()

	config.OpenDB()
	
	data := models.Answer{
		Opsi: "A",
		Jawaban: "Contoh",
		Score: 70,
		QuestionID: 1,
	}
	utils.InsertAnswer(data) 
}

func TestGetByIDAnswerSuccess(t *testing.T) {
	InitAnswer()

	answerData := models.Answer{
		Opsi: "A",
		Jawaban: "Contoh",
		Score: 70,
		QuestionID: 1,
	}

	err := answerData.Create(config.Postgres.DB)
	assert.Nil(t, err)

	Data := models.Answer{
		Model: models.Model{
			ID: 2,
		},
	}

	data, err := Data.GetByID(config.Postgres.DB)
	assert.Nil(t, err)

	fmt.Println(data)
}



func TestGetAllAnswerSuccess(t *testing.T) {
	InitAnswer()

	answerData := models.Answer{
		Opsi: "A",
		Jawaban: "Contoh",
		Score: 70,
		QuestionID: 1,
	}

	err := answerData.Create(config.Postgres.DB)
	assert.Nil(t, err)

	data, err := answerData.GetAll(config.Postgres.DB)
	assert.Nil(t, err)
	assert.GreaterOrEqual(t, len(data), 1)

	fmt.Println(data)
	
}

func TestUpdateDataAnswerSuccess( *testing.T)  {
	InitAnswer()

	config.OpenDB()
	
	data := models.Answer{
		Opsi: "b",
		Jawaban: "Contoh di update",
		Score: 2,
		QuestionID: 1,
	}
	utils.UpdateAnswerByID(2, data)
}

func TestDeleteDataAnswerSuccess( *testing.T)  {
	InitAnswer()

	config.OpenDB()

	utils.DeleteAnswerByID(1)
}