package controllers

import (
	"kelasbeta/finalproject/utils"
	"log"

	"github.com/gofiber/fiber/v2"
)

func ExamRoutes(app *fiber.App) {
	examRoute := app.Group("exam")
	examRoute.Get("/question/:identifier", GetExamQuestionsByID)
}

func GetExamQuestionsByID(c *fiber.Ctx) error {
	res, err := utils.GetExamQuestions(c.Params("identifier"))
	if err != nil {
		if err.Error() == "record not found" {
			return c.Status(fiber.StatusNotFound).JSON(map[string]any{
				"message": "ID Not Found",
			})
		}
		log.Printf("An Error occured on GetCategorys, err:%s", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(map[string]any{
			"message": "Server Error",
		})
	}

	return c.Status(fiber.StatusOK).JSON(map[string]any{
		"data":    res,
		"message": "Success",
	})
}
