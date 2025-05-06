package biathlon

import (
	"github.com/MartinezGitHub/system-prototype-for-biathlon-competitions/internal/entity"
	"time"
)

func (b *biathlonImpl) ProcessEnterPenalty(event entity.Event) error {
	err := b.competitorRepository.MarkAsOnPenalty(event.CompetitorID)
	if err != nil {
		return err
	}
	competitor, err := b.competitorRepository.GetCompetitor(event.CompetitorID)
	if err != nil {
		return err
	}
	penalty := competitor.Penalties[len(competitor.Penalties)-1]

	penalty.Time = event.Time
	return nil
}

func (b *biathlonImpl) ProcessLeavePenalty(event entity.Event) error {
	err := b.competitorRepository.MarkAsRacing(event.CompetitorID)
	if err != nil {
		return err
	}

	competitor, err := b.competitorRepository.GetCompetitor(event.CompetitorID)
	if err != nil {
		return err
	}
	penalty := competitor.Penalties[len(competitor.Penalties)-1]

	diff := event.Time.Sub(penalty.Time)
	seconds := diff.Seconds()
	base := time.Date(0, 1, 1, 0, 0, 0, 0, time.UTC)

	penalty.Time = base.Add(diff)
	penalty.AverageSpeed = float64(b.config.PenaltyLen*penalty.ValueOfMissedShots) / seconds

	return nil
}
