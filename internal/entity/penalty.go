package entity

import "time"

type Penalty struct {
	Lap                int
	ValueOfMissedShots int
	Time               time.Time
	AverageSpeed       float64
}
