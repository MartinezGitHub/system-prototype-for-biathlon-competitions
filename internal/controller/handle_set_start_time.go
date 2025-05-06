package controller

import (
	"fmt"
	"github.com/MartinezGitHub/system-prototype-for-biathlon-competitions/internal/entity"
	"github.com/MartinezGitHub/system-prototype-for-biathlon-competitions/internal/parser/lexer"
	"time"
)

func (i *implementation) HandleSetStartTime(rest string) (string, error) {
	tokens := lexer.AllTokens(rest)
	eventTime, _, competitorID, err := parseBaseEvent(tokens)
	if err != nil {
		return "", fmt.Errorf("set start time: %w", err)
	}

	if len(tokens) < 3 {
		return "", fmt.Errorf("missing start time value")
	}

	startTime, err := time.Parse("15:04:05.000", tokens[3])
	if err != nil {
		return "", fmt.Errorf("invalid start time value: %w", err)
	}

	event := entity.Event{
		Time:         eventTime,
		Type:         entity.EventSetStartTime,
		CompetitorID: competitorID,
		Payload: entity.StartTimePayload{
			Time: startTime,
		},
	}
	return fmt.Sprintf("[%s] The start time for the competitor(%d) was set by a draw to %s",
		event.Time.Format("15:04:05.000"),
		event.CompetitorID,
		startTime.Format("15:04:05.000")), i.competitorsUseCase.ProcessSetStartTime(event)
}
