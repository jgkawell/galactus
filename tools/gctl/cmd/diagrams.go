package cmd

import (
	"gctl/cmd/diagrams"

	"github.com/spf13/cobra"
)

// diagramsCmd represents the api command
var diagramsCmd = &cobra.Command{
	Use:   "diagrams",
	Short: "The parent to all diagram commands",
}

// diagramsBuildCmd represents the build command
var diagramsBuildCmd = &cobra.Command{
	Use:   "build",
	Short: "A brief description of your command",
	RunE:  diagrams.Build,
}

// diagramsInitCmd represents the init command
var diagramsInitCmd = &cobra.Command{
	Use:   "init",
	Short: "A brief description of your command",
	RunE:  diagrams.Init,
}

// diagramsCleanCmd represents the init command
var diagramsCleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "A brief description of your command",
	RunE:  diagrams.Clean,
}

func init() {
	// add parent
	rootCmd.AddCommand(diagramsCmd)
	// add children
	diagramsCmd.AddCommand(diagramsBuildCmd)
	diagramsCmd.AddCommand(diagramsInitCmd)
	diagramsCmd.AddCommand(diagramsCleanCmd)
}
