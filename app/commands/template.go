package commands

import (
	"fmt"

	"github.com/lem3s/fg/app/cmd"
)

type HelloCommand struct {
	Ctx *cmd.AppContext
}

func (h *HelloCommand) Run(args []string) error {
	//aqui você retornaria uma info para o usuario
	h.Ctx.Interactor.Info(fmt.Sprintf("Hello %s!", h.Ctx.Config.GetString("jar")))
	return nil
}

func init() {
	cmd.Register("hello", func(ctx *cmd.AppContext) cmd.Command {
		return &HelloCommand{Ctx: ctx}
	})
}
