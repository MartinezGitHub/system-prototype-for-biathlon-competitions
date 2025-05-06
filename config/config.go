package config

import (
	"encoding/json"
	"github.com/pkg/errors"

	"fmt"
	"go.uber.org/zap"
	"os"
)

const DefaultValueOfRegisteredCompetitors = 0
const DefaultValueOfCompetitorsWithoutSetTime = 0

type Config struct {
	Laps        int    `json:"laps" validate:"required,min=1"`
	LapLen      int    `json:"lapLen" validate:"required,min=1"`
	PenaltyLen  int    `json:"penaltyLen" validate:"required,min=1"`
	FiringLines int    `json:"firingLines" validate:"required,min=1"`
	Start       string `json:"start" validate:"required"`
	StartDelta  string `json:"startDelta" validate:"required"`

	ValueOfRegisteredCompetitors     int
	ValueOfCompetitorsWithoutSetTime int
}

func NewConfig(logger *zap.Logger, path string) *Config {
	config, err := LoadFromFile(path)
	if err != nil {
		logger.Error("failed to load config data from file: ", zap.Error(err))
		return nil
	}

	config.ValueOfRegisteredCompetitors = DefaultValueOfRegisteredCompetitors
	config.ValueOfCompetitorsWithoutSetTime = DefaultValueOfCompetitorsWithoutSetTime

	return config
}

func LoadFromFile(path string) (*Config, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read config file")
	}

	var cfg Config
	if unmarshalErr := json.Unmarshal(file, &cfg); unmarshalErr != nil {
		return nil, errors.Wrap(unmarshalErr, "failed to unmarshal config file")
	}

	return &cfg, nil
}

func PrintConfig(config *Config) {
	fmt.Println("Laps:", config.Laps)
	fmt.Println("LapLen:", config.LapLen)
	fmt.Println("PenaltyLen:", config.PenaltyLen)
	fmt.Println("FiringLines:", config.FiringLines)
	fmt.Println("Start:", config.Start)
	fmt.Println("StartDelta:", config.StartDelta)
}
