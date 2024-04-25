package controllers

import (
	"kelasbeta/finalproject/models"
	"kelasbeta/finalproject/utils"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func RouteAnswers(app *fiber.App) {
	answerGroup := app.Group("/answer")
	answerGroup.Get("/", GetAnswerList)
	answerGroup.Get("/:id", GetAnswerByID)
	answerGroup.Post("/", InsertAnswer)
	answerGroup.Put("/:id", UpdateAnswerByID)
	answerGroup.Delete("/:id", DeleteAnswerByID)
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
	type AddAnswerRequest struct {
		Opsi       string `json:"opsi" valid:"required,type(string)"`
		Jawaban    string `json:"jawaban" valid:"required,type(string)"`
		Score      int    `json:"score" valid:"required"`
		QuestionID uint   `json:"question_id" valid:"required"`
	}

	req := new(AddAnswerRequest)

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

func GetAnswerByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{
			"message": "ID not valid",
		})
	}

	answerData, err := utils.GetAnswerByID(uint(id))
	if err != nil {
		if err.Error() == "record not found" {
			return c.Status(fiber.StatusNotFound).JSON(map[string]interface{}{
				"message": "ID not found",
			})
		}
		logrus.Error("Error on get question data: ", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(map[string]interface{}{
			"message": "Server Error",
		})
	}

	return c.Status(fiber.StatusOK).JSON(map[string]interface{}{
		"data":    answerData,
		"message": "Success",
	})
}

func UpdateAnswerByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{
			"message": "ID not valid",
		})
	}

	var answerData models.Answer
	if err := c.BodyParser(&answerData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{
			"message": "Invalid request body",
		})
	}

	if err := utils.UpdateAnswerByID(uint(id), answerData); err != nil {
		logrus.Errorf("Error updating question: %s", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(map[string]interface{}{
			"message": "Failed to update question",
		})
	}

	return c.Status(fiber.StatusOK).JSON(map[string]interface{}{
		"message": "Answers updated successfully",
	})
}

func DeleteAnswerByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "ID not valid",
		})
	}

	err = utils.DeleteAnswerByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to delete question",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Answers deleted successfully",
	})
}
