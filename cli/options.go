package cli

import (
	"fmt"
	"module/common"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "fg",
	Short: "base of cli",
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
