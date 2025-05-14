package cmd

import (
    "fmt"
)

type Command interface {
    Run(args []string) error
}

type CommandFactory func(ctx *AppContext) Command

var registry = map[string]CommandFactory{}

func Register(name string, factory CommandFactory) {
    registry[name] = factory
}

func CreateCommand(name string, ctx *AppContext) (Command, error) {
    factory, ok := registry[name]
    if !ok {
        return nil, fmt.Errorf("comando '%s' não encontrado", name)
    }
    return factory(ctx), nil
}
