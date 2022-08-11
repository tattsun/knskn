package watcher

import (
	"knskn/sensor"
	"log"
	"time"

	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	Temparture = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "temparture",
		Help: "temparture of the room",
	})
	Humidity = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "humidity",
		Help: "humidity of the room",
	})
	Pressure = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "pressure",
		Help: "air pressure of the room",
	})
	DiscomfortIndex = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "discomfort_index",
		Help: "discomfort index of the room",
	})
)

type Watcher struct {
	Sensor sensor.Sensor
}

func NewWatcher(sensor sensor.Sensor) *Watcher {
	return &Watcher{
		Sensor: sensor,
	}
}

func (w *Watcher) Start() {
	for {
		if err := w.RunTick(); err != nil {
			log.Printf("failed to get values from sensor: %+v", err)
		}
		time.Sleep(1 * time.Second)
	}
}

func (w *Watcher) RunTick() error {
	temp, press, hum, err := w.Sensor.Get()
	if err != nil {
		return errors.Wrap(err, "")
	}

	Temparture.Set(temp)
	Humidity.Set(hum)
	Pressure.Set(press)
	DiscomfortIndex.Set(CalcDiscomfortIndex(temp, hum))

	return nil
}
