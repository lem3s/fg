package commands

import (
	"fmt"

	"github.com/lem3s/fg/app/cmd"
)

type LogsCommand struct {
	Ctx *cmd.AppContext
}

func (l *LogsCommand) Run(args []string) error {
	//aqui você retornaria uma info para o usuario
	l.Ctx.Interactor.Info(fmt.Sprintf("Hello %s!", l.Ctx.Config.GetString("jar")))
	return nil
}

func init() {
	cmd.Register("logs", func(ctx *cmd.AppContext) cmd.Command {
		return &LogsCommand{Ctx: ctx}
	})
}
