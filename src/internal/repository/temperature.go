package repository

import (
        "context"
        "go.mongodb.org/mongo-driver/mongo"
        "smart_cooling_system/internal/domain"
        "time"
)

type TemperatureRepository struct {
        collection *mongo.Collection
}

func NewTemperatureRepository(db *mongo.Database) *TemperatureRepository {
        return &TemperatureRepository{
                collection: db.Collection("temperature_readings"),
        }
}

func (r *TemperatureRepository) Save(ctx context.Context, data *domain.TemperatureData) error {
        data.Timestamp = time.Now()
        _, err := r.collection.InsertOne(ctx, data)
        return err
}