package state_machine

import "github.com/MartinezGitHub/system-prototype-for-biathlon-competitions/internal/entity"

type StateMachine interface {
	CanTransition(currentStatus entity.Status, event entity.EventType) bool
	ApplyTransition(currentStatus entity.Status, event entity.EventType) (entity.Status, error)
}
