package controller

import (
	"fmt"
	"github.com/MartinezGitHub/system-prototype-for-biathlon-competitions/internal/entity"
	"github.com/MartinezGitHub/system-prototype-for-biathlon-competitions/internal/parser/lexer"
	"strconv"
)

func (i *implementation) HandleTargetHit(rest string) (string, error) {
	tokens := lexer.AllTokens(rest)
	eventTime, _, competitorID, err := parseBaseEvent(tokens)
	if err != nil {
		return "", fmt.Errorf("target hit: %w", err)
	}

	if len(tokens) < 4 {
		return "", fmt.Errorf("missing target number")
	}

	targetNum, err := strconv.Atoi(tokens[3])
	if err != nil {
		return "", fmt.Errorf("invalid target number: %w", err)
	}

	event := entity.Event{
		Time:         eventTime,
		Type:         entity.EventTargetHit,
		CompetitorID: competitorID,
		Payload: entity.TargetHitPayload{
			Target: targetNum,
		},
	}
	return fmt.Sprintf("[%s] The target(%d) has been hit by competitor(%d)",
		event.Time.Format("15:04:05.000"),
		targetNum,
		event.CompetitorID), i.fireRangeUseCase.ProcessTargetHit(event)
}
