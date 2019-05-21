package response

type BusTimeTable struct {
	Actual          bool    `json:"Actual"`
	BlockNumber     int     `json:"BlockNumber"`
	DepartureText   string  `json:"DepartureText"`
	DepartureTime   string  `json:"DepartureTime"`
	Description     string  `json:"Description"`
	Gate            string  `json:"Gate"`
	Route           string  `json:"Route"`
	RouteDirection  string  `json:"RouteDirection"`
	Terminal        string  `json:"Terminal"`
	VehicleLatitude float64 `json:"VehicleLatitude"`
}
