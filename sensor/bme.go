package sensor

import (
	"github.com/pkg/errors"
	"github.com/quhar/bme280"
	"golang.org/x/exp/io/i2c"
)

type BMESensor struct {
	bme *bme280.BME280
}

func NewBMESensor() (*BMESensor, error) {
	dev, err := i2c.Open(&i2c.Devfs{Dev: "/dev/i2c-1"}, 0x76)
	if err != nil {
		return nil, errors.Wrap(err, "failed to initialize i2c")
	}

	bme := bme280.New(dev,
		bme280.Standby(bme280.Stbdby1000),
		bme280.Mode(bme280.NormalMode),
		bme280.TempOverSmpl(bme280.OverSmpl1<<5),
		bme280.PressOverSmpl(bme280.OverSmpl1<<3),
		bme280.HumOverSmpl(bme280.OverSmpl1),
	)

	err = bme.Init()
	if err != nil {
		return nil, errors.Wrap(err, "failed to initialize bme")
	}

	return &BMESensor{bme}, nil
}

func (bme *BMESensor) Get() (temp float64, press float64, hum float64, err error) {
	temp, press, hum, err = bme.bme.EnvData()
	if err != nil {
		err = errors.Wrap(err, "failed to get envs from bme")
	}
	return temp, press, hum, err
}
