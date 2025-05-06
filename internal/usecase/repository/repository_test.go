package repository

import (
	"github.com/MartinezGitHub/system-prototype-for-biathlon-competitions/config"
	"testing"
	"time"

	"github.com/MartinezGitHub/system-prototype-for-biathlon-competitions/internal/entity"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func TestAddCompetitor(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name string
		id   int
		want *entity.Competitor
	}{
		{
			name: "Add competitor 1",
			id:   1,
			want: &entity.Competitor{
				ID:           1,
				Status:       entity.Registered,
				LapTimes:     []*entity.LapTime{},
				Shots:        []*entity.Shot{},
				Disqualified: false,
			},
		},
	}

	repo := NewInMemory(zap.NewNop(), &config.Config{})
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := repo.AddCompetitor(tc.id)
			require.NoError(t, err)
			competitor, err := repo.GetCompetitor(tc.id)
			require.NoError(t, err)
			require.Equal(t, tc.want, competitor)
		})
	}
}

func TestAddCompetitorAlreadyExists(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name string
		id   int
		want error
	}{
		{
			name: "Add first competitor",
			id:   1,
			want: nil,
		},
		{
			name: "Add same competitor",
			id:   1,
			want: entity.ErrCompetitorExists,
		},
	}

	repo := NewInMemory(zap.NewNop(), &config.Config{})
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := repo.AddCompetitor(tc.id)
			require.Equal(t, tc.want, err)
		})
	}
}

func TestSetStartTime(t *testing.T) {
	t.Parallel()
	now := time.Now()
	testCases := []struct {
		name      string
		id        int
		startTime time.Time
		want      error
	}{
		{
			name:      "Set start time - success",
			id:        1,
			startTime: now,
			want:      nil,
		},
		{
			name:      "Set start time - not found",
			id:        999,
			startTime: now,
			want:      entity.ErrCompetitorNotFound,
		},
	}

	repo := NewInMemory(zap.NewNop(), &config.Config{})
	err := repo.AddCompetitor(1)
	require.NoError(t, err)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err = repo.SetStartTime(tc.id, tc.startTime)
			require.Equal(t, tc.want, err)
		})
	}
}

func TestMarkNotStarted(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name    string
		id      int
		message string
		want    error
	}{
		{
			name:    "Mark not started - success",
			id:      1,
			message: "test",
			want:    nil,
		},
		{
			name:    "Mark not started - not found",
			id:      999,
			message: "test",
			want:    entity.ErrCompetitorNotFound,
		},
	}

	repo := NewInMemory(zap.NewNop(), &config.Config{})
	err := repo.AddCompetitor(1)
	require.NoError(t, err)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err = repo.MarkNotStarted(tc.id, tc.message)
			require.Equal(t, tc.want, err)
		})
	}
}

func TestSetOnStartLine(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name string
		id   int
		want error
	}{
		{
			name: "Set on start line - success",
			id:   1,
			want: nil,
		},
		{
			name: "Set on start line - not found",
			id:   999,
			want: entity.ErrCompetitorNotFound,
		},
	}

	repo := NewInMemory(zap.NewNop(), &config.Config{})
	err := repo.AddCompetitor(1)
	require.NoError(t, err)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err = repo.SetOnStartLine(tc.id)
			require.Equal(t, tc.want, err)
		})
	}
}

func TestGetCompetitor(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name string
		id   int
		want error
	}{
		{
			name: "Exists",
			id:   1,
			want: nil,
		},
		{
			name: "Does not exist",
			id:   2,
			want: entity.ErrCompetitorNotFound,
		},
	}

	repo := NewInMemory(zap.NewNop(), &config.Config{})
	err := repo.AddCompetitor(1)
	require.NoError(t, err)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err = repo.GetCompetitor(tc.id)
			require.Equal(t, tc.want, err)
		})
	}
}

func TestUpdateCompetitor(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name       string
		competitor *entity.Competitor
		want       error
	}{
		{
			name: "Update existing",
			competitor: &entity.Competitor{
				ID:     1,
				Status: entity.Racing,
			},
			want: nil,
		},
		{
			name: "Update not existing",
			competitor: &entity.Competitor{
				ID:     999,
				Status: entity.Racing,
			},
			want: entity.ErrCompetitorNotFound,
		},
	}

	repo := NewInMemory(zap.NewNop(), &config.Config{})
	err := repo.AddCompetitor(1)
	require.NoError(t, err)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err = repo.UpdateCompetitor(tc.competitor)
			require.Equal(t, tc.want, err)
		})
	}
}

