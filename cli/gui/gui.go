package gui

import (
	"fmt"
)

type App struct {
}

func NewApp() *App {
	return &App{}
}

func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s!", name)
}
