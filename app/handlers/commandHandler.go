package handlers

import "github.com/lem3s/fg/app/cmd"

func HandleCallback(err error, ctx *cmd.AppContext) {
	if err != nil {
		ctx.Interactor.Error(err.Error())
		return
	}

	ctx.Interactor.Info(ctx.SuccessMessage)
}

func HandleParams(params []string, ctx *cmd.AppContext) {
	//logica para dar handle nos parametros globais, --dir e afins
	//parametros locais serão validados dentro do comando
}
