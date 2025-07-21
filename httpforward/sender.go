package httpforward

import (
	"bytes"
	"log"
	"net/http"
)

var targetAPI = "http://example.com/iot-data"

func ForwardMessage(data []byte) {
	log.Println("Forwarding message to HTTP API...")
	resp, err := http.Post(targetAPI, "application/json", bytes.NewBuffer(data))
	if err != nil {
		log.Printf("HTTP forward error: %v", err)
		return
	}
	defer resp.Body.Close()
	log.Printf("HTTP response status: %s", resp.Status)
}
