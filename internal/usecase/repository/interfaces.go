package repository

import (
	"github.com/MartinezGitHub/system-prototype-for-biathlon-competitions/internal/entity"
	"time"
)

type (
	CompetitorRepository interface {
		AddCompetitor(id int) error
		SetStartTime(id int, startTime time.Time) error
		MarkNotStarted(id int, message string) error
		SetOnStartLine(id int) error
		GetCompetitor(id int) (*entity.Competitor, error)
		UpdateCompetitor(competitor *entity.Competitor) error
		MarkOnShootingRange(id int) error
		AddShot(id int, shot *entity.Shot) error
		MarkAsRacing(id int) error
		AddPenalty(id int, penalty *entity.Penalty) error
		MarkAsOnPenalty(id int) error
		MarkNotFinished(id int, message string) error
		GetCompetitors() []*entity.Competitor
	}
)
