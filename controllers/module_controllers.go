package controllers

import (
	"FINALPROJECT/model"
	"FINALPROJECT/utils"

	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v2"
	"github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

func RouteModules(app *fiber.App) {
	modulesgroup := app.Group("/modules",)
	modulesgroup.Post("/", InsertModulesData)
	// modulesgroup.Get("/", GetModulesList)
	// modulesgroup.Get("/:id", GetModulesByID)
	// modulesgroup.Put("/:id", UpdateModulesByID)
	// modulesgroup.Delete("/:id", DeleteModulesById)
}


func InsertModulesData(c *fiber.Ctx) error {
	type AddModuleRequest struct{
		Identifier   string `json:"identifier" valid:"required, type(string)"`
		Name         string `json:"name" valid:"required, type(string)"`
		QuestionIDS pq.Int64Array `json:"question_ids" valid:"required, type(string)"`
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

module := model.Module{
	Identifier: req.Identifier,
	Name: req.Name,
	QuestionIDS: req.QuestionIDS,
}

InsertedModulesData, errCreateModule:= utils.InsertModuleData(module)


if errCreateModule != nil {
	logrus.Printf("Error creating categori: %s\n", errCreateModule.Error())
	return c.Status(fiber.StatusInternalServerError).
		JSON(map[string]interface{}{
			"message": "Server Error",
		})
}

return c.Status(fiber.StatusCreated).JSON(map[string]interface{}{
	"message": "Success Insert Data",
	"category": InsertedModulesData,
})
}


