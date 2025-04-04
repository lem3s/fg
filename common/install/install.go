package install

import (
	"github.com/lem3s/fg/cli"

	"github.com/spf13/cobra"
)

func init() {
	cli.RootCmd.AddCommand(installCmd)
}

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Installs a specific version of the FHIR Guard application",
	Long:  `Example: fg install 2.1.3`,
	Run: install,
}

func install(cmd *cobra.Command, args []string) {
}