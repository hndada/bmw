package main

import "math"

type XY struct {
	X float64 // Longitude
	Y float64 // Latitude
}

// BusStop, TrainStation, Landmark implement Place interface.
type Place interface {
	Type() int
	Name() string
	Address() string
	Coordinates() (float64, float64)
}

// km/h
const (
	WalkSpeed  = 5.0
	BusSpeed   = 30.0
	TrainSpeed = 60.0
)

func TravelTime(p, q XY, speed float64) float64 {
	distance := HaversineDistance(p, q)
	return distance / speed
}

// HaversineDistance calculates the distance between two points on the Earth
// given their latitude and longitude using the Haversine formula.
func HaversineDistance(p XY, q XY) float64 {
	const R = 6371.0 // Earth's radius in kilometers

	// Convert degrees to radians
	lat1 := p.X * math.Pi / 180.0
	lon1 := p.Y * math.Pi / 180.0
	lat2 := q.X * math.Pi / 180.0
	lon2 := q.Y * math.Pi / 180.0

	// Haversine formula
	dlat := lat2 - lat1
	dlon := lon2 - lon1
	a := math.Sin(dlat/2)*math.Sin(dlat/2) + math.Cos(lat1)*math.Cos(lat2)*math.Sin(dlon/2)*math.Sin(dlon/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	// Distance in kilometers
	return R * c
}
