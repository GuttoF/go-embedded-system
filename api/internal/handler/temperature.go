package handler

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"go-embedded-system/internal/domain"
	"go-embedded-system/internal/usecase"
)

type TemperatureHandler struct {
	useCase *usecase.TemperatureUseCase
	fanStatus string
}

func NewTemperatureHandler(useCase *usecase.TemperatureUseCase) *TemperatureHandler {
	return &TemperatureHandler{
		useCase: useCase,
		fanStatus: "off",
	}
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

func (h *TemperatureHandler) ControlFan(c *fiber.Ctx) error {
	var fanControl domain.FanControl
	if err := c.BodyParser(&fanControl); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse JSON",
		})
	}

	if fanControl.State == "on" || fanControl.State == "off" {
		h.fanStatus = fanControl.State
		return c.JSON(fiber.Map{"message": "Fan status updated", "state": h.fanStatus})
	} else {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid state. Use 'on' or 'off'.",
		})
	}
}

func (h *TemperatureHandler) GetFanStatus(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"state": h.fanStatus})
}