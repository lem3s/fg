package commands

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/lem3s/fg/app/cmd"
	"gopkg.in/yaml.v3"
)

type VersionInfo struct {
    Version     string    `yaml:"version"`
    InstallDate time.Time `yaml:"install_date"`
    IsActive    bool      `yaml:"is_active"`
    Components  []string  `yaml:"components,omitempty"`
}

type VersionSystem struct {
    AppName           string        `yaml:"app_name"`
    InstalledVersions []VersionInfo `yaml:"installed_versions"`
    LatestVersion     string        `yaml:"latest_version"`
    UpdateAvailable   bool          `yaml:"update_available"`
}

type VersionMetadata struct {
	Version     string    `yaml:"version"`
	InstallDate time.Time `yaml:"install_date"`
	IsActive    bool      `yaml:"is_active"`
	Components  []string  `yaml:"components,omitempty"`
	JarFiles    []string  `yaml:"jar_files,omitempty"`
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
		return fmt.Errorf("erro ao listar versões: %v", err)
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
		versionInfo, err := h.readVersionMetadata(versionPath, versionName)
		if err != nil {
			h.Ctx.Interactor.Warn(fmt.Sprintf("Aviso: Não foi possível ler metadados para versão %s: %v", versionName, err))
			versionInfo = h.createBasicVersionInfo(versionPath, versionName)
		}

		versions = append(versions, versionInfo)
	}

	if len(versions) == 0 {
		h.Ctx.Interactor.Info("Nenhuma versão instalada.")
		return nil
	}

	versionSystem := VersionSystem{
		AppName:           "FG",
		InstalledVersions: versions,
		LatestVersion:     "desconhecida",
		UpdateAvailable:   false,
	}

	h.displayVersionInfo(versionSystem)
	return nil
}

func (h *ListCmd) readVersionMetadata(versionPath, versionName string) (VersionInfo, error) {

	metadataFiles := []string{".fg.yaml", "metadata.yaml", "config.yaml", "version.yaml"}

	for _, filename := range metadataFiles {
		metadataPath := filepath.Join(versionPath, filename)
		if _, err := os.Stat(metadataPath); err == nil {
			data, err := os.ReadFile(metadataPath)
			if err != nil {
				continue
			}

			var metadata VersionMetadata
			if err := yaml.Unmarshal(data, &metadata); err != nil {
				continue
			}

			return VersionInfo{
				Version:     metadata.Version,
				InstallDate: metadata.InstallDate,
				IsActive:    metadata.IsActive,
				Components:  metadata.Components,
			}, nil
		}
	}

	return VersionInfo{}, fmt.Errorf("nenhum arquivo de metadados encontrado")
}

func (h *ListCmd) createBasicVersionInfo(versionPath, versionName string) VersionInfo {
	dirInfo, err := os.Stat(versionPath)
	installDate := time.Now()
	if err == nil {
		installDate = dirInfo.ModTime()
	}

	jarCount := h.countJarFiles(versionPath)
	components := []string{}
	if jarCount > 0 {
		components = append(components, fmt.Sprintf("%d arquivo(s) JAR", jarCount))
	}

	return VersionInfo{
		Version:     versionName,
		InstallDate: installDate,
		IsActive:    false, 
		Components:  components,
	}
}

func (h *ListCmd) countJarFiles(dirPath string) int {
	count := 0
	filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil 
		}
		if !info.IsDir() && strings.HasSuffix(strings.ToLower(info.Name()), ".jar") {
			count++
		}
		return nil
	})
	return count
}

func (h *ListCmd) displayVersionInfo(vs VersionSystem) {
	h.Ctx.Interactor.Info(fmt.Sprintf("=== Informações de Versão para %s ===\n", vs.AppName))

	sort.Slice(vs.InstalledVersions, func(i, j int) bool {
		return vs.InstalledVersions[i].InstallDate.After(vs.InstalledVersions[j].InstallDate)
	})

	versionAtualExibida := false

	h.Ctx.Interactor.Info("Versões instaladas:")

	for i, v := range vs.InstalledVersions {
		status := "inativa"
		if v.IsActive {
			status = "ATIVA"

			if !versionAtualExibida {
				h.Ctx.Interactor.Info(fmt.Sprintf("\nVersão atual: %s (instalada em %s)",
					v.Version, v.InstallDate.Format("02/01/2006")))
				versionAtualExibida = true
			}
		}

		h.Ctx.Interactor.Info(fmt.Sprintf("%d. Versão %s - %s (Status: %s)",
			i+1, v.Version, v.InstallDate.Format("02/01/2006"), status))

		if len(v.Components) > 0 {
			h.Ctx.Interactor.Info(fmt.Sprintf("   Componentes: %v", v.Components))
		}
	}

	if !versionAtualExibida {
		h.Ctx.Interactor.Info("\nVersão atual: Nenhuma versão ativa encontrada")
	}

	if vs.LatestVersion != "desconhecida" {
		h.Ctx.Interactor.Info(fmt.Sprintf("\nÚltima versão disponível: %s", vs.LatestVersion))

		if vs.UpdateAvailable {
			h.Ctx.Interactor.Info("Status: Atualização disponível!")
		} else {
			h.Ctx.Interactor.Info("Status: Sistema atualizado")
		}
	}
}

func init() {
	cmd.Register("list", func(ctx *cmd.AppContext) cmd.Command {
		return &ListCmd{Ctx: ctx}
	})
}