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

// buildCmd represents the build command
var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "A brief description of your command",
	RunE:  api.Build,
}

// cleanCmd represents the clean command
var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: api.Clean,
}

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "A brief description of your command",
	RunE:  api.Init,
}

func init() {
	// add parent
	rootCmd.AddCommand(apiCmd)
	// add children
	apiCmd.AddCommand(buildCmd)
	apiCmd.AddCommand(cleanCmd)
	apiCmd.AddCommand(initCmd)
}
