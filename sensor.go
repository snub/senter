package senter

import (
	"database/sql"
	"github.com/jinzhu/gorm"
	"time"
)

const sensorTableName string = "sensor"

type Sensor struct {
	Id            int64
	DeviceAddress string
	Name          sql.NullString
	Description   sql.NullString
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func NewSensor(deviceAddress string) *Sensor {
	return &Sensor{Id: 0, DeviceAddress: deviceAddress}
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
	return sensorTableName
}

func (s *Sensor) New() bool {
	return getDb().NewRecord(s)
}

func (s *Sensor) Create() {
	db := getDb()
	if db.NewRecord(s) {
		if err := db.Create(s).Error; err != nil {
			logger.Printf("unable to create sensor: %s\n", err)
		}
	} else {
		logger.Printf("cannot create, sensor already exists with id: %d\n", s.Id)
	}
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
