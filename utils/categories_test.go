package utils_test

import (
	"FINALPROJECT/config"
	"FINALPROJECT/model"
	"FINALPROJECT/utils"
	"context"
	"fmt"
	"testing"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func Init() {
	err := godotenv.Load("../.env")
	if err != nil {
		logrus.Println(".env not found, using global variable")
	}
}


func TestCreateData(t *testing.T)  {
	Init()
	conn, err := config.OpenConn()
	assert.Nil(t, err)

	bdy := model.Category{
		name: "TNI",
		order: 3,
	}
	err = utils.InsertCategoryData(conn, bdy, context.TODO())
	assert.Nil(t, err)
}


func TestGetByID(t *testing.T)  {
	Init()

	conn, err := config.OpenConn()
	assert.Nil(t, err)
	ctx := context.TODO()

	res, err := utils.GetCategoriesByID(conn, "12345", ctx)
	assert.Nil(t, err)

	fmt.Println(res)
}