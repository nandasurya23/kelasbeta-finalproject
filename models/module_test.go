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

func InitModules() {
	err := godotenv.Load(".env")
	if err != nil {
		logrus.Println(".env not found, using global variable")
	}
}

func TestCreateDataModulesSuccess( *testing.T)  {
	InitModules()

	config.OpenDB()
	
	data := models.Module{
		Identifier: "MDL-1712057343",
		Name: "module a",
		QuestionIDS:[]int64{1, 2, 3},
	}
	utils.InsertModuleData(data)
}

// func TestGetByIDModuleSuccess(t *testing.T) {
// 	InitModules()

// 	moduleData := models.Module{
// 		Identifier: "MDL-1712057343",
// 		Name: "module a",
// 		QuestionIDS:[]int64{1, 2, 3},

		
// 	}

// 	err := moduleData.Create(config.Postgres.DB)
// 	assert.Nil(t, err)

// 	Data := models.Module{
// 		Model: models.Model{
// 			ID: 12,
// 		},
// 	}

// 	data, err := Data.GetByID(config.Postgres.DB)
// 	assert.Nil(t, err)

// 	fmt.Println(data)
// }



func TestGetAllModuleSuccess(t *testing.T) {
	InitModules()

	moduleData := models.Module{
		Identifier: "MDL-1712057343",
		Name: "module a",
		QuestionIDS:[]int64{1, 2, 3},
	}

	err := moduleData.Create(config.Postgres.DB)
	assert.Nil(t, err)

	data, err := moduleData.GetAll(config.Postgres.DB)
	assert.Nil(t, err)
	assert.GreaterOrEqual(t, len(data), 1)

	fmt.Println(data)
	
}

func TestUpdateDataModulesSuccess( *testing.T)  {
	InitModules()

	config.OpenDB()
	
	data := models.Module{
		Identifier: "MDL-1712057343",
		Name: "module a updated",
		QuestionIDS:[]int64{1, 2, 3},
	}
	utils.UpdateModuleByID(2, data)
}

func TestDeleteDataModulesSuccess( *testing.T)  {
	InitModules()

	config.OpenDB()

	utils.DeleteModuleByID(2)
}