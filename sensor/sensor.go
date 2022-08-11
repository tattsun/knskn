package sensor

type Sensor interface {
	Get() (temp float64, press float64, hum float64, err error)
}
