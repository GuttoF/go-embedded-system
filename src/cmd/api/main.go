package main

import (
	"go-embedded-system/src/internal/db"
	"go-embedded-system/src/internal/handler"
	"go-embedded-system/src/internal/domain"
	"go-embedded-system/src/internal/repository"
	"go-embedded-system/src/internal/usecase"
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

	log.Fatal(app.Listen(":3000"))
}
