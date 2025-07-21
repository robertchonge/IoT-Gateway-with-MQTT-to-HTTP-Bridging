package mqtt

import (
	"fmt"
	"iot-gateway/httpforward"
	"iot-gateway/devices"
	"github.com/eclipse/paho.mqtt.golang"
	"log"
	"time"
)

var broker = "tcp://localhost:1883"
var topic = "iot/data"

func messageHandler(client mqtt.Client, msg mqtt.Message) {
	log.Printf("Received message on topic %s: %s", msg.Topic(), msg.Payload())
	devices.UpdateDevice(msg.Topic())
	httpforward.ForwardMessage(msg.Payload())
}

func StartMQTTClient() {
	opts := mqtt.NewClientOptions().AddBroker(broker)
	opts.SetClientID("iot-gateway-client")
	opts.SetDefaultPublishHandler(messageHandler)

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	if token := client.Subscribe(topic, 0, nil); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	log.Println("MQTT client connected and subscribed to topic.")

	for {
		time.Sleep(1 * time.Second)
	}
}
