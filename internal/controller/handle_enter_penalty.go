package controller

import (
	"fmt"
	"github.com/MartinezGitHub/system-prototype-for-biathlon-competitions/internal/entity"
	"github.com/MartinezGitHub/system-prototype-for-biathlon-competitions/internal/parser/lexer"
)

func (i *implementation) HandleEnterPenalty(rest string) (string, error) {
	tokens := lexer.AllTokens(rest)
	eventTime, _, competitorID, err := parseBaseEvent(tokens)
	if err != nil {
		return "", fmt.Errorf("enter penalty: %w", err)
	}

	event := entity.Event{
		Time:         eventTime,
		Type:         entity.EventEnterPenalty,
		CompetitorID: competitorID,
	}

	return fmt.Sprintf("[%s] The competitor(%d) entered the penalty laps",
			event.Time.Format("15:04:05.000"),
			event.CompetitorID),
		i.penaltyUseCase.ProcessEnterPenalty(event)
}
