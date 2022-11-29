package cmd

import (
	"gctl/cmd/infra"

	"github.com/spf13/cobra"
)

// infraCmd represents the run command
var infraCmd = &cobra.Command{
	Use:   "infra",
	Short: "Manage all local galactus infra (Docker containers)",
}

// infraCleanCmd represents the infra command
var infraCleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Clean up infra resources (Docker containers)",
	RunE:  infra.Clean,
}

// infraInitCmd represents the infra command
var infraInitCmd = &cobra.Command{
	Use:   "init",
	Short: "Build all custom galactus infra Docker images",
	RunE:  infra.Init,
}

// infraStartCmd represents the infra command
var infraStartCmd = &cobra.Command{
	Use:   "start",
	Short: "Start all galactus infra locally",
	RunE:  infra.Start,
}

// infraStopCmd represents the infra command
var infraStopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop all running galactus infra",
	RunE:  infra.Stop,
}

func init() {
	// add parent
	rootCmd.AddCommand(infraCmd)
	// add children
	infraCmd.AddCommand(infraCleanCmd)
	infraCmd.AddCommand(infraInitCmd)
	infraCmd.AddCommand(infraStartCmd)
	infraStartCmd.Flags().BoolVarP(&infra.Follow, "follow", "f", false, "whether or not to follow the output of the infra docker containers (true/false)")
	infraCmd.AddCommand(infraStopCmd)
}
