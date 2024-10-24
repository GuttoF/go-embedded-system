package repository

import (
        "context"
        "log"
        "go-embedded-system/src/internal/db"
        "go-embedded-system/src/internal/domain"
        "time"
)

type TemperatureRepository struct{}

func NewTemperatureRepository() *TemperatureRepository {
        return &TemperatureRepository{}
}

func (r *TemperatureRepository) Save(ctx context.Context, data *domain.TemperatureData) error {
        // add timestamp before saving
        data.Timestamp=time.Now()

        _, err := db.FirebaseClient.NewRef("temperature_readings").Push(ctx, data)
        if err != nil {
                log.Printf("error saving temperature data: %v", err)
                return err
        }
        return nil
}

func (r *TemperatureRepository) GetAll(ctx context.Context) ([]*domain.TemperatureData, error) {
        var results []*domain.TemperatureData
        err := db.FirebaseClient.NewRef("temperature_readings").Get(ctx, &results)
        if err != nil {
                log.Printf("error retrieving temperature data: %v", err)
        return nil, err
        }
        return results, nil
}