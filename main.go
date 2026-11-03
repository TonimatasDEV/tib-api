package main

import (
	"estamos-zumbados/updater"
	"fmt"
	"log"
	"time"
)

const (
	WebsocketURL = "wss://sae.tib.org/saews/public-events/{tripId}"
)

func main() {
	manager := updater.NewTripManager()

	err := manager.Start()
	if err != nil {
		log.Fatal(err)
	}

	for {
		time.Sleep(30 * time.Second)
		fmt.Printf("Trip count: %d\n", manager.GetTripCount())
	}
}
