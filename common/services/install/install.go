package install

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var InstallCmd = &cobra.Command{
	Use:   "install",
	Short: "Installs a specific version of the FHIR Guard application",
	Long:  `Example: fg install 2.1.3`,
	Run:   install,
}

func install(cmd *cobra.Command, args []string) {
	appConfig, err := ParseYAMLFromURL("https://raw.githubusercontent.com/lem3s/fg-example-app/refs/heads/main/setup.yaml")
	if err != nil {
		log.Fatalf("Failed to parse YAML: %v", err)
	}

	fmt.Println("Downloading targz file...")
	err = downloadFile(appConfig)
	if err != nil {
		log.Fatal("Failed to download .targz: " + err.Error())
	}
	fmt.Println("Downloaded")
}
