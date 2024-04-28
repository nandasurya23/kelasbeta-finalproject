package config_test

import (
	"kelasbeta/finalproject/config"
	"testing"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func Init() {
	err := godotenv.Load("../.env")
	if err != nil {
		logrus.Println(".env not found, using global variable")
	}
}

func TestConnection(t *testing.T) {
	Init()
	config.OpenDB()
}
