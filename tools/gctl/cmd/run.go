package cmd

import (
	"gctl/cmd/run"

	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run a single service locally",
	RunE:  run.Run,
}

// runCoreCmd represents the core command
var runCoreCmd = &cobra.Command{
	Use:   "core",
	Short: "Run all core galactus services locally",
	RunE:  run.Core,
}

// runInfraCmd represents the infra command
var runInfraCmd = &cobra.Command{
	Use:   "infra",
	Short: "Run all galactus infra locally",
	RunE:  run.Infra,
}

// runInitCmd represents the infra command
var runInitCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize local machine for running infra and services",
	RunE:  run.Init,
}

func init() {
	// add parent
	rootCmd.AddCommand(runCmd)
	// add children
	runCmd.AddCommand(runCoreCmd)
	runCmd.Flags().StringVarP(&run.Domain, "domain", "d", "core", "domain for service")
	runCmd.Flags().StringVarP(&run.Service, "service", "s", "registry", "service to run")
	runCmd.Flags().StringVarP(&run.Version, "version", "v", "v1", "version of service to run")
	runCmd.AddCommand(runInfraCmd)
	runCmd.AddCommand(runInitCmd)
}
