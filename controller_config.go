package senter

import (
	"github.com/jinzhu/gorm"
	"time"
)

const controllerConfigTableName string = "sensor_controller_config"

type ControllerConfig struct {
	Id             int64
	ControllerId   int64
	IpAddress      string
	UpdateInterval int64
	NtpIpAddress   string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func NewControllerConfig(controllerId int64, ipAddress string, updateInterval int64, ntpIpAddress string) *ControllerConfig {
	return &ControllerConfig{ControllerId: controllerId, Id: 0, IpAddress: ipAddress, UpdateInterval: updateInterval, NtpIpAddress: ntpIpAddress}
}

func LoadControllerConfigByControllerId(controllerId int64) *ControllerConfig {
	logger.Printf("load controller config by controller id: %d\n", controllerId)
	db := getDb()
	var cs []ControllerConfig
	if err := db.Where("controller_id = ?", controllerId).Find(&cs).Error; err != nil {
		if err != gorm.RecordNotFound {
			logger.Printf("unable to load controller config bycontroller id: %s\n", err)
		} else {
			logger.Printf("no record found: controller id = %d", controllerId)
		}
		return nil
	}
	logger.Printf("cs: %v\n", cs)
	if len(cs) == 0 {
		return nil
	}
	if len(cs) > 1 {
		logger.Println("more than one result by controller id: %d\n", controllerId)
		return nil
	}
	return &(cs[0])
}

func (c ControllerConfig) TableName() string {
	return controllerConfigTableName
}
