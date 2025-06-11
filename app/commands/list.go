package commands

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/lem3s/fg/app/cmd"
)

type VersionInfo struct {
	Version     string    `yaml:"version"`
	InstallDate time.Time `yaml:"install_date"`
	IsActive    bool      `yaml:"is_active"`
}

type ListCmd struct {
	Ctx *cmd.AppContext
}

func (h *ListCmd) Run(args []string) error {
	fgHome := h.Ctx.FgHome
	if fgHome == "" {
		return fmt.Errorf("diretório FG Home não configurado")
	}

	err := h.listVersionsFromFgHome(fgHome)
	if err != nil {
		return err
	}

	return nil
}

func (h *ListCmd) listVersionsFromFgHome(fgHome string) error {
	if _, err := os.Stat(fgHome); os.IsNotExist(err) {
		h.Ctx.Interactor.Info("Nenhuma versão instalada. Diretório principal não existe.")
		return nil
	}

	entries, err := os.ReadDir(fgHome)
	if err != nil {
		return fmt.Errorf("erro ao ler diretório principal: %v", err)
	}

	var versions []VersionInfo

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		versionName := entry.Name()
		if strings.HasPrefix(versionName, ".") {
			continue
		}

		versionPath := filepath.Join(fgHome, versionName)
		versionInfo := h.createVersionInfo(versionPath, versionName)
		versions = append(versions, versionInfo)
	}

	if len(versions) == 0 {
		h.Ctx.Interactor.Info("Nenhuma versão instalada.")
		return nil
	}

	h.setActiveVersion(versions)
	h.displayVersionInfo(versions)
	return nil
}

func (h *ListCmd) createVersionInfo(versionPath, versionName string) VersionInfo {
	dirInfo, err := os.Stat(versionPath)
	installDate := time.Now()
	if err == nil {
		installDate = dirInfo.ModTime()
	}

	return VersionInfo{
		Version:     versionName,
		InstallDate: installDate,
		IsActive:    false,
	}
}

func (h *ListCmd) setActiveVersion(versions []VersionInfo) {
	if len(versions) == 0 {
		return
	}

	sort.Slice(versions, func(i, j int) bool {
		return versions[i].InstallDate.After(versions[j].InstallDate)
	})

	for i := range versions {
		versions[i].IsActive = false
	}

	versions[0].IsActive = true
}

func (h *ListCmd) displayVersionInfo(versions []VersionInfo) {
	h.Ctx.Interactor.Info("=== Informações das Versões Instaladas ===")

	sort.Slice(versions, func(i, j int) bool {
		return versions[i].InstallDate.After(versions[j].InstallDate)
	})

	versionAtualExibida := false

	h.Ctx.Interactor.Info("Versões instaladas:")

	for i, v := range versions {
		status := "inativa"
		if v.IsActive {
			status = "ATIVA"

			if !versionAtualExibida {
				h.Ctx.Interactor.Info(fmt.Sprintf("Versão atual: %s (instalada em %s)", v.Version, v.InstallDate.Format("02/01/2006")))
				versionAtualExibida = true
			}
		}

		h.Ctx.Interactor.Info(fmt.Sprintf("%d. Versão %s - %s (Status: %s)", i+1, v.Version, v.InstallDate.Format("02/01/2006"), status))
	}

	if !versionAtualExibida {
		h.Ctx.Interactor.Info("Versão atual: Nenhuma versão ativa encontrada")
	}
}

func init() {
	cmd.Register("list", func(ctx *cmd.AppContext) cmd.Command {
		return &ListCmd{Ctx: ctx}
	})
}