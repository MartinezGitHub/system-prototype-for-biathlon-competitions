package entity

import "time"

type Event struct {
	Time         time.Time
	Type         EventType
	CompetitorID int
	Payload      any
}

type EventType int

const (
	EventRegister EventType = iota + 1
	EventSetStartTime
	EventOnStartLine
	EventStart
	EventOnShootingRange
	EventTargetHit
	EventLeaveShootingRange
	EventEnterPenalty
	EventLeavePenalty
	EventFinishLap
	EventCannotContinue

	EventDisqualified = 32
	EventFinished     = 33
)

type RangePayload struct {
	Range int
}

type TargetHitPayload struct {
	Target int
}

type CommentPayload struct {
	Text string
}

type StartTimePayload struct {
	Time time.Time
}
