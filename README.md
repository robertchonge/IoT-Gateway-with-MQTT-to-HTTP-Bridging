# IoT-Gateway-with-MQTT-to-HTTP-Bridging

#Overview

Designed a gateway service that:

-Receives telemetry data from IoT devices via MQTT.

-Converts and forwards this data to external REST APIs.

-Supports real-time monitoring of connected devices.

-Secures communication via lightweight authentication.

## Features
- MQTT subscription & message processing
- HTTP forwarding of IoT telemetry
- Device tracking & management
- Prometheus metrics exporter
- REST API with basic token authentication

## Build & Run
```bash
docker build -t iot-gateway .
docker run -p 8080:8080 -p 9100:9100 iot-gateway
```

## Endpoints
- `/devices` – List active devices (requires Bearer token)
- `/metrics` – Prometheus metrics endpoint

## Configuration
- MQTT broker at `tcp://localhost:1883`
- HTTP forwarding to `http://example.com/iot-data`
- API token set in `auth/middleware.go`
