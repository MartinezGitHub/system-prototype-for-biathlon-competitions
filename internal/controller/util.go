package controller

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func parseBaseEvent(tokens []string) (time.Time, int, int, error) {
	if len(tokens) < 3 {
		return time.Time{}, 0, 0, fmt.Errorf("expected at least 3 tokens")
	}

	eventTime, err := time.Parse("15:04:05.000", strings.Trim(tokens[0], "[]"))
	if err != nil {
		return time.Time{}, 0, 0, fmt.Errorf("invalid time format: %w", err)
	}

	eventID, err := strconv.Atoi(tokens[1])
	if err != nil {
		return time.Time{}, 0, 0, fmt.Errorf("invalid event ID: %w", err)
	}

	competitorID, err := strconv.Atoi(tokens[2])
	if err != nil {
		return time.Time{}, 0, 0, fmt.Errorf("invalid competitor ID: %w", err)
	}
	return eventTime, eventID, competitorID, nil
}
