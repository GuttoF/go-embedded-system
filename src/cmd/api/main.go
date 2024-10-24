package main

import (
	"go-embedded-system/src/internal/db"
	"go-embedded-system/src/internal/handler"
	"go-embedded-system/src/internal/repository"
	"go-embedded-system/src/internal/usecase"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	mongoURI := os.Getenv("MONGO_URI")
	database, err := db.ConnectMongo(mongoURI)
	if err != nil {
		log.Fatal("Failed to connect to MongoDB: ", err)
	}

	repo := repository.NewTemperatureRepository(database)
	useCase := usecase.NewTemperatureUseCase(repo)
	handler := handler.NewTemperatureHandler(useCase)

	app.Post("/temperature", handler.SaveTemperature)

	log.Fatal(app.Listen(":3000"))
}
