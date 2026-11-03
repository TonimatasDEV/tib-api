package structs

type TripsResponse struct {
	Count int `json:"tripCount"`
}

type MainResponse struct {
	Msg          string `json:"msg"`
	Repository   string `json:"repository"`
	IssueTracker string `json:"issueTracker"`
}
