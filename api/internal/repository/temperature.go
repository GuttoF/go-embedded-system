package repository

import (
	"context"
	"go-embedded-system/internal/db"
	"go-embedded-system/internal/domain"
	"log"
)

type TemperatureRepository struct{}

func NewTemperatureRepository() *TemperatureRepository {
	return &TemperatureRepository{}
}

func (r *TemperatureRepository) Save(ctx context.Context, data *domain.TemperatureData) error {
	if err := db.DB.WithContext(ctx).Create(data).Error; err != nil {
			log.Printf("error saving temperature data: %v", err)
			return err
	}
	return nil
}

func (r *TemperatureRepository) GetAll(ctx context.Context) ([]*domain.TemperatureData, error) {
	var dataList []*domain.TemperatureData
	if err := db.DB.WithContext(ctx).Find(&dataList).Error; err != nil {
			log.Printf("error retrieving temperature data: %v", err)
			return nil, err
	}
	return dataList, nil
}