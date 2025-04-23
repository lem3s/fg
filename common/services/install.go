package services

import (
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Installs a specific version of the FHIR Guard application",
	Long:  `Example: fg install 2.1.3`,
	Run:   install,
}

func install(cmd *cobra.Command, args []string) {
}
