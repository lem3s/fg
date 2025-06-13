package logger

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

type LoggerService struct {
	logsPath string
}

func NewLoggerService() *LoggerService {
	logsPath := filepath.Join(".", "logs")

	if _, err := os.Stat(logsPath); os.IsNotExist(err) {
		file, err := os.Create(logsPath)
		if err != nil {
			panic(err)
		}
		file.Close()
	}

	return &LoggerService{
		logsPath: logsPath,
	}
}

func (l *LoggerService) Info(message string) error {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	logMessage := fmt.Sprintf("[%s] INFO: %s\n", timestamp, message)

	file, err := os.OpenFile(l.logsPath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open logs file: %w", err)
	}
	defer file.Close()

	if _, err := file.WriteString(logMessage); err != nil {
		return fmt.Errorf("failed to write to logs file: %w", err)
	}

	return nil
}

func (l *LoggerService) Warn(message string) error {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	logMessage := fmt.Sprintf("[%s] WARNING: %s\n", timestamp, message)

	file, err := os.OpenFile(l.logsPath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open logs file: %w", err)
	}
	defer file.Close()

	if _, err := file.WriteString(logMessage); err != nil {
		return fmt.Errorf("failed to write to logs file: %w", err)
	}

	return nil
}

func (l *LoggerService) Error(message string) error {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	logMessage := fmt.Sprintf("[%s] ERROR: %s\n", timestamp, message)

	file, err := os.OpenFile(l.logsPath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open logs file: %w", err)
	}
	defer file.Close()

	if _, err := file.WriteString(logMessage); err != nil {
		return fmt.Errorf("failed to write to logs file: %w", err)
	}

	return nil
}
