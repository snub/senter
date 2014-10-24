package senter

import "github.com/jinzhu/gorm"

type Sensor struct {
	Id            int64
	DeviceAddress string
}

func NewSensor(deviceAddress string) *Sensor {
	return &Sensor{0, deviceAddress}
}

// TODO better error handling
func LoadSensorByDeviceAddress(deviceAddress string) *Sensor {
	logger.Printf("load sensor by device address: %s\n", deviceAddress)
	db := getDb()
	var ss []Sensor
	if err := db.Where("device_address = ?", deviceAddress).Find(&ss).Error; err != nil {
		if err != gorm.RecordNotFound {
			logger.Printf("unable to load sensor by device address: %s\n", err)
		} else {
			logger.Printf("no record found: device address = %s", deviceAddress)
		}
		return NewSensor(deviceAddress)
	}
	logger.Printf("ss: %v\n", ss)
	if len(ss) == 0 {
		return NewSensor(deviceAddress)
	}
	if len(ss) > 1 {
		logger.Printf("more than one result by device address: %s\n", deviceAddress)
		return nil

	}
	return &(ss[0])
}

func (s Sensor) TableName() string {
	return "sensor"
}

// TODO on update check rows affected
func (s *Sensor) Save() {
	db := getDb()
	if db.NewRecord(s) {
		if err := db.Create(s).Error; err != nil {
			logger.Printf("unable to create sensor: %s\n", err)
		}
	} else {
		if err := db.Save(s).Error; err != nil {
			logger.Printf("unable to save sensor: %s\n", err)
		}
	}
}
