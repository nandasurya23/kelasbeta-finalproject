package controllers

import (
	"kelasbeta/finalproject/models"
	"kelasbeta/finalproject/utils"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func RouteQuestions(app *fiber.App) {
	questionGroup := app.Group("/question")
	questionGroup.Get("/", GetQuestionList)
	questionGroup.Get("/:id", GetQuestionByID)
	questionGroup.Post("/",utils.CheckRole, InsertQuestionData)
	questionGroup.Put("/:id",utils.CheckRole,  UpdateQuestionByID)
	questionGroup.Delete("/:id",utils.CheckRole,  DeleteQuestionByID)
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
		Question   string `json:"question" valid:"required,type(string)"`
		CategoryID uint   `json:"category_id" valid:"required"`
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
		Question:   req.Question,
		CategoryID: req.CategoryID,
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

func GetQuestionByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{
			"message": "ID not valid",
		})
	}

	questionData, err := utils.GetQuestionByID(uint(id))
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
		"data":    questionData,
		"message": "Success",
	})
}

func UpdateQuestionByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{
			"message": "ID not valid",
		})
	}

	var questionData models.Question
	if err := c.BodyParser(&questionData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{
			"message": "Invalid request body",
		})
	}

	if err := utils.UpdateQuestionByID(uint(id), questionData); err != nil {
		logrus.Errorf("Error updating question: %s", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(map[string]interface{}{
			"message": "Failed to update question",
		})
	}

	return c.Status(fiber.StatusOK).JSON(map[string]interface{}{
		"message": "Question updated successfully",
	})
}

func DeleteQuestionByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "ID not valid",
		})
	}

	err = utils.DeleteQuestionByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to delete question",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Question deleted successfully",
	})
}
