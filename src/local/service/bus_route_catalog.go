package service

import (
	"fmt"
	"strings"

	"github.com/GoPracPro/src/local/service/metrotransit"
	resModel "github.com/GoPracPro/src/local/service/metrotransit/model/response"
)

const (
	emptyString = ""
)

func GetNextBusTimeToGo(busRouteName, busStopName, busDirection string) (string, error) {
	busRoute, err := metrotransit.GetRouteDetailsByName(busRouteName)
	if err != nil {
		return emptyString, fmt.Errorf("Unable to find route details for input route : %s. Reason : %e", busRouteName, err)
	}
	busDirectionDet, err := getDirectionForRoute(busDirection, busRoute)
	if err != nil {
		return emptyString, err
	}
	stopDetail, err := getStopDetailsForRoute(busStopName, busRoute, busDirectionDet)
	if err != nil {
		return emptyString, err
	}
	return getTimeToGO(busRoute, busDirectionDet, stopDetail)
}

func getDirectionForRoute(inpBusDirection string, route resModel.BusRoute) (resModel.TextValuePair, error) {
	var directionDet resModel.TextValuePair
	busDirection := strings.ToUpper(strings.TrimSpace(inpBusDirection))
	directions, err := metrotransit.GetBusRouteDirections(route.Route)
	if err != nil {
		return directionDet, fmt.Errorf("Unable to find route directions for input route : %s. Reason : %+v", route.Description, err)
	}
	for _, namVal := range directions {
		direction := strings.ToUpper(namVal.Text)
		if strings.Contains(direction, busDirection) {
			directionDet = namVal
			break
		}
	}
	if directionDet.Text == emptyString {
		return directionDet, fmt.Errorf("Unable to find route directions for input route : %s. Direction : %s", route.Description, inpBusDirection)
	}
	return directionDet, nil
}

func getStopDetailsForRoute(inpBusStopName string, route resModel.BusRoute, dir resModel.TextValuePair) (resModel.TextValuePair, error) {
	busStopName := strings.ToUpper(strings.TrimSpace(inpBusStopName))
	var stopDetail resModel.TextValuePair
	busStops, err := metrotransit.GetBusStops(route.Route, dir.Value)
	if err != nil {
		return stopDetail, fmt.Errorf("Unable to find stop details  for input route : %s. Direction : %s", route.Description, dir.Text)
	}
	for _, namVal := range busStops {
		nameValText := strings.ToUpper(strings.TrimSpace(namVal.Text))
		if nameValText == busStopName {
			stopDetail = namVal
			break
		}
	}
	if stopDetail.Text == emptyString {
		return stopDetail, fmt.Errorf("Unable to find stop details for input route : %s. Direction : %s and matching input stop name : %s", route.Description, dir.Text, inpBusStopName)
	}
	return stopDetail, nil
}

func getTimeToGO(route resModel.BusRoute, dir resModel.TextValuePair, stopDetail resModel.TextValuePair) (string, error) {
	timeTableDetails, err := metrotransit.GetBusStopTimeTable(route.Route, dir.Value, stopDetail.Value)
	if err != nil {
		return "", fmt.Errorf("Unable to find stop time details  for input route : %s. Direction : %s, stop details : %s. Reason : %+v", route.Description, dir.Text, stopDetail.Text, err)
	}
	for _, timeTabDet := range timeTableDetails {
		if timeTabDet.Actual {
			return timeTabDet.DepartureText, nil
		}

	}

	return "", nil
}
