package biathlon

import (
	"fmt"
	"github.com/MartinezGitHub/system-prototype-for-biathlon-competitions/config"
	"github.com/MartinezGitHub/system-prototype-for-biathlon-competitions/internal/entity"
	"github.com/MartinezGitHub/system-prototype-for-biathlon-competitions/internal/usecase/repository"
	"go.uber.org/zap"
	"strconv"
	"strings"
	"time"
)

type (
	ReportUseCase interface {
		GenerateReportFromRepository(path string) error
	}

	PenaltyUseCase interface {
		ProcessEnterPenalty(event entity.Event) error
		ProcessLeavePenalty(event entity.Event) error

		CalculatePenaltyTime(competitorID int) (time.Duration, error)
	}

	FiringRangeUseCase interface {
		ProcessOnShootingRange(event entity.Event) error
		ProcessTargetHit(event entity.Event) error
		ProcessLeaveShootingRange(event entity.Event) error
	}

	CompetitorsUseCase interface {
		ProcessRegistration(event entity.Event) error
		ProcessSetStartTime(event entity.Event) error
		ProcessOnStartLine(event entity.Event) error
		ProcessStart(event entity.Event) error
		ProcessFinishLap(event entity.Event) error
		ProcessCannotContinue(event entity.Event) error
	}
)

var _ ReportUseCase = (*biathlonImpl)(nil)
var _ PenaltyUseCase = (*biathlonImpl)(nil)
var _ FiringRangeUseCase = (*biathlonImpl)(nil)
var _ CompetitorsUseCase = (*biathlonImpl)(nil)

type biathlonImpl struct {
	logger               *zap.Logger
	config               *config.Config
	competitorRepository repository.CompetitorRepository
}

func New(
	logger *zap.Logger,
	config *config.Config,
	competitorRepository repository.CompetitorRepository,
) *biathlonImpl {
	return &biathlonImpl{
		logger:               logger,
		config:               config,
		competitorRepository: competitorRepository,
	}
}

func parseCustomDuration(durationStr string) (time.Duration, error) {
	parts := strings.Split(durationStr, ":")
	if len(parts) != 3 {
		return 0, fmt.Errorf("invalid duration format, expected HH:MM:SS")
	}

	hours, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, err
	}

	minutes, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0, err
	}

	seconds, err := strconv.Atoi(parts[2])
	if err != nil {
		return 0, err
	}

	totalSeconds := hours*3600 + minutes*60 + seconds
	return time.Duration(totalSeconds) * time.Second, nil
}

//func formatDuration(d time.Duration) string {
//	hours := int(d.Hours())
//	minutes := int(d.Minutes()) % 60
//	seconds := int(d.Seconds()) % 60
//	milliseconds := (d.Nanoseconds() % 1e9) / 1e6
//
//	return fmt.Sprintf("[%02d:%02d:%02d.%03d]", hours, minutes, seconds, milliseconds)
//}
