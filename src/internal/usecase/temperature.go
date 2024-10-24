package usecase

import (
	"context"
	"go-embedded-system/src/internal/domain"
	"go-embedded-system/src/internal/repository"
)

type TemperatureUseCase struct {
	repo *repository.TemperatureRepository
}

func NewTemperatureUseCase(repo *repository.TemperatureRepository) *TemperatureUseCase {
	return &TemperatureUseCase{repo: repo}
}

func (uc *TemperatureUseCase) SaveTemperatureData(ctx context.Context, data *domain.TemperatureData) error {
	return uc.repo.Save(ctx, data)
}

func (uc *TemperatureUseCase) GetAllTemperatureData(ctx context.Context) ([]*domain.TemperatureData, error) {
	return uc.repo.GetAll(ctx)
}