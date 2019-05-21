package metrotransit

import (
	"fmt"

	"strings"

	"github.com/GoPracPro/src/local/platform/config"
	"github.com/GoPracPro/src/local/service/metrotransit/model/response"
	"github.com/GoPracPro/src/local/service/metrotransit/util"
)

func GetRouteDetailsByName(name string) (response.BusRoute, error) {
	routeName := strings.TrimSpace(name)
	if routeName == "" {
		return response.BusRoute{}, fmt.Errorf("route name is mandatory input")
	}
	var routes []response.BusRoute
	err := util.ExecuteGET(config.BusRouteURL, &routes)
	if err != nil || routes == nil || len(routes) < 1 {
		return response.BusRoute{}, fmt.Errorf("GetRouteDetailsByName unable to fetch routes. Reason : %+v", err)
	}
	for _, route := range routes {
		if strings.ToLower(strings.TrimSpace(route.Description)) == strings.ToLower(routeName) {
			return route, nil
		}
	}

	return response.BusRoute{}, fmt.Errorf("GetRouteDetailsByName unable to find matching route for  route name : %s", name)
}

//GetBusRouteDirections : Get the bus routes as per given bus routes
func GetBusRouteDirections(routeID string) ([]response.TextValuePair, error) {
	extURL := fmt.Sprintf(config.BusRouteDirURL, routeID)
	var resModels []response.TextValuePair
	err := util.ExecuteGET(extURL, &resModels)
	return resModels, err
}

//GetBusStops : Get the bus stopes as per given bus route and direction
func GetBusStops(routeID, directionID string) ([]response.TextValuePair, error) {
	extURL := fmt.Sprintf(config.BusStopsURL, routeID, directionID)
	var resModels []response.TextValuePair
	err := util.ExecuteGET(extURL, &resModels)
	return resModels, err
}

//GetBusStopTimeTable : Get the bus stop time tales as per given bus route and direction and stop code
func GetBusStopTimeTable(routeID, directionID, stopCode string) ([]response.BusTimeTable, error) {
	extURL := fmt.Sprintf(config.BusStopTimeTableURL, routeID, directionID, stopCode)
	var resModels []response.BusTimeTable
	err := util.ExecuteGET(extURL, &resModels)
	return resModels, err
}
