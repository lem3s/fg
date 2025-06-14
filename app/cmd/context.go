package cmd

import (
	"runtime"
	"slices"

	"github.com/lem3s/fg/app/logger"
	"github.com/spf13/viper"
)

type InteractionHandler interface {
	Prompt(message string) string
	Confirm(message string) bool
	Info(message string)
	Warn(message string)
	Error(message string)
}

type AppContext struct {
	Config         *viper.Viper
	OS             string
	FgHome         string
	LogLevel       string
	Interactor     InteractionHandler
	SuccessMessage string
}

func NewAppContext(cfg *viper.Viper, FgHome string, LogLevel string) *AppContext {
	return &AppContext{Config: cfg, OS: runtime.GOOS, FgHome: FgHome, LogLevel: LogLevel, Interactor: &logger.DefaultInteractionHandler{}}
}

func IsVersionDeppendant(commandName string) bool {
	versionDeppendantCommands := []string{"install", "uninstall", "config", "start"}

	return slices.Contains(versionDeppendantCommands, commandName)
}
