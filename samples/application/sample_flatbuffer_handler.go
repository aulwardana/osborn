package main

import (
	"fmt"

	"osborn/samples/application/sensors"

	flatbuffers "github.com/google/flatbuffers/go"
)

func MakeSensor(b *flatbuffers.Builder, id int32, temperature float32, humidity float32, topic []byte) []byte {
	b.Reset()

	topic_position := b.CreateByteString(topic)

	sensors.SensorStart(b)
	sensors.SensorAddTopic(b, topic_position)
	sensors.SensorAddId(b, id)
	sensors.SensorAddTemperature(b, temperature)
	sensors.SensorAddHumidity(b, humidity)

	sensor_position := sensors.SensorEnd(b)

	b.Finish(sensor_position)

	return b.Bytes[b.Head():]
}

func ReadSensor(buf []byte) (id int32, temperature float32, humidity float32, topic []byte) {
	sensor := sensors.GetRootAsSensor(buf, 0)

	topic = sensor.Topic()
	id = sensor.Id()
	humidity = sensor.Humidity()
	temperature = sensor.Temperature()

	return
}

func FlatProc() {
	b := flatbuffers.NewBuilder(0)
	buf := MakeSensor(b, 7, 20.5, 30.5, []byte("sensing"))
	id, temperature, humidity, topic := ReadSensor(buf)
	fmt.Printf("Topic %s has id %d. The temperature %6.2f and humudity %6.2f. Payload %d bytes long.\n", topic, id, temperature, humidity, len(buf))
}
