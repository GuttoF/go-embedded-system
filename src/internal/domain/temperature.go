package domain

import "time"

type TemperatureData struct {
	ID          string      `gorm:"primaryKey"`
	Temperature float64
	Humidity    float64
	Timestamp   time.Time   `gorm:"autoCreateTime"`
}
