package main

import (
	"github.com/MartinezGitHub/system-prototype-for-biathlon-competitions/config"
	"github.com/MartinezGitHub/system-prototype-for-biathlon-competitions/internal/app"
	"go.uber.org/zap"
	"log"
)

func main() {
	var logger *zap.Logger
	logger, err := zap.NewProduction()

	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}

	cfg := config.NewConfig(logger, "config/config.json")

	app.Run(logger, cfg)
}
