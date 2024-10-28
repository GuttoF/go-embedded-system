package domain

import "time"

type TemperatureData struct {
	ID          uint `gorm:"primaryKey"`
	Temperature float64
	Humidity    float64
	Timestamp   time.Time `gorm:"autoCreateTime"`
}

func (t *TemperatureData) AdjustTime() {
	location, _ := time.LoadLocation("America/Sao_Paulo")
	t.Timestamp = t.Timestamp.In(location)
}

type FanControl struct {
	State string `json:"state"` // on or off
}
