package main

import (
	"math/rand"
	"time"
)

// 지선버스 3212: 100100212
type Transport struct {
	ID    int
	Type  int
	Color int
	// Direction
	Name     string
	Stations []Station
}

const (
	RedBus   = iota
	BlueBus  // main line bus
	GreenBus // branch bus
	TownBus
	GBus // Gyeonggi-do bus
	Subway
)

func (tp Transport) RandomWaitingTime() time.Duration {
	max := []int{15, 10, 8, 6, 25, 6}[tp.Type]
	minute := time.Duration(float64(max) * rand.Float64())
	return minute * time.Minute
}
