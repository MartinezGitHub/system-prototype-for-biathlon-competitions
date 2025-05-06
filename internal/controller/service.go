package controller

import (
	"github.com/MartinezGitHub/system-prototype-for-biathlon-competitions/config"
	"github.com/MartinezGitHub/system-prototype-for-biathlon-competitions/internal/usecase/biathlon"
	"go.uber.org/zap"
)

type BiathlonController interface {
	HandleRegistration(rest string) (string, error)
	HandleSetStartTime(rest string) (string, error)
	HandleOnStartLine(rest string) (string, error)
	HandleStart(rest string) (string, error)
	HandleOnShootingRange(rest string) (string, error)
	HandleTargetHit(rest string) (string, error)
	HandleLeaveShootingRange(rest string) (string, error)
	HandleEnterPenalty(rest string) (string, error)
	HandleLeavePenalty(rest string) (string, error)
	HandleFinishLap(rest string) (string, error)
	HandleCannotContinue(rest string) (string, error)
	GenerateReport(path string) error
}

var _ BiathlonController = (*implementation)(nil)

type implementation struct {
	logger             *zap.Logger
	config             *config.Config
	fireRangeUseCase   biathlon.FiringRangeUseCase
	penaltyUseCase     biathlon.PenaltyUseCase
	competitorsUseCase biathlon.CompetitorsUseCase
	eventUseCase       biathlon.PenaltyUseCase
	reportUseCase      biathlon.ReportUseCase
}

func New(
	logger *zap.Logger,
	config *config.Config,
	fireRangeUseCase biathlon.FiringRangeUseCase,
	penaltyUseCase biathlon.PenaltyUseCase,
	competitorsUseCase biathlon.CompetitorsUseCase,
	eventUseCase biathlon.PenaltyUseCase,
	reportUseCase biathlon.ReportUseCase,
) *implementation {
	return &implementation{
		logger:             logger,
		config:             config,
		fireRangeUseCase:   fireRangeUseCase,
		penaltyUseCase:     penaltyUseCase,
		competitorsUseCase: competitorsUseCase,
		eventUseCase:       eventUseCase,
		reportUseCase:      reportUseCase,
	}
}
