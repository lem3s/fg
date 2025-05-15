package cmd

import (
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
	Interactor     InteractionHandler
	SuccessMessage string
}

func NewAppContext(cfg *viper.Viper) *AppContext {
	return &AppContext{
		Config:     cfg,
		Interactor: &logger.DefaultInteractionHandler{},
	}
}
