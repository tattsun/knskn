package main

import (
	"knskn/sensor"
	"knskn/watcher"
	"log"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func IsDebugMode() bool {
	return os.Getenv("DEBUG") != ""
}

func main() {
	// Initialize sensor
	var sen sensor.Sensor
	var err error
	if IsDebugMode() {
		sen = sensor.NewMockSensor()
	} else {
		sen, err = sensor.NewBMESensor()
		if err != nil {
			log.Fatal(err)
		}
	}

	// Start watcher
	w := watcher.NewWatcher(sen)
	go func() {
		w.Start()
	}()

	// Start server
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}