func TestMarkOnShootingRange(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name string
		id   int
		want error
	}{
		{
			name: "Mark on shooting - success",
			id:   1,
			want: nil,
		},
		{
			name: "Mark on shooting - not found",
			id:   999,
			want: entity.ErrCompetitorNotFound,
		},
	}

	repo := NewInMemory(zap.NewNop(), &config.Config{})
	err := repo.AddCompetitor(1)
	require.NoError(t, err)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err = repo.MarkOnShootingRange(tc.id)
			require.Equal(t, tc.want, err)
		})
	}
}

func TestAddShot(t *testing.T) {
	t.Parallel()
	now := time.Now()
	testCases := []struct {
		name string
		id   int
		shot *entity.Shot
		want error
	}{
		{
			name: "Add shot - success",
			id:   1,
			shot: &entity.Shot{
				Lap:    1,
				Range:  2,
				Target: 3,
				Time:   now,
			},
			want: nil,
		},
		{
			name: "Add shot - not found",
			id:   999,
			shot: &entity.Shot{
				Lap:    1,
				Range:  2,
				Target: 3,
				Time:   now,
			},
			want: entity.ErrCompetitorNotFound,
		},
	}

	repo := NewInMemory(zap.NewNop(), &config.Config{})
	err := repo.AddCompetitor(1)
	require.NoError(t, err)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err = repo.AddShot(tc.id, tc.shot)
			require.Equal(t, tc.want, err)
		})
	}
}

func TestMarkAsRacing(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name string
		id   int
		want error
	}{
		{
			name: "Mark as racing - success",
			id:   1,
			want: nil,
		},
		{
			name: "Mark as racing - not found",
			id:   999,
			want: entity.ErrCompetitorNotFound,
		},
	}

	repo := NewInMemory(zap.NewNop(), &config.Config{})
	err := repo.AddCompetitor(1)
	require.NoError(t, err)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err = repo.MarkAsRacing(tc.id)
			require.Equal(t, tc.want, err)
		})
	}
}

func TestAddPenalty(t *testing.T) {
	t.Parallel()
	now := time.Now()
	testCases := []struct {
		name    string
		id      int
		penalty *entity.Penalty
		want    error
	}{
		{
			name: "Add penalty - success",
			id:   1,
			penalty: &entity.Penalty{
				Lap:                1,
				ValueOfMissedShots: 2,
				Time:               now,
				AverageSpeed:       10.5,
			},
			want: nil,
		},
		{
			name: "Add penalty - not found",
			id:   999,
			penalty: &entity.Penalty{
				Lap:                1,
				ValueOfMissedShots: 2,
				Time:               now,
				AverageSpeed:       10.5,
			},
			want: entity.ErrCompetitorNotFound,
		},
	}

	repo := NewInMemory(zap.NewNop(), &config.Config{})
	err := repo.AddCompetitor(1)
	require.NoError(t, err)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err = repo.AddPenalty(tc.id, tc.penalty)
			require.Equal(t, tc.want, err)
		})
	}
}

func TestMarkAsOnPenalty(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name string
		id   int
		want error
	}{
		{
			name: "Mark on penalty - success",
			id:   1,
			want: nil,
		},
		{
			name: "Mark on penalty - not found",
			id:   999,
			want: entity.ErrCompetitorNotFound,
		},
	}

	repo := NewInMemory(zap.NewNop(), &config.Config{})
	err := repo.AddCompetitor(1)
	require.NoError(t, err)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err = repo.MarkAsOnPenalty(tc.id)
			require.Equal(t, tc.want, err)
		})
	}
}

func TestMarkNotFinished(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name    string
		id      int
		message string
		want    error
	}{
		{
			name:    "Mark not finished - success",
			id:      1,
			message: "test",
			want:    nil,
		},
		{
			name:    "Mark not finished - not found",
			id:      999,
			message: "test",
			want:    entity.ErrCompetitorNotFound,
		},
	}

	repo := NewInMemory(zap.NewNop(), &config.Config{})
	err := repo.AddCompetitor(1)
	require.NoError(t, err)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err = repo.MarkNotFinished(tc.id, tc.message)
			require.Equal(t, tc.want, err)
		})
	}
}

func TestGetCompetitors(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name  string
		ids   []int
		count int
	}{
		{
			name:  "Get all",
			ids:   []int{1, 2, 3},
			count: 3,
		},
	}

	repo := NewInMemory(zap.NewNop(), &config.Config{})
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			for _, id := range tc.ids {
				err := repo.AddCompetitor(id)
				require.NoError(t, err)
			}
			competitors := repo.GetCompetitors()
			require.Equal(t, tc.count, len(competitors))
		})
	}
}
