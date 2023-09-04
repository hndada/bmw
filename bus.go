package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"time"
)

type BusStop struct {
	ID        int
	Name      string
	BusRoutes []BusRoute
	X, Y      float64
	Address   string
}

type BusRoute struct {
	ID       int
	Name     string
	BusStops []BusStop
}

var BusRoutesMap = make(map[int]BusRoute)
var BusStopsMap = make(map[int]BusStop)

type busDataRecord struct {
	RouteID     int
	RouteName   string
	Order       int
	BusStopID   int
	BusStopName string
	X           float64
	Y           float64
}

func parseBusData() {
	raws := readAllRecords("data/bus.csv")

	// Populate the data structures
	var rs []busDataRecord
	for _, raw := range raws {
		routeID, _ := strconv.Atoi(raw[0])
		routeName := raw[1]
		order, _ := strconv.Atoi(raw[2])
		nodeID, _ := strconv.Atoi(raw[3])
		nodeName := raw[4]
		x, _ := strconv.ParseFloat(raw[5], 64)
		y, _ := strconv.ParseFloat(raw[6], 64)

		rs = append(rs, busDataRecord{
			RouteID:     routeID,
			RouteName:   routeName,
			Order:       order,
			BusStopID:   nodeID,
			BusStopName: nodeName,
			X:           x,
			Y:           y,
		})
	}

	// Sort by route name and order.
	sort.SliceStable(rs, func(i, j int) bool {
		return rs[i].Order < rs[j].Order
	})
	sort.SliceStable(rs, func(i, j int) bool {
		return rs[i].RouteID < rs[j].RouteID
	})
	validateBusData(rs)

	var route BusRoute
	for _, r := range rs {
		if _, ok := BusRoutesMap[r.RouteID]; !ok {
			// Add the previous route to the map if it's not an blank route.
			if route.ID != 0 {
				BusRoutesMap[route.ID] = route
			}
			route = BusRoute{
				ID:   r.RouteID,
				Name: r.RouteName,
			}
		}
		route.BusStops = append(route.BusStops, BusStopsMap[r.BusStopID])
	}

	sort.SliceStable(rs, func(i, j int) bool {
		return rs[i].RouteName < rs[j].RouteName
	})
	sort.SliceStable(rs, func(i, j int) bool {
		return rs[i].BusStopID < rs[j].BusStopID
	})
	var busStop BusStop
	for _, r := range rs {
		if _, ok := BusStopsMap[r.BusStopID]; !ok {
			// Add the previous route to the map if it's not an blank route.
			if busStop.ID != 0 {
				BusStopsMap[r.BusStopID] = busStop
			}
			busStop = BusStop{
				ID:   r.BusStopID,
				Name: r.BusStopName,
				X:    r.X,
				Y:    r.Y,
			}
		}
		busStop.BusRoutes = append(busStop.BusRoutes, BusRoutesMap[r.RouteID])
	}

	fmt.Println(len(BusRoutesMap), BusRoutesMap[100100106])
	fmt.Println(len(BusStopsMap), BusStopsMap[124900026])
}

// Panic when order increases by more than 1.
func validateBusData(rs []busDataRecord) {
	p := func(last, now busDataRecord) {
		fmt.Printf("%+v\n", last)
		fmt.Printf("%+v\n", now)
	}
	//	var exceptionRouteIDs = []int{
	//		100100106, 100100163, 100100373, 100100389,
	//	}

	lastRecord := rs[0]
	for _, r := range rs[1:] {
		if r.RouteID == lastRecord.RouteID {
			if r.RouteName != lastRecord.RouteName {
				p(lastRecord, r)
				panic("name mismatch")
			}
			// if r.Order != lastRecord.Order+1 {
			if r.Order <= lastRecord.Order {
				p(lastRecord, r)
				panic("order increases by more than 1")
			}
		}
		lastRecord = r
	}
}

type Bus struct {
	BusRoute

	current int
	next    int
}

func (b Bus) Direction() string {
	return BusStopsMap[b.next].Name
}

func (b Bus) WaitingTime() time.Duration {
	max := 10 * time.Minute
	wt := rand.Float64() * float64(max)
	return time.Duration(int64(wt))
}
