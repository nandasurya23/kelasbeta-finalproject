package controllers

import (
	"kelasbeta/finalproject/models"
	"kelasbeta/finalproject/utils"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v2"
	"github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

func RouteModules(app *fiber.App) {
	modulesgroup := app.Group("/modules",)
	modulesgroup.Post("/", InsertModulesData)
	modulesgroup.Get("/", GetModulesList)
	modulesgroup.Get("/:id", GetModulesByID)
	modulesgroup.Put("/:id", UpdateModulesByID)
	modulesgroup.Delete("/:id", DeleteModulessByID)
}

func InsertModulesData(c *fiber.Ctx) error {
	type AddModuleRequest struct{
		Identifier   string `json:"identifier" valid:"required, type(string)"`
		Name         string `json:"name" valid:"required, type(string)"`
		QuestionIDS pq.Int64Array `json:"question_ids" valid:"required"`
	}
	req := new(AddModuleRequest)

	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).
		JSON(map[string]interface{}{
			"message": "Body not valid",
		})
	}
if _, err := govalidator.ValidateStruct(req); err != nil {
	return c.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{
		"message": err.Error(),
	})
}

module := models.Module{
	Identifier: req.Identifier,
	Name: req.Name,
	QuestionIDS: req.QuestionIDS,
}

InsertedModulesData, errCreateModule:= utils.InsertModuleData(module)


if errCreateModule != nil {
	logrus.Printf("Error creating modules: %s\n", errCreateModule.Error())
	return c.Status(fiber.StatusInternalServerError).
		JSON(map[string]interface{}{
			"message": "Server Error",
		})
}

return c.Status(fiber.StatusCreated).JSON(map[string]interface{}{
	"message": "Success Insert Data",
	"module": InsertedModulesData,
})
}

func GetModulesList(c *fiber.Ctx) error {
	moduleData, err := utils.GetModulesList()
	if err != nil {
		logrus.Error("Error on get modules list: ",
		err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(
			map[string]any{
				"message": "Server Error",
			},
		)
	}
	return c.Status(fiber.StatusOK).JSON(
		map[string]any{
			"data":    moduleData,
			"message": "Success",
		},
	)
}


func GetModulesByID(c *fiber.Ctx) error {
	moduleId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			map[string]any{
				"message": "ID not valid",
			},
		)
	}
	modulesData, err := utils.GetModulesByID(uint(moduleId))
	if err != nil {
		if err.Error() == "record not found"{
			return c.Status(fiber.StatusNotFound).JSON(
				map[string]any{
					"message": "ID not found",
				},
			)
		}
		logrus.Error("Error on get modules data: ", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(
			map[string]any{
				"message": "Server Error",
			},
		)
	}
	return c.Status(fiber.StatusOK).JSON(
		map[string]any{
			"data":    modulesData,
			"message": "Success",
		},
	)
}


func UpdateModulesByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{
			"message": "ID not valid",
		})
	}

	var modulesData models.Module
	if err := c.BodyParser(&modulesData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{
			"message": "Invalid request body",
		})
	}
if err := utils.UpdateModulesByID(uint(id), modulesData); err != nil {
		logrus.Errorf("Error updating modules: %s", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(map[string]interface{}{
			"message": "Failed to update modules",
		})
	}

	return c.Status(fiber.StatusOK).JSON(map[string]interface{}{
		"message": "modules updated successfully",
	})
}

func DeleteModulessByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "ID not valid",
		})
	}

	err = utils.DeleteModulesByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to delete modules",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "modules deleted successfully",
	})
}