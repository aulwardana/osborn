package sensor

import (
	"errors"
)

var (
	FailGet         = errors.New("sensor : Data Not Found")
	FailInsert      = errors.New("sensor : Fail Input Data")
	FailInsertExist = errors.New("sensor : Data Already Exist")
	FailUpdate      = errors.New("sensor : Fail Update Data")
	FailDelete      = errors.New("sensor : Fail Delete Data")
)

type Sensor struct {
	Id          interface{} `json:"_id" bson:"_id"`
	Code        string      `json:"code" bson:"code"`
	Temperature int         `json:"temperature" bson:"temperature"`
	Humidity    int         `json:"humidity" bson:"humidity"`
}

type SensorInsert struct {
	Code        string `json:"code" bson:"code"`
	Temperature int    `json:"temperature" bson:"temperature"`
	Humidity    int    `json:"humidity" bson:"humidity"`
}

type SensorService interface {
	//GetSensorData(code string) (*Sensor, error)
	InsertSensorData(r *SensorInsert) error
	//UpdateSensorData(code string, s *Sensor) error
	//DeleteSensorData(code string) error
}
