package config

//TODO: move to a config file
const (
	BusRouteURL         string = "http://svc.metrotransit.org/NexTrip/Routes"
	BusRouteDirURL      string = "http://svc.metrotransit.org/NexTrip/Directions/%s"
	BusStopsURL         string = "http://svc.metrotransit.org/NexTrip/Stops/%s/%s"
	BusStopTimeTableURL string = "http://svc.metrotransit.org/NexTrip/%s/%s/%s"
)
