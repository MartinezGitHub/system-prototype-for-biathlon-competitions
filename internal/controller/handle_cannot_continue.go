package controller

import (
	"fmt"
	"github.com/MartinezGitHub/system-prototype-for-biathlon-competitions/internal/entity"
	"github.com/MartinezGitHub/system-prototype-for-biathlon-competitions/internal/parser/lexer"
	"strings"
)

func (i *implementation) HandleCannotContinue(rest string) (string, error) {
	tokens := lexer.AllTokens(rest)
	eventTime, _, competitorID, err := parseBaseEvent(tokens)
	if err != nil {
		return "", fmt.Errorf("cannot continue: %w", err)
	}

	if len(tokens) < 4 {
		return "", fmt.Errorf("missing reason")
	}

	text := strings.Join(tokens[3:], " ")
	event := entity.Event{
		Time:         eventTime,
		Type:         entity.EventCannotContinue,
		CompetitorID: competitorID,
		Payload: entity.CommentPayload{
			Text: text,
		},
	}
	return fmt.Sprintf("[%s] The competitor(%d) can`t continue: %s",
			eventTime.Format("15:04:05.000"),
			event.CompetitorID, text),
		i.competitorsUseCase.ProcessCannotContinue(event)
}
