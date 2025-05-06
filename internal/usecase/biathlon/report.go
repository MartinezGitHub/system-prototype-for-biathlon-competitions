package biathlon

import (
	"fmt"
	"github.com/MartinezGitHub/system-prototype-for-biathlon-competitions/internal/entity"
	"os"
	"sort"
	"time"
)

func (b *biathlonImpl) GenerateReportFromRepository(path string) error {
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer func() {
		closeErr := file.Close()
		if closeErr != nil {
			b.logger.Warn("failed to close file")
		}
	}()

	competitors := b.competitorRepository.GetCompetitors()

	var notStartedFinished []*entity.Competitor
	var otherCompetitors []*entity.Competitor

	for _, competitor := range competitors {
		if competitor.Status == entity.NotStarted || competitor.Status == entity.NotFinished {
			notStartedFinished = append(notStartedFinished, competitor)
		} else {
			otherCompetitors = append(otherCompetitors, competitor)
		}
	}

	sort.Slice(otherCompetitors, func(i, j int) bool {
		diffI := otherCompetitors[i].FinishTime.Sub(otherCompetitors[i].StartTime)
		diffJ := otherCompetitors[j].FinishTime.Sub(otherCompetitors[j].StartTime)
		return diffI < diffJ
	})

	sortedCompetitors := append(otherCompetitors, notStartedFinished...)

	for _, competitor := range sortedCompetitors {
		var statusPart string
		if competitor.Status == entity.NotStarted || competitor.Status == entity.NotFinished {
			statusPart = fmt.Sprintf("%s", competitor.Status)
		} else {
			diff := competitor.FinishTime.Sub(competitor.StartTime)
			base := time.Date(0, 1, 1, 0, 0, 0, 0, time.UTC)

			statusPart = fmt.Sprintf("%s", base.Add(diff).Format("15:04:05.000"))
		}

		entityPart := fmt.Sprintf("[%s] %d", statusPart, competitor.ID)

		lapsInfo := make([]string, 0, len(competitor.LapTimes))
		for _, lap := range competitor.LapTimes {
			if lap.Active {
				lapsInfo = append(lapsInfo, fmt.Sprintf("{%s, %.3f}",
					lap.Time.Format("15:04:05.000"), lap.AverageSpeed))
			} else {
				if competitor.Status == entity.NotStarted || competitor.Status == entity.NotFinished {
					lapsInfo = append(lapsInfo, fmt.Sprint("{,}"))
				}
			}
		}

		lapsStr := fmt.Sprintf("[%s]", joinStrings(lapsInfo))

		penaltiesInfo := make([]string, 0, len(competitor.Penalties))
		for _, penalty := range competitor.Penalties {
			penaltiesInfo = append(penaltiesInfo, fmt.Sprintf("{%s, %.3f, %d/5}",
				penalty.Time.Format("15:04:05.000"), penalty.AverageSpeed, 5-penalty.ValueOfMissedShots))
		}
		penaltiesStr := fmt.Sprintf("[%s]", joinStrings(penaltiesInfo))

		line := fmt.Sprintf("%s %s %s\n", entityPart, lapsStr, penaltiesStr)
		if _, err = file.WriteString(line); err != nil {
			return fmt.Errorf("failed to write to file: %v", err)
		}
	}

	return nil
}

func joinStrings(items []string) string {
	result := ""
	for i, item := range items {
		if i > 0 {
			result += ", "
		}
		result += item
	}
	return result
}
