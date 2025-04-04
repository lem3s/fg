package cli

import (
	"fmt"
	"module/common"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "fg",
	Short: "A tool for managing and running the FHIR Guard application",
	Long: `This cli application provides a consistent and easy-to-use interface
	 for installing, updating, starting, stopping, and monitoring different versions
	 of the FHIR Guard application. By default, fg operates in CLI mode, but a 
	 graphical interface can be launched by using the gui command.`,
}

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Display configuration", //fazer i18n
	Run: func(cmd *cobra.Command, args []string) {
		resultChan := make(chan string)
		errChan := make(chan error)

		go func() {
			result, err := common.ReadAndParseConfig()
			if err != nil {
				errChan <- err
				return
			}
			resultChan <- result
		}()

		select {
		case result := <-resultChan:
			fmt.Println(result)
		case err := <-errChan:
			fmt.Printf("%v\n", err)
		}
	},
}

func Execute() error {
	rootCmd.AddCommand(configCmd)
	return rootCmd.Execute()
}
