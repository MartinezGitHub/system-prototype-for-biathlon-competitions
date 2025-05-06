package parser

import (
	"bufio"
	"fmt"
	"github.com/MartinezGitHub/system-prototype-for-biathlon-competitions/internal/controller"
	"github.com/MartinezGitHub/system-prototype-for-biathlon-competitions/internal/parser/lexer"
	"go.uber.org/zap"
	"os"
)

type handlerFunc func(string) (string, error)

type InputParser struct {
	ctrl     controller.BiathlonController
	handlers map[string]handlerFunc
	logger   *zap.Logger
}

func New(logger *zap.Logger, ctrl controller.BiathlonController) *InputParser {
	return &InputParser{
		ctrl:   ctrl,
		logger: logger,
		handlers: map[string]handlerFunc{
			"1":  ctrl.HandleRegistration,
			"2":  ctrl.HandleSetStartTime,
			"3":  ctrl.HandleOnStartLine,
			"4":  ctrl.HandleStart,
			"5":  ctrl.HandleOnShootingRange,
			"6":  ctrl.HandleTargetHit,
			"7":  ctrl.HandleLeaveShootingRange,
			"8":  ctrl.HandleEnterPenalty,
			"9":  ctrl.HandleLeavePenalty,
			"10": ctrl.HandleFinishLap,
			"11": ctrl.HandleCannotContinue,
		},
	}
}

func (p *InputParser) ParseFile(inputPath string, outputPath string) error {
	file, err := os.Open(inputPath)
	if err != nil {
		return err
	}

	fileToWrite, err := os.Create(outputPath)
	if err != nil {
		return err
	}

	defer func() {
		closeErr := file.Close()
		if closeErr != nil {
			p.logger.Warn("failed to close input file", zap.Error(closeErr))
		}
		closeErr = fileToWrite.Close()
		if closeErr != nil {
			p.logger.Warn("failed to close output file", zap.Error(closeErr))
		}
	}()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		var str string
		str, err = p.parseLine(line)
		if err != nil {
			p.logger.Warn("Parse error", zap.String("line", line), zap.Error(err))
			continue
		}
		_, err = fileToWrite.WriteString(str + "\n")
		if err != nil {
			return err
		}
	}
	return scanner.Err()

}

func (p *InputParser) parseLine(line string) (string, error) {
	eventType, _ := lexer.SecondToken(line)
	if handler, exists := p.handlers[eventType]; exists {
		return handler(line)
	}
	return "", fmt.Errorf("unknown event type: %s", eventType)
}
