package cmd

import (
	"gctl/cmd/infra"

	"github.com/spf13/cobra"
)

var infraCmd = &cobra.Command{
	Use:   "infra",
	Short: "Manage all local galactus infra (Docker containers)",
}

var infraCleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Clean up infra resources (Docker containers)",
	RunE:  infra.Clean,
}

var infraInitCmd = &cobra.Command{
	Use:   "init",
	Short: "Build all custom galactus infra Docker images",
	RunE:  infra.Init,
}

var infraStartCmd = &cobra.Command{
	Use:   "start",
	Short: "Start all galactus infra locally",
	RunE:  infra.Start,
}

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
