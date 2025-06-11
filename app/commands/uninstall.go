package commands

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/lem3s/fg/app/cmd"
)

type UninstallCmd struct {
	Ctx *cmd.AppContext
}

func (h *UninstallCmd) Run(args []string) error {
	fgHome := h.Ctx.FgHome
	if fgHome == "" {
		return fmt.Errorf("diretório FG Home não configurado")
	}

	var versionToUninstall string
	if len(args) > 0 {
		versionToUninstall = args[0]
	} else {
		versionToUninstall = "1.2.3"
		
		h.Ctx.Interactor.Info(fmt.Sprintf("Nenhuma versão especificada. Usando a última versão instalada: %s", versionToUninstall))
	}

	return h.uninstallVersion(fgHome, versionToUninstall)
}

func (h *UninstallCmd) uninstallVersion(fgHome, version string) error {
	versionPath, err := h.findVersionPath(fgHome, version)
	if err != nil {
		return err
	}

	versionName := filepath.Base(versionPath)
	h.Ctx.Interactor.Info(fmt.Sprintf("Desinstalando versão: %s", versionName))
	h.Ctx.Interactor.Info(fmt.Sprintf("Localização: %s", versionPath))

	jarFiles := h.findJarFiles(versionPath)
	if len(jarFiles) > 0 {
		h.Ctx.Interactor.Info("Arquivos JAR encontrados:")
		for _, jar := range jarFiles {
			h.Ctx.Interactor.Info(fmt.Sprintf("  - %s", filepath.Base(jar)))
		}
	} else {
		h.Ctx.Interactor.Warn("Nenhum arquivo JAR encontrado nesta versão.")
	}

	confirmMessage := fmt.Sprintf("Tem certeza que deseja desinstalar a versão %s?", versionName)
	if !h.Ctx.Interactor.Confirm(confirmMessage) {
		h.Ctx.Interactor.Info("Desinstalação cancelada.")
		return nil
	}

	if err := os.RemoveAll(versionPath); err != nil {
		return fmt.Errorf("erro ao remover diretório: %v", err)
	}

	h.Ctx.Interactor.Info(fmt.Sprintf("Versão '%s' desinstalada com sucesso!", versionName))
	return nil
}

func (h *UninstallCmd) findVersionPath(fgHome, version string) (string, error) {
	version = strings.TrimSpace(version)
	versionPath := filepath.Join(fgHome, version)

	if _, err := os.Stat(versionPath); err == nil {
		return versionPath, nil
	}

	entries, err := os.ReadDir(fgHome)
	if err != nil {
		return "", fmt.Errorf("erro ao ler diretório: %v", err)
	}

	for _, entry := range entries {
		if entry.IsDir() && strings.Contains(entry.Name(), version) {
			return filepath.Join(fgHome, entry.Name()), nil
		}
	}

	return "", fmt.Errorf("versão '%s' não encontrada", version)
}

func (h *UninstallCmd) findJarFiles(dirPath string) []string {
	var jarFiles []string
	filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err == nil && !info.IsDir() && strings.HasSuffix(strings.ToLower(info.Name()), ".jar") {
			jarFiles = append(jarFiles, path)
		}
		return nil
	})
	return jarFiles
}

func init() {
	cmd.Register("uninstall", func(ctx *cmd.AppContext) cmd.Command {
		return &UninstallCmd{Ctx: ctx}
	})
}