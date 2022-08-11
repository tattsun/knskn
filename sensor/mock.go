package sensor

import (
	"math/rand"
	"sync"
)

type MockSensor struct {
	DeltaTemp  float64
	DeltaPress float64
	DeltaHum   float64
	Temp       float64
	Press      float64
	Hum        float64
	mtx        *sync.Mutex
}

func NewMockSensor() *MockSensor {
	return &MockSensor{
		Temp:  25,
		Press: 1002,
		Hum:   60,
		mtx:   new(sync.Mutex),
	}
}

func (m *MockSensor) Get() (temp float64, press float64, hum float64, err error) {
	m.mtx.Lock()
	defer m.mtx.Unlock()

	m.Temp += rand.Float64() - 0.5
	m.Press += rand.Float64() - 0.5
	m.Hum += rand.Float64() - 0.5

	return m.Temp, m.Press, m.Hum, nil
}
