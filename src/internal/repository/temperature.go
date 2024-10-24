package repository

import (
	"context"
	"go-embedded-system/src/internal/db"
	"go-embedded-system/src/internal/domain"
	"log"
	"time"
)

type TemperatureRepository struct{}

func NewTemperatureRepository() *TemperatureRepository {
	return &TemperatureRepository{}
}

func (r *TemperatureRepository) Save(ctx context.Context, data *domain.TemperatureData) error {
	// add timestamp before saving
	data.Timestamp = time.Now()

	ref, err := db.FirebaseClient.NewRef("temperature_readings").Push(ctx, data)
	if err != nil {
		log.Printf("error saving temperature data: %v", err)
		return err
	}

	data.ID = ref.Key
	if err := ref.Set(ctx, data); err != nil {
		log.Printf("error updating temperature data with ID: %v", err)
		return err
	}

	return nil
}

func (r *TemperatureRepository) GetAll(ctx context.Context) ([]*domain.TemperatureData, error) {
	var results map[string]*domain.TemperatureData

	err := db.FirebaseClient.NewRef("temperature_readings").Get(ctx, &results)
	if err != nil {
		log.Printf("error retrieving temperature data: %v", err)
		return nil, err
	}

	if results == nil {
		log.Println("no data found")
		return []*domain.TemperatureData{}, nil
	}

	var dataList []*domain.TemperatureData
	for _, data := range results {
		dataList = append(dataList, data)
	}

	return dataList, nil
}
