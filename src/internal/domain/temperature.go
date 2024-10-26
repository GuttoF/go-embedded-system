package domain

import "time"

type TemperatureData struct {
	ID          uint `gorm:"primaryKey"`
	Temperature float64
	Humidity    float64
	Timestamp   time.Time `gorm:"autoCreateTime"`
}

type FanControl struct {
	State string `json:"state"` // on or off
}
