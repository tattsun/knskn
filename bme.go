package main

import (
	"github.com/pkg/errors"
	"github.com/quhar/bme280"
	"golang.org/x/exp/io/i2c"
)

type BME struct {
	bme *bme280.BME280
}

func NewBME() (*BME, error) {
	dev, err := i2c.Open(&i2c.Devfs{Dev: "/dev/i2c-1"}, 0x76)
	if err != nil {
		return nil, errors.Wrap(err, "failed to initialize i2c")
	}

	bme := bme280.New(dev)

	return &BME{bme}, nil
}

func (bme *BME) GetEnv() (temp float64, press float64, hum float64, err error) {
	temp, press, hum, err = bme.bme.EnvData()
	if err != nil {
		err = errors.Wrap(err, "failed to get envs from bme")
	}
	return temp, press, hum, err
}
