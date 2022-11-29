package cmd

import (
	"gctl/cmd/run"

	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run a single service locally",
	RunE:  run.Run,
}

var runCoreCmd = &cobra.Command{
	Use:   "core",
	Short: "Run all core galactus services locally",
	RunE:  run.Core,
}

func init() {
	// add parent
	rootCmd.AddCommand(runCmd)
	// add children
	runCmd.AddCommand(runCoreCmd)
	runCmd.Flags().StringVarP(&run.Domain, "domain", "d", "core", "domain for service")
	runCmd.Flags().StringVarP(&run.Service, "service", "s", "registry", "service to run")
	runCmd.Flags().StringVarP(&run.Version, "version", "v", "v1", "version of service to run")
}
