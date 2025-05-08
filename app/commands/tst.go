package commands

import (
    "fmt"
    "github.com/lem3s/fg/app/cmd"
)

type HelloCommand struct {
    Ctx *cmd.AppContext
}

func (h *HelloCommand) Run(args []string) {
    nome := h.Ctx.Config.GetString("nome")
    if nome == "" && len(args) > 0 {
        nome = args[0]
    }
    if nome == "" {
        nome = "Mundo"
    }
    fmt.Printf("Olá, %s!\n", nome)
}

func init() {
    cmd.Register("hello", func(ctx *cmd.AppContext) cmd.Command {
        return &HelloCommand{Ctx: ctx}
    })
}
