package repository

import (
	"github.com/MartinezGitHub/system-prototype-for-biathlon-competitions/config"
	"github.com/MartinezGitHub/system-prototype-for-biathlon-competitions/internal/entity"
	"go.uber.org/zap"
	"time"
)

var _ CompetitorRepository = (*InMemoryImpl)(nil)

type InMemoryImpl struct {
	competitors map[int]*entity.Competitor
	config      *config.Config
	logger      *zap.Logger
}

func NewInMemory(logger *zap.Logger, config *config.Config) *InMemoryImpl {
	return &InMemoryImpl{
		competitors: make(map[int]*entity.Competitor),
		config:      config,
		logger:      logger,
	}
}

func (i *InMemoryImpl) GetCompetitors() []*entity.Competitor {
	ret := make([]*entity.Competitor, 0, len(i.competitors))
	for _, competitor := range i.competitors {
		ret = append(ret, competitor)
	}
	return ret
}

func (i *InMemoryImpl) MarkNotFinished(id int, message string) error {
	if _, ok := i.competitors[id]; !ok {
		return entity.ErrCompetitorNotFound
	}

	i.competitors[id].Status = entity.NotFinished
	i.competitors[id].Comment = message
	return nil
}

func (i *InMemoryImpl) MarkAsOnPenalty(id int) error {
	if _, ok := i.competitors[id]; !ok {
		return entity.ErrCompetitorNotFound
	}

	i.competitors[id].Status = entity.OnPenalty
	return nil
}

func (i *InMemoryImpl) AddPenalty(id int, penalty *entity.Penalty) error {
	competitor, ok := i.competitors[id]

	if !ok {
		return entity.ErrCompetitorNotFound
	}

	competitor.Penalties = append(competitor.Penalties, penalty)

	return nil
}

func (i *InMemoryImpl) MarkAsRacing(id int) error {
	if _, ok := i.competitors[id]; !ok {
		return entity.ErrCompetitorNotFound
	}

	i.competitors[id].Status = entity.Racing
	return nil
}

func (i *InMemoryImpl) AddShot(id int, shot *entity.Shot) error {

	competitor, ok := i.competitors[id]

	if !ok {
		return entity.ErrCompetitorNotFound
	}

	competitor.Shots = append(competitor.Shots, shot)

	return nil
}

func (i *InMemoryImpl) MarkOnShootingRange(id int) error {

	competitor, ok := i.competitors[id]

	if !ok {
		return entity.ErrCompetitorNotFound
	}

	competitor.Status = entity.OnShooting
	return nil
}

func (i *InMemoryImpl) SetOnStartLine(id int) error {

	if _, ok := i.competitors[id]; !ok {
		return entity.ErrCompetitorNotFound
	}

	i.competitors[id].Status = entity.OnStartLine
	return nil
}

func (i *InMemoryImpl) MarkNotStarted(id int, message string) error {

	if _, ok := i.competitors[id]; !ok {
		return entity.ErrCompetitorNotFound
	}

	i.competitors[id].Status = entity.NotStarted
	i.competitors[id].Comment = message

	return nil
}

func (i *InMemoryImpl) SetStartTime(id int, startTime time.Time) error {
	if _, exists := i.competitors[id]; !exists {
		return entity.ErrCompetitorNotFound
	}
	i.competitors[id].StartTime = startTime
	return nil
}

func (i *InMemoryImpl) AddCompetitor(id int) error {

	if _, exists := i.competitors[id]; exists {
		return entity.ErrCompetitorExists
	}

	i.competitors[id] = &entity.Competitor{
		ID:           id,
		StartTime:    time.Time{},
		Status:       entity.Registered,
		LapTimes:     make([]*entity.LapTime, 0),
		Shots:        make([]*entity.Shot, 0),
		Disqualified: false,
	}
	return nil

}

func (i *InMemoryImpl) GetCompetitor(id int) (*entity.Competitor, error) {

	competitor, exists := i.competitors[id]
	if !exists {
		return nil, entity.ErrCompetitorNotFound
	}
	return competitor, nil
}

func (i *InMemoryImpl) UpdateCompetitor(competitor *entity.Competitor) error {

	if _, exists := i.competitors[competitor.ID]; !exists {
		return entity.ErrCompetitorNotFound
	}

	i.competitors[competitor.ID] = competitor
	return nil
}
