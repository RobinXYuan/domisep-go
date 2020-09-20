package sensors

import (
	"github.com/jinzhu/gorm"
)

type SensorModel struct {
	gorm.Model
	Name   string `json:"name"`
	Type   string `json:"type"`
	Module string `json:"module`
}

type SensorDataModel struct {
	gorm.Model
	SensorID   uint16  `json:"sensorId"`
	CreateTime string  `json:"createTime"`
	Value      float32 `json:"value"`
}
