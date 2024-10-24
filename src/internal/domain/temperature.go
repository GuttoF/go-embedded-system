package domain

import "time"

type TemperatureData struct {
        ID          string    `json:"id" bson:"_id,omitempty"`
        Temperature float64   `json:"temperature" bson:"temperature"`
        Humidity    float64   `json:"humidity" bson:"humidity"`
        Timestamp   time.Time `json:"timestamp" bson:"timestamp"`
}