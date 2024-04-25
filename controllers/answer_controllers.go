package controllers

import (
	"kelasbeta/finalproject/models"
	"kelasbeta/finalproject/utils"

	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func RouteAnswers(app *fiber.App) {
	questionGroup := app.Group("/answer")
	questionGroup.Get("/", GetAnswerList)
	questionGroup.Post("/", InsertAnswer)
}

func GetAnswerList(c *fiber.Ctx) error {
	answerData, err := utils.GetAnswerList()
	if err != nil {
		logrus.Error("Error on get cars list: ", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(
			map[string]any{
				"message": "Server Error",
			},
		)
	}
	return c.Status(fiber.StatusOK).JSON(
		map[string]any{
			"data":    answerData,
			"message": "Success",
		},
	)
}

func InsertAnswer(c *fiber.Ctx) error {
	type AddQuestionRequest struct {
		Opsi       string `json:"opsi" valid:"required,type(string)"`
		Jawaban    string `json:"jawaban" valid:"required,type(string)"`
		Score      int    `json:"score" valid:"required"`
		QuestionID uint   `json:"question_id" valid:"required"`
	}

	req := new(AddQuestionRequest)

	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(map[string]interface{}{
				"message": "Body Not Valid",
			})
	}

	if _, err := govalidator.ValidateStruct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{
			"message": err.Error(),
		})
	}

	answer := models.Answer{
		Opsi:       req.Opsi,
		Jawaban:    req.Jawaban,
		Score:      req.Score,
		QuestionID: req.QuestionID,
	}

	insertedAnswer, errCreateAnswer := utils.InsertAnswer(answer)

	if errCreateAnswer != nil {
		logrus.Printf("Error creating question: %s\n", errCreateAnswer.Error())
		return c.Status(fiber.StatusInternalServerError).
			JSON(map[string]interface{}{
				"message": "Server Error",
			})
	}

	return c.Status(fiber.StatusOK).JSON(map[string]interface{}{
		"message": "Success Insert Data",
		"answer":  insertedAnswer,
	})
}
