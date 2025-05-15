package main

import (
	"os"

	"github.com/lem3s/fg/app"
	"github.com/lem3s/fg/app/cmd"
	_ "github.com/lem3s/fg/app/commands"
	"github.com/lem3s/fg/app/handlers"
)

func main() {
	commandName := os.Args[1]
	args := os.Args[2:]
	cfg := app.GetConfig()
	ctx := cmd.NewAppContext(cfg)
	handlers.HandleParams(args, ctx)

	command, _ := cmd.CreateCommand(commandName, ctx)

	err := command.Run(args)
	handlers.HandleCallback(err, ctx)
}
