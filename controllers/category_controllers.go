package controllers

import (
	"kelasbeta/finalproject/models"
	"kelasbeta/finalproject/utils"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func RouteCategories(app *fiber.App)  {
	categoriesgroup := app.Group("/categories",)
	categoriesgroup.Post("/", InsertCategoryData)
	categoriesgroup.Get("/", GetCategoriesList)
	categoriesgroup.Get("/:id", GetCategoriesByID)
	categoriesgroup.Put("/:id", UpdatecategoriesByID)
	categoriesgroup.Delete("/:id", DeleteCategoriesByID)
}

// func CheckRole(c *fiber.Ctx) error {
// 	client := string(c.Request().Header.Peek("Role"))
// 	if client == "Admin" {
// 		return c.Next()
// 	}
// 	return c.Status(fiber.StatusUnauthorized).JSON(map[string]any{
// 		"message": "User Unauthorized",
// 	})
// }

func InsertCategoryData(c *fiber.Ctx) error {
	type AddCategoryRequest struct {
		Name string `json:"name" valid:"required,type(string)"` 
		Order int `json:"order" valid:"required"`
	}
	req := new(AddCategoryRequest)
	
	if err := c.BodyParser(req); err != nil{
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

	categori := models.Category{
		Name:  req.Name,
		Order: req.Order,
	}

	InsertedCategoryData, errCreateCategori:= utils.InsertCategoryData(categori)

	if errCreateCategori != nil {
		logrus.Printf("Error creating category: %s\n", errCreateCategori.Error())
		return c.Status(fiber.StatusInternalServerError).
			JSON(map[string]interface{}{
				"message": "Server Error",
			})
	}

	return c.Status(fiber.StatusCreated).JSON(map[string]interface{}{
		"message": "Success Insert Data",
		"category": InsertedCategoryData,
	})
}


func GetCategoriesList(c *fiber.Ctx) error {
	catagoriesData, err := utils.GetCategoriesList()
	if err != nil {
		logrus.Error("Error on get categories list: ",
		err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(
			map[string]any{
				"message": "Server Error",
			},
		)
	}
	return c.Status(fiber.StatusOK).JSON(
		map[string]any{
			"data":    catagoriesData,
			"message": "Success",
		},
	)
}


func GetCategoriesByID(c *fiber.Ctx) error {
	categoryId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			map[string]any{
				"message": "ID not valid",
			},
		)
	}
	categoriesData, err := utils.GetCategoriesByID(uint(categoryId))
	if err != nil {
		if err.Error() == "record not found"{
			return c.Status(fiber.StatusNotFound).JSON(
				map[string]any{
					"message": "ID not found",
				},
			)
		}
		logrus.Error("Error on get categories data: ", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(
			map[string]any{
				"message": "Server Error",
			},
		)
	}
	return c.Status(fiber.StatusOK).JSON(
		map[string]any{
			"data":    categoriesData,
			"message": "Success",
		},
	)
}


func UpdatecategoriesByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{
			"message": "ID not valid",
		})
	}

	var categoriesData models.Category
	if err := c.BodyParser(&categoriesData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{
			"message": "Invalid request body",
		})
	}
if err := utils.UpdateCategoriesByID(uint(id), categoriesData); err != nil {
		logrus.Errorf("Error updating categories: %s", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(map[string]interface{}{
			"message": "Failed to update categories",
		})
	}

	return c.Status(fiber.StatusOK).JSON(map[string]interface{}{
		"message": "categories updated successfully",
	})
}

func DeleteCategoriesByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "ID not valid",
		})
	}

	err = utils.DeleteCategoriesByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to delete categories",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "categories deleted successfully",
	})
}