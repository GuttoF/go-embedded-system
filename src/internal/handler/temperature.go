package handler

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"go-embedded-system/src/internal/domain"
	"go-embedded-system/src/internal/usecase"
)

type TemperatureHandler struct {
	useCase *usecase.TemperatureUseCase
}

func NewTemperatureHandler(useCase *usecase.TemperatureUseCase) *TemperatureHandler {
	return &TemperatureHandler{useCase: useCase}
}

func (h *TemperatureHandler) SaveTemperature(c *fiber.Ctx) error {
	var data domain.TemperatureData
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse JSON",
		})
	}

	if err := h.useCase.SaveTemperatureData(context.Background(), &data); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "could not save data",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "temperature data saved successfully",
	})
}

func (h *TemperatureHandler) GetAllTemperatures(c *fiber.Ctx) error {
	data, err := h.useCase.GetAllTemperatureData(context.Background())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "could not retrieve data",
		})
	}

	return c.Status(fiber.StatusOK).JSON(data)
}