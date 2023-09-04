package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"time"
)

type TrainLine struct {
	Name     string // There is no ID for train line.
	Stations []TrainStation
}

// Station doesn't have Direction, Transport does.
type TrainStation struct {
	ID      int
	Name    string // Outer code is attached to the name.
	X, Y    float64
	Address string
}

var TrainLines = make(map[string]TrainLine)
var TrainStationsMap = make(map[int]TrainStation)

type TrainDataRecord struct {
	StationID   int // I guess some of them in train_code are not numbers.
	StationName string
	StationCode string // Outer code
	LineName    string
	X           float64
	Y           float64
}

func parseTrainData() {
	var trainStationCodes = make(map[int]string)
	codeRaws := readAllRecords("data/train_code.csv")
	for _, raw := range codeRaws {
		id, _ := strconv.Atoi(raw[0])
		trainStationCodes[id] = raw[4]
	}

	raws := readAllRecords("data/train.csv")

	// Populate the data structures
	var rs []TrainDataRecord
	for _, raw := range raws {
		nodeID, _ := strconv.Atoi(raw[0])
		nodeName := raw[1]
		// StationCode is derived from trainStationCode.
		lineName := raw[2]
		x, _ := strconv.ParseFloat(raw[3], 64)
		y, _ := strconv.ParseFloat(raw[4], 64)

		rs = append(rs, TrainDataRecord{
			StationID:   nodeID,
			StationName: nodeName,
			StationCode: trainStationCodes[nodeID],
			LineName:    lineName,
			X:           x,
			Y:           y,
		})
	}

	sort.SliceStable(rs, func(i, j int) bool {
		return rs[i].StationCode < rs[j].StationCode
	})
	sort.SliceStable(rs, func(i, j int) bool {
		return rs[i].LineName < rs[j].LineName
	})

	var line TrainLine
	for _, r := range rs {
		if _, ok := TrainLines[r.LineName]; !ok {
			if line.Name != "" {
				TrainLines[line.Name] = line
			}
			line = TrainLine{
				Name: r.LineName,
			}
		}
		if _, ok := TrainStationsMap[r.StationID]; !ok {
			TrainStationsMap[r.StationID] = TrainStation{
				ID:   r.StationID,
				Name: r.StationName,
				X:    r.X,
				Y:    r.Y,
			}
		}
		line.Stations = append(line.Stations, TrainStationsMap[r.StationID])
	}

	fmt.Println(len(TrainLines), TrainLines["김포골드라인"])
	fmt.Println(len(TrainStationsMap), TrainStationsMap[9996])
}

type Train struct {
	TrainLine

	current int
	next    int
}

func (t Train) Direction() string {
	return TrainStationsMap[t.next].Name
}

func (t Train) WaitingTime() time.Duration {
	max := 8 * time.Minute
	wt := rand.Float64() * float64(max)
	return time.Duration(int64(wt))
}
