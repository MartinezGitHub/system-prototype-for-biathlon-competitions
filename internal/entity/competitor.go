package entity

import (
	"github.com/pkg/errors"
	"time"
)

type Competitor struct {
	ID                       int
	StartTime                time.Time
	FinishTime               time.Time
	Status                   Status
	LapTimes                 []*LapTime
	Penalties                []*Penalty
	Shots                    []*Shot
	Disqualified             bool
	Comment                  string
	CurrentLap               int
	SucceedShotsOnCurrentLap int
}

type LapTime struct {
	Number       int
	Time         time.Time
	AverageSpeed float64
	Active       bool
}

type Shot struct {
	Lap    int
	Range  int
	Target int
	Time   time.Time
}

type Status string

const (
	Registered  Status = "Registered"
	OnStartLine Status = "OnStartLine"
	Racing      Status = "Racing"
	OnShooting  Status = "OnShooting"
	OnPenalty   Status = "OnPenalty"
	NotStarted  Status = "NotStarted"
	NotFinished Status = "NotFinished"
)

var (
	ErrCompetitorExists   = errors.New("Competitor already exists")
	ErrCompetitorNotFound = errors.New("Competitor not found")
)
