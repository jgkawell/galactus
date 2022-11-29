package cmd

import (
	"gctl/cmd/infra"

	"github.com/spf13/cobra"
)

// infraCmd represents the run command
var infraCmd = &cobra.Command{
	Use:   "infra",
	Short: "Parent",
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
	Short: "Stop all galactus infra",
	RunE:  infra.Stop,
}

// infraInitCmd represents the infra command
var infraInitCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize infra (build Docker images, etc.)",
	RunE:  infra.Init,
}

// infraCleanCmd represents the infra command
var infraCleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Clean up infra resources (Docker images, containers, etc.)",
	RunE:  infra.Clean,
}

func init() {
	// root
	rootCmd.AddCommand(infraCmd)
	// run
	infraCmd.AddCommand(infraStartCmd)
	infraStartCmd.Flags().BoolVarP(&infra.Follow, "follow", "f", false, "whether or not to follow the output of the infra docker containers")
	infraCmd.AddCommand(infraStopCmd)
	infraCmd.AddCommand(infraInitCmd)
	infraCmd.AddCommand(infraCleanCmd)
}
