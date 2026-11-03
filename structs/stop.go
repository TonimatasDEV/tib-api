package structs

import (
	"encoding/json"
	"strings"
	"tib/utils"
)

var stopsCache []Stop

type Stop struct {
	StopCode string `json:"stopCode"`
	StopId   int    `json:"stopId"`
	StopName string `json:"stopName"`
}

const (
	StopsURL      = "https://www.tib.org/o/manager/stops/ctmr4"
	DeparturesURL = "https://www.tib.org/o/manager/stop-code/{stopCode}/departures/ctmr4"
)

func GetStops() ([]Stop, error) {
	if stopsCache != nil {
		return stopsCache, nil
	}

	body, err := utils.Request(StopsURL)
	if err != nil {
		return nil, err
	}

	var stops []Stop
	if err := json.Unmarshal(body, &stops); err != nil {
		return nil, err
	}

	stopsCache = stops

	return stops, nil
}

func (s *Stop) GetStopDepartures() ([]Departure, error) {
	body, _ := utils.Request(strings.ReplaceAll(DeparturesURL, "{stopCode}", s.StopCode))

	var departuresStop []Departure
	if err := json.Unmarshal(body, &departuresStop); err != nil {
		return nil, err
	}

	return departuresStop, nil
}
