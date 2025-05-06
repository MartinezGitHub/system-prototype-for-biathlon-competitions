package biathlon

import (
	"fmt"
	"github.com/MartinezGitHub/system-prototype-for-biathlon-competitions/internal/entity"
	"time"
)

func (b *biathlonImpl) ProcessOnShootingRange(event entity.Event) error {

	err := b.competitorRepository.MarkOnShootingRange(event.CompetitorID)
	if err != nil {
		return err
	}
	competitor, err := b.competitorRepository.GetCompetitor(event.CompetitorID)
	if err != nil {
		return err
	}
	competitor.SucceedShotsOnCurrentLap = 0
	competitor.LapTimes[len(competitor.LapTimes)-1].Active = true

	return nil
}

func (b *biathlonImpl) ProcessTargetHit(event entity.Event) error {
	shotPayload, ok := event.Payload.(entity.TargetHitPayload)
	if !ok {
		return fmt.Errorf("invalid payload type for ProcessTargetHit event")
	}

	competitor, err := b.competitorRepository.GetCompetitor(event.CompetitorID)
	if err != nil {
		return err
	}
	competitor.SucceedShotsOnCurrentLap++
	shot := entity.Shot{
		Lap:    competitor.CurrentLap,
		Target: shotPayload.Target,
		Time:   event.Time,
	}

	err = b.competitorRepository.AddShot(event.CompetitorID, &shot)

	return nil
}

func (b *biathlonImpl) ProcessLeaveShootingRange(event entity.Event) error {
	competitor, err := b.competitorRepository.GetCompetitor(event.CompetitorID)
	if err != nil {
		return err
	}

	penalty := entity.Penalty{
		Lap:                competitor.CurrentLap,
		Time:               time.Time{},
		ValueOfMissedShots: 5 - competitor.SucceedShotsOnCurrentLap,
		AverageSpeed:       0,
	}

	err = b.competitorRepository.AddPenalty(event.CompetitorID, &penalty)
	if err != nil {
		return err
	}

	err = b.competitorRepository.MarkAsRacing(event.CompetitorID)
	if err != nil {
		return err
	}
	return nil
}
