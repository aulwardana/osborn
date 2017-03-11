package main

import (
	"osborn/core/libs/mqttlib"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type mqttSession struct {
	c mqtt.Client
}

func sessionMqtt(c mqtt.Client) *mqttSession {
	return &mqttSession{
		c: c,
	}
}

func initMqtt() {
	sessionMqtt(mqttlib.Connect(cnf.Mqtt().Protocol, cnf.Mqtt().Url, cnf.Mqtt().Port))
}
