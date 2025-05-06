package controller

import (
	"fmt"
	"github.com/MartinezGitHub/system-prototype-for-biathlon-competitions/internal/entity"
	"github.com/MartinezGitHub/system-prototype-for-biathlon-competitions/internal/parser/lexer"
	"strconv"
)

func (i *implementation) HandleOnShootingRange(rest string) (string, error) {
	tokens := lexer.AllTokens(rest)
	eventTime, _, competitorID, err := parseBaseEvent(tokens)
	if err != nil {
		return "", fmt.Errorf("on shooting range: %w", err)
	}

	if len(tokens) < 3 {
		return "", fmt.Errorf("missing range number")
	}

	rangeNum, err := strconv.Atoi(tokens[2])
	if err != nil {
		return "", fmt.Errorf("invalid range number: %w", err)
	}

	event := entity.Event{
		Time:         eventTime,
		Type:         entity.EventOnShootingRange,
		CompetitorID: competitorID,
		Payload: entity.RangePayload{
			Range: rangeNum,
		},
	}
	return fmt.Sprintf("[%s] The competitor(%d) is on the firing range(%d)",
		event.Time.Format("15:04:05.000"),
		event.CompetitorID, rangeNum), i.fireRangeUseCase.ProcessOnShootingRange(event)
}
