package main

import (
	"fmt"
	"log"
)

func main() {
	bme, err := NewBME()
	if err != nil {
		log.Fatalf("failed to initialize BME: %+v", err)
	}

	temp, press, hum, err := bme.GetEnv()
	if err != nil {
		log.Fatalf("failed to get environments: %+v", err)
	}

	fmt.Printf("Temp: %f\n", temp)
	fmt.Printf("Pressure: %f\n", press)
	fmt.Printf("Humidity: %f\n", hum)
}
