package biathlon

import (
	"fmt"
	"github.com/MartinezGitHub/system-prototype-for-biathlon-competitions/internal/entity"
	"time"
)

func (b *biathlonImpl) ProcessRegistration(event entity.Event) error {
	err := b.competitorRepository.AddCompetitor(event.CompetitorID)
	if err != nil {
		return err
	}

	return nil
}

func (b *biathlonImpl) ProcessSetStartTime(event entity.Event) error {
	timePayload, ok := event.Payload.(entity.StartTimePayload)
	if !ok {
		return fmt.Errorf("invalid payload type for SetStartTime event")
	}

	err := b.competitorRepository.SetStartTime(event.CompetitorID, timePayload.Time)
	if err != nil {
		return err
	}

	return nil
}

func (b *biathlonImpl) ProcessOnStartLine(event entity.Event) error {
	err := b.competitorRepository.SetOnStartLine(event.CompetitorID)
	if err != nil {
		return err
	}

	return nil
}

func (b *biathlonImpl) ProcessStart(event entity.Event) error {
	competitor, err := b.competitorRepository.GetCompetitor(event.CompetitorID)
	if err != nil {
		return err
	}

	parsedDuration, err := parseCustomDuration(b.config.StartDelta)
	if err != nil {
		return err
	}

	if event.Time.Before(competitor.StartTime) ||
		event.Time.After(competitor.StartTime.Add(parsedDuration)) {
		err = b.competitorRepository.MarkNotStarted(event.CompetitorID, "Disqualified")
		if err != nil {
			return err
		}
		return nil
	}

	err = b.competitorRepository.MarkAsRacing(event.CompetitorID)
	competitor, err = b.competitorRepository.GetCompetitor(event.CompetitorID)
	if err != nil {
		return err
	}
	competitor.CurrentLap = 0
	lapTime := entity.LapTime{
		Number:       0,
		Time:         event.Time,
		AverageSpeed: 0,
		Active:       true,
	}
	competitor.LapTimes = append(competitor.LapTimes, &lapTime)

	return nil
}

func (b *biathlonImpl) ProcessFinishLap(event entity.Event) error {
	competitor, err := b.competitorRepository.GetCompetitor(event.CompetitorID)
	if err != nil {
		return err
	}

	competitor.FinishTime = event.Time

	lapTime := competitor.LapTimes[len(competitor.LapTimes)-1]

	diff := event.Time.Sub(lapTime.Time)
	seconds := diff.Seconds()
	base := time.Date(0, 1, 1, 0, 0, 0, 0, time.UTC)

	lapTime.Time = base.Add(diff)

	lapTime.AverageSpeed = float64(b.config.LapLen) / seconds

	err = b.competitorRepository.MarkAsRacing(event.CompetitorID)
	if err != nil {
		return err
	}

	competitor.CurrentLap++

	newLapTime := entity.LapTime{
		Number:       competitor.CurrentLap,
		Time:         event.Time,
		AverageSpeed: 0,
		Active:       false,
	}
	competitor.LapTimes = append(competitor.LapTimes, &newLapTime)

	return nil
}

func (b *biathlonImpl) ProcessCannotContinue(event entity.Event) error {
	commentPayload, ok := event.Payload.(entity.CommentPayload)
	if !ok {
		return fmt.Errorf("invalid payload type for SetStartTime event")
	}
	err := b.competitorRepository.MarkNotFinished(event.CompetitorID, commentPayload.Text)
	if err != nil {
		return err
	}
	return nil
}
