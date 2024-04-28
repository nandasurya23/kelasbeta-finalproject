package models_test

import (
	"fmt"
	"kelasbeta/finalproject/config"
	"kelasbeta/finalproject/models"
	"kelasbeta/finalproject/utils"
	"testing"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func InitCategories() {
	err := godotenv.Load("../.env")
	if err != nil {
		logrus.Println(".env not found, using global variable")
	}
	config.OpenDB()
}

func TestCreateDataCategoriesSuccess( *testing.T)  {
	InitCategories()

	config.OpenDB()
	
	data := models.Category{
		Name: "kategori A",
		Order: 1,
	}
	utils.InsertCategoryData(data)
}

func TestGetByIDCategoriesSuccess(t *testing.T) {
	InitCategories()

	categoriData := models.Category{
		Name: "kategori dua",
		Order: 2,

		
	}

	err := categoriData.Create(config.Postgres.DB)
	assert.Nil(t, err)

	Data := models.Category{
		Model: models.Model{
			ID: 2,
		},
	}

	data, err := Data.GetByID(config.Postgres.DB)
	assert.Nil(t, err)

	fmt.Println(data)
}



func TestGetAllCategoriesSuccess(t *testing.T) {
	InitCategories()

	categoriData := models.Category{
		Name: "kategori dua belas",
		Order: 12,

		
	}

	err := categoriData.Create(config.Postgres.DB)
	assert.Nil(t, err)

	data, err := categoriData.GetAll(config.Postgres.DB)
	assert.Nil(t, err)
	assert.GreaterOrEqual(t, len(data), 1)

	fmt.Println(data)
	
}

func TestUpdateDataCategoriesSuccess( *testing.T)  {
	InitCategories()

	config.OpenDB()
	
	data := models.Category{
		Name: "kategori update",
		Order: 3,
	}
	utils.UpdateCategoriesByID(1, data)
}

func TestDeleteDataCategoriesSuccess( *testing.T)  {
	InitCategories()

	config.OpenDB()

	utils.DeleteCategoriesByID(1)
}