package controllers

import (
	"FINALPROJECT/model"
	"FINALPROJECT/utils"
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
	// categoriesgroup.Put("/:id", UpdateCategoriesByID)
	// categoriesgroup.Delete("/:id", DeleteCategoriesById)
}

func CheckRole(c *fiber.Ctx) error {
	client := string(c.Request().Header.Peek("Role"))
	if client == "Admin" {
		return c.Next()
	}
	return c.Status(fiber.StatusUnauthorized).JSON(map[string]any{
		"message": "User Unauthorized",
	})
}

func InsertCategoryData(c *fiber.Ctx) error {
	type AddCategoryRequest struct {
		name string
		order uint 
	}
	req := new(AddCategoryRequest)
	
	if err := c.BodyParser(req); err != nil{
		return c.Status(fiber.StatusBadRequest).
		JSON(map[string]any{
			"message": "Body not valid",
		})
	}
	isValid, err := govalidator.ValidateStruct(req)
	if !isValid && err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(map[string]any{
			"message": err.Error(),
		})
	}

	category, errCreatecategory := utils.InsertCategoryData(model.Category{})

	if errCreatecategory != nil {
		logrus.Printf("Terjadi error : %s\n",errCreatecategory.Error())
		return c.Status(fiber.StatusInternalServerError).
			JSON(map[string]any{
				"message": "Server Error",
			})
	}
	return c.Status(fiber.StatusOK).JSON(map[string]any{
		"message": "Success Insert Data",
		"category": category,
	})
}


func GetCategoriesList(c *fiber.Ctx) error {
	catagoriesData, err := utils.GetCategoriesList()
	if err != nil {
		logrus.Error("Error on get cars list: ",
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