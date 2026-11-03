package structs

type Departure struct {
	Dt        string    `json:"dt"`
	Aet       string    `json:"aet"`
	Snam      string    `json:"snam"`
	TripId    int       `json:"trip_id"`
	LineColor string    `json:"lineColor"`
	Dem       bool      `json:"dem"`
	Lcod      string    `json:"lcod"`
	Etn       string    `json:"etn"`
	Dtn       string    `json:"dtn"`
	Et        string    `json:"et"`
	RealTrip  *RealTrip `json:"realTrip,omitempty"`
}

type RealTrip struct {
	Aet        string `json:"aet"`
	Id         string `json:"id"`
	LastCoords struct {
		Lng float64 `json:"lng"`
		Lat float64 `json:"lat"`
	} `json:"lastCoords"`
}
