package main

import (
	"go-embedded-system/src/internal/db"
	"go-embedded-system/src/internal/handler"
	"go-embedded-system/src/internal/repository"
	"go-embedded-system/src/internal/usecase"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

        db.InitFirebase()

	repo := repository.NewTemperatureRepository()
	useCase := usecase.NewTemperatureUseCase(repo)
	handler := handler.NewTemperatureHandler(useCase)

	app.Post("/temperature", handler.SaveTemperature)

	log.Fatal(app.Listen(":3000"))
}
