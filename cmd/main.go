package main

import (
	"iot-gateway/mqtt"
	"iot-gateway/api"
	"iot-gateway/monitoring"
	"log"
)

func main() {
	log.Println("Starting IoT Gateway Service...")
	go monitoring.StartPrometheusServer() // Expose /metrics
	go api.StartHTTPServer()              // Expose API endpoints
	mqtt.StartMQTTClient()                // Block: subscribe & process
}

