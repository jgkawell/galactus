package cmd

import (
	"gctl/cmd/api"

	"github.com/spf13/cobra"
)

// apiCmd represents the api command
var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "The parent to all api commands",
}

// apiBuildCmd represents the build command
var apiBuildCmd = &cobra.Command{
	Use:   "build",
	Short: "A brief description of your command",
	RunE:  api.Build,
}

// apiCleanCmd represents the clean command
var apiCleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: api.Clean,
}

// apiInitCmd represents the init command
var apiInitCmd = &cobra.Command{
	Use:   "init",
	Short: "A brief description of your command",
	RunE:  api.Init,
}

func init() {
	// add parent
	rootCmd.AddCommand(apiCmd)
	// add children
	apiCmd.AddCommand(apiBuildCmd)
	apiCmd.AddCommand(apiCleanCmd)
	apiCmd.AddCommand(apiInitCmd)
}
