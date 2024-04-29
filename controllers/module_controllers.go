package controllers

import (
	"kelasbeta/finalproject/models"
	"kelasbeta/finalproject/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func RouteModules(app *fiber.App) {
	modulesgroup := app.Group("/modules")
	modulesgroup.Post("/", utils.CheckRole, InsertModulesData)
	modulesgroup.Get("/", utils.CheckRole, GetModulesList)
	modulesgroup.Get("/:id", utils.CheckRole, GetModuleByID)
	modulesgroup.Put("/:id", utils.CheckRole, UpdateModuleByID)
	modulesgroup.Delete("/:id", utils.CheckRole, DeleteModuleByID)
}

func InsertModulesData(c *fiber.Ctx) error {
	var req models.Module
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}
	insertedModuleData, err := utils.InsertModuleData(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Server Error",
		})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Success Insert Data",
		"module":  insertedModuleData,
	})
}

func GetModulesList(c *fiber.Ctx) error {
	moduleData, err := utils.GetModulesList()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Server Error",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    moduleData,
		"message": "Success",
	})
}

func GetModuleByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "ID not valid",
		})
	}
	moduleData, err := utils.GetModuleByID(uint(id))
	if err != nil {
		if err.Error() == "record not found" {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "ID not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Server Error",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    moduleData,
		"message": "Success",
	})
}

func UpdateModuleByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "ID not valid",
		})
	}

	var moduleData models.Module
	if err := c.BodyParser(&moduleData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}
	if err := utils.UpdateModuleByID(uint(id), moduleData); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to update module",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Module updated successfully",
	})
}

func DeleteModuleByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "ID not valid",
		})
	}
	if err := utils.DeleteModuleByID(uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to delete module",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Module deleted successfully",
	})
}
