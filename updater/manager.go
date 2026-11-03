package updater

import (
	"net/http"
	"sync"
	"tib/structs"
	"time"

	"github.com/gin-gonic/gin"
)

type TripManager struct {
	mu    sync.RWMutex
	trips map[string]*structs.RealTrip
}

func NewTripManager() *TripManager {
	return &TripManager{
		trips: make(map[string]*structs.RealTrip),
	}
}

func (tm *TripManager) Start() error {
	stops, err := structs.GetStops()
	if err != nil {
		return err
	}

	for _, stop := range stops {
		go tm.createTripTicker(&stop)
	}

	return nil
}

func (tm *TripManager) createTripTicker(stop *structs.Stop) {
	ticker := time.NewTicker(time.Minute)

	for range ticker.C {
		departures, err := stop.GetStopDepartures()
		if err != nil {
			continue
		}

		for _, departure := range departures {
			trip := departure.RealTrip
			if trip != nil {
				tm.SetTrip(trip.Id, trip)
			}
		}
	}
}

func (tm *TripManager) SetTrip(id string, trip *structs.RealTrip) {
	tm.mu.Lock()
	defer tm.mu.Unlock()
	tm.trips[id] = trip
}

func (tm *TripManager) GetTrip(id string) (*structs.RealTrip, bool) {
	tm.mu.RLock()
	defer tm.mu.RUnlock()
	trip, ok := tm.trips[id]
	return trip, ok
}

func (tm *TripManager) GetTripCount() int {
	tm.mu.RLock()
	defer tm.mu.RUnlock()
	return len(tm.trips)
}

func (tm *TripManager) TripCountHandler(c *gin.Context) {
	tripsResponse := structs.TripsResponse{
		Count: tm.GetTripCount(),
	}

	c.JSON(http.StatusOK, tripsResponse)
}
