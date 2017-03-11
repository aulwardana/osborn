package mqttlib

import (
	"fmt"
	"log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var f mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("TOPIC: %s\n", msg.Topic())
	fmt.Printf("MSG: %s\n", msg.Payload())
}

func Connect(protocol string, url string, port int) mqtt.Client {
	addr := fmt.Sprintf("%s://%s:%d", protocol, url, port)
	opts := mqtt.NewClientOptions().AddBroker(addr)
	opts.SetClientID("osborn")
	opts.SetDefaultPublishHandler(f)

	c := mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	log.Println("MQTT connected:" + addr)

	return c
}
