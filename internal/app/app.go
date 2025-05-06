package app

import (
	"github.com/MartinezGitHub/system-prototype-for-biathlon-competitions/config"
	"github.com/MartinezGitHub/system-prototype-for-biathlon-competitions/internal/controller"
	inputParser "github.com/MartinezGitHub/system-prototype-for-biathlon-competitions/internal/parser"
	"github.com/MartinezGitHub/system-prototype-for-biathlon-competitions/internal/usecase/biathlon"
	"github.com/MartinezGitHub/system-prototype-for-biathlon-competitions/internal/usecase/repository"
	"go.uber.org/zap"
)

const inputFilePath = "test_data/input.txt"
const outputFilePath = "test_data/output.txt"
const resultsPath = "test_data/results.txt"

func Run(logger *zap.Logger, cfg *config.Config) {
	repo := repository.NewInMemory(logger, cfg)

	useCases := biathlon.New(logger, cfg, repo)

	ctrl := controller.New(logger, cfg, useCases, useCases, useCases, useCases, useCases)

	parser := inputParser.New(logger, ctrl)

	if err := parser.ParseFile(inputFilePath, outputFilePath); err != nil {
		logger.Error(err.Error())
	}

	if err := ctrl.GenerateReport(resultsPath); err != nil {
		logger.Error(err.Error())
	}

	return

}
