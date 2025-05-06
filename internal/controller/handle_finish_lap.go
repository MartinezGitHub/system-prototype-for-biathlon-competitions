package controller

import (
	"fmt"
	"github.com/MartinezGitHub/system-prototype-for-biathlon-competitions/internal/entity"
	"github.com/MartinezGitHub/system-prototype-for-biathlon-competitions/internal/parser/lexer"
)

func (i *implementation) HandleFinishLap(rest string) (string, error) {
	tokens := lexer.AllTokens(rest)
	eventTime, _, competitorID, err := parseBaseEvent(tokens)
	if err != nil {
		return "", fmt.Errorf("finish lap: %w", err)
	}

	event := entity.Event{
		Time:         eventTime,
		Type:         entity.EventFinishLap,
		CompetitorID: competitorID,
	}
	return fmt.Sprintf("[%s] The competitor(%d) ended the main lap",
		event.Time.Format("15:04:05.000"),
		event.CompetitorID), i.competitorsUseCase.ProcessFinishLap(event)
}
