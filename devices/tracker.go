package devices

import (
	"sync"
	"time"
)

var (
	devices      = make(map[string]time.Time)
	devicesMutex = sync.Mutex{}
)

func UpdateDevice(deviceID string) {
	devicesMutex.Lock()
	defer devicesMutex.Unlock()
	devices[deviceID] = time.Now()
}

func GetActiveDevices() map[string]time.Time {
	devicesMutex.Lock()
	defer devicesMutex.Unlock()
	copy := make(map[string]time.Time)
	for k, v := range devices {
		copy[k] = v
	}
	return copy
}
