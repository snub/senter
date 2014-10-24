package senter

import "time"

type Temperature struct {
	Id        int64
	SensorId  int64
	Timestamp time.Time
	Value     float32
}

func NewTemperature(sensor *Sensor, timestamp int64, value float32) *Temperature {
	return &Temperature{0, sensor.Id, time.Unix(timestamp, 0).UTC(), value}
}

func (t Temperature) TableName() string {
	return "sensor_temperature"
}

func (t *Temperature) Save() {
	db := getDb()
	if db.NewRecord(t) {
		if err := db.Create(t).Error; err != nil {
			logger.Printf("unable to create temperature: %s\n", err)
		}
	} else {
		logger.Println("temperature does not suppor updating")
	}
}
