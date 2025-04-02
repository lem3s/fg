package services

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"time"
)

type VersionInfo struct {
	Version     string    `json:"version"`
	InstallDate time.Time `json:"install_date"`
	IsActive    bool      `json:"is_active"`
	Components  []string  `json:"components,omitempty"`
}

type VersionSystem struct {
	AppName           string        `json:"app_name"`
	InstalledVersions []VersionInfo `json:"installed_versions"`
	LatestVersion     string        `json:"latest_version"`
	UpdateAvailable   bool          `json:"update_available"`
}

func ListVersions() {
	versionSystem := getMockVersionData()

	displayVersionInfo(versionSystem)
}

func ListVersionsFromFile(filePath string) error {
	_, err := os.Stat(filePath)
	if err != nil {
		return fmt.Errorf("arquivo não encontrado: %s", filePath)
	}
	
	data, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("erro ao ler o arquivo: %v", err)
	}
	
	var versionSystem VersionSystem
	err = json.Unmarshal(data, &versionSystem)
	if err != nil {
		return fmt.Errorf("erro ao processar o JSON: %v", err)
	}
	
	displayVersionInfo(versionSystem)
	return nil
}

func GetMockVersionsJSON() string {
	versionSystem := getMockVersionData()
	jsonBytes, err := json.MarshalIndent(versionSystem, "", "  ")
	if err != nil {
		return fmt.Sprintf("Erro ao gerar JSON: %v", err)
	}
	return string(jsonBytes)
}

func getMockVersionData() VersionSystem {
	now := time.Now()
	
	return VersionSystem{
		AppName: "MeuSistema",
		InstalledVersions: []VersionInfo{
			{
				Version:     "1.0.0",
				InstallDate: now.AddDate(0, -6, 0),
				IsActive:    false,
				Components:  []string{"core", "ui", "db"},
			},
			{
				Version:     "1.1.0",
				InstallDate: now.AddDate(0, -4, 0),
				IsActive:    false,
				Components:  []string{"core", "ui", "db", "api"},
			},
			{
				Version:     "1.2.0",
				InstallDate: now.AddDate(0, -2, 0),
				IsActive:    false,
				Components:  []string{"core", "ui", "db", "api", "auth"},
			},
			{
				Version:     "2.0.0",
				InstallDate: now.AddDate(0, 0, -15),
				IsActive:    true,
				Components:  []string{"core-v2", "ui-v2", "db-v2", "api-v2", "auth-v2", "analytics"},
			},
		},
		LatestVersion:   "2.1.0",
		UpdateAvailable: true,
	}
}


func displayVersionInfo(vs VersionSystem) {
	fmt.Printf("=== Informações de Versão para %s ===\n\n", vs.AppName)
	
	fmt.Printf("Versão atual: ")
	for _, v := range vs.InstalledVersions {
		if v.IsActive {
			fmt.Printf("%s (instalada em %s)\n", v.Version, v.InstallDate.Format("02/01/2006"))
			break
		}
	}
	
	fmt.Printf("Última versão disponível: %s\n", vs.LatestVersion)
	
	if vs.UpdateAvailable {
		fmt.Println("Status: Atualização disponível!")
	} else {
		fmt.Println("Status: Sistema atualizado")
	}
	
	fmt.Println("\nHistórico de versões instaladas:")
	
	sort.Slice(vs.InstalledVersions, func(i, j int) bool {
		return vs.InstalledVersions[i].InstallDate.After(vs.InstalledVersions[j].InstallDate)
	})
	
	for i, v := range vs.InstalledVersions {
		status := "inativa"
		if v.IsActive {
			status = "ATIVA"
		}
		
		fmt.Printf("%d. Versão %s - %s (Status: %s)\n", 
			i+1, v.Version, v.InstallDate.Format("02/01/2006"), status)
		
		if len(v.Components) > 0 {
			fmt.Printf("   Componentes: %v\n", v.Components)
		}
	}
}