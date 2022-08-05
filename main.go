package main

import (
	"fmt"
	"log"
)

func main() {
	bme, err := NewBME()
	if err != nil {
		log.Fatalf("failed to: %+v", err)
	}
	fmt.Println(bme.GetEnv())
}
