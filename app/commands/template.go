package commands

import (
	"fmt"

	"github.com/lem3s/fg/app/cmd"
)

type HelloCommand struct {
	Ctx *cmd.AppContext
}

func (h *HelloCommand) Run(args []string) error {
	nome := h.Ctx.Config.GetString("nome")
	fmt.Printf("Olá, %s!\n", nome)
	return nil
}

func init() {
	cmd.Register("hello", func(ctx *cmd.AppContext) cmd.Command {
		return &HelloCommand{Ctx: ctx}
	})
}
