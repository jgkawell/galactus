package cmd

import (
	"gctl/cmd/api"

	"github.com/spf13/cobra"
)

// apiCmd represents the api command
var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "Commands for managing the galactus API (protobufs)",
}

// apiBuildCmd represents the build command
var apiBuildCmd = &cobra.Command{
	Use:   "build",
	Short: "Compile all protobuf files",
	RunE:  api.Build,
}

// apiCleanCmd represents the clean command
var apiCleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Clean up all compiled (generated) protobuf files",
	RunE: api.Clean,
}

// apiInitCmd represents the init command
var apiInitCmd = &cobra.Command{
	Use:   "init",
	Short: "Build the Docker image that compiles protobufs",
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
