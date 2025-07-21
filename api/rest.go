package api

import (
	"encoding/json"
	"iot-gateway/auth"
	"iot-gateway/devices"
	"log"
	"net/http"
)

func StartHTTPServer() {
	http.Handle("/devices", auth.AuthMiddleware(http.HandlerFunc(devicesHandler)))
	log.Println("API server running on :8080")
	http.ListenAndServe(":8080", nil)
}

func devicesHandler(w http.ResponseWriter, r *http.Request) {
	devices := devices.GetActiveDevices()
	json.NewEncoder(w).Encode(devices)
}

