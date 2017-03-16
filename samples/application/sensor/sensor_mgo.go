package sensor

import (
	"gopkg.in/mgo.v2"
)

const sensorCollection string = "coba"

type SensorServiceMongo struct {
	db     *mgo.Session
	dbName string
}

func (m *SensorServiceMongo) InsertSensorData(r *SensorInsert) error {
	c := m.db.DB(m.dbName).C(sensorCollection)
	if err := c.Insert(r); err != nil {
		if mgo.IsDup(err) {
			return FailInsertExist
		}

		return FailInsert
	}

	return nil
}

func New(db *mgo.Session, dbName string) SensorService {
	svc := &SensorServiceMongo{
		db:     db,
		dbName: dbName,
	}

	return svc
}
