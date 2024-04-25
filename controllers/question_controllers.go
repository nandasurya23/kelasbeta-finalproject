package controllers

import (
	"kelasbeta/finalproject/models"
	"kelasbeta/finalproject/utils"

	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func RouteQuestions(app *fiber.App) {
	questionGroup := app.Group("/question")
	questionGroup.Get("/", GetQuestionList)
	questionGroup.Post("/", InsertQuestionData)
}

func GetQuestionList(c *fiber.Ctx) error {
	questionData, err := utils.GetQuestionList()
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
			"data":    questionData,
			"message": "Success",
		},
	)
}

func InsertQuestionData(c *fiber.Ctx) error {
	type AddQuestionRequest struct {
		Question     string `json:"question" valid:"required,type(string)"`
		CategoriesID uint   `json:"categories_id" valid:"required"`
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

	question := models.Question{
		Question:     req.Question,
		CategoriesID: req.CategoriesID,
	}

	insertedQuestion, errCreateQuestion := utils.InsertQuestionData(question)

	if errCreateQuestion != nil {
		logrus.Printf("Error creating question: %s\n", errCreateQuestion.Error())
		return c.Status(fiber.StatusInternalServerError).
			JSON(map[string]interface{}{
				"message": "Server Error",
			})
	}

	return c.Status(fiber.StatusOK).JSON(map[string]interface{}{
		"message":  "Success Insert Data",
		"question": insertedQuestion,
	})
}
