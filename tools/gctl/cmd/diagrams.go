package cmd

import (
	"gctl/cmd/diagrams"

	"github.com/spf13/cobra"
)

var diagramsCmd = &cobra.Command{
	Use:   "diagrams",
	Short: "Manage generated diagram (PlantUML) files within `docs/diagrams`",
}

var diagramsBuildCmd = &cobra.Command{
	Use:   "build",
	Short: "Compile all diagram source files into visual documents",
	Long: `Compile all diagram source files into visual documents.
Defaults to PNGs but can take any file type from here: https://plantuml.com/command-line
Make sure to only include the type and not include the '-t'. For example, to create
PDFs use the flag '-f=pdf' NOT '-f=-tpdf'`,
	RunE:  diagrams.Build,
}

var diagramsCleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Delete all generated PlantUML diagram files",
	RunE:  diagrams.Clean,
}

var diagramsInitCmd = &cobra.Command{
	Use:   "init",
	Short: "Build the PlantUML builder Docker image",
	RunE:  diagrams.Init,
}

func init() {
	// add parent
	rootCmd.AddCommand(diagramsCmd)
	// add children
	diagramsCmd.AddCommand(diagramsBuildCmd)
	diagramsBuildCmd.Flags().StringVarP(&diagrams.Type, "type", "t", "png", "file type to generate (png, pdf, etc.)")
	diagramsCmd.AddCommand(diagramsCleanCmd)
	diagramsCmd.AddCommand(diagramsInitCmd)
}
