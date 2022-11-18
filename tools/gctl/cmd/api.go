package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// apiCmd represents the api command
var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "A brief description of your command",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("api called")
	},
}

func init() {
	rootCmd.AddCommand(apiCmd)

}
