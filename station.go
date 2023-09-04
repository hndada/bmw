package main

import "time"

type Station struct {
	ID        int
	Type      int
	Name      string
	Direction string
	Address   string
	X, Y      float64 // GPS
}

const (
	BusStation = iota
	TrainStation
)

func (s Station) Duration() time.Time {

}

// func (s Station) Type() int {
// 	if s.ID >= 100_000 {
// 		return BusStop
// 	}
// 	return TrainStation
// }
