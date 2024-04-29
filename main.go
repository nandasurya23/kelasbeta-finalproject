package main

import (
	"kelasbeta/finalproject/config"
	"kelasbeta/finalproject/controllers"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func InitEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		logrus.Warn("Cannot load env file, using system env")
	}
}

func main() {
	InitEnv()
	config.OpenDB()

	app := fiber.New()

	controllers.RouteQuestions(app)
	controllers.RouteAnswers(app)
	controllers.RouteCategories(app)
	controllers.RouteModules(app)
	controllers.ExamRoutes(app)

	err := app.Listen(":3000")
	if err != nil {
		logrus.Fatal(
			"Error on Running Fiber, ",
			err.Error())
	}

}
