package main

import (
	"go-embedded-system/app/internal/db"
	"go-embedded-system/app/internal/handler"
	"go-embedded-system/app/internal/domain"
	"go-embedded-system/app/internal/repository"
	"go-embedded-system/app/internal/usecase"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	db.InitPostgres()

	db.DB.AutoMigrate(&domain.TemperatureData{})

	repo := repository.NewTemperatureRepository()
	useCase := usecase.NewTemperatureUseCase(repo)
	temperatureHandler := handler.NewTemperatureHandler(useCase)

	app.Post("/temperature", temperatureHandler.SaveTemperature)
	app.Get("/temperatures", temperatureHandler.GetAllTemperatures)

	app.Post("/fan", temperatureHandler.ControlFan)
	app.Get("/fan/status", temperatureHandler.GetFanStatus)

	log.Fatal(app.Listen(":3000"))
}
