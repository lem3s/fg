package commands

import (
	"github.com/lem3s/fg/app/cmd"
)

type InstallCmd struct {
	Ctx *cmd.AppContext
}

func (h *InstallCmd) Run(args []string) error {
	h.Ctx.Config.Set("jar", "teste")
	h.Ctx.Config.WriteConfig()
	return nil
}

func init() {
	cmd.Register("install", func(ctx *cmd.AppContext) cmd.Command {
		return &InstallCmd{Ctx: ctx}
	})
}
