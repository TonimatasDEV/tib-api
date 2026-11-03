package main

import (
	"log"
	"net/http"
	"tib/structs"
	"tib/updater"

	"github.com/gin-gonic/gin"
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

	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())

	router.GET("/", Handler)
	router.GET("/trips", manager.TripCountHandler)

	address := "0.0.0.0:8080"
	err = router.Run(address)
	if err != nil {
		log.Fatal(err)
	}
}

func Handler(c *gin.Context) {
	response := structs.MainResponse{
		Msg:          "It works!",
		Repository:   "https://github.com/TonimatasDEV/tib-api",
		IssueTracker: "https://github.com/TonimatasDEV/tib-api/issues",
	}

	c.JSON(http.StatusOK, response)
}
