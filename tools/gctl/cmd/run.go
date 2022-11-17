package cmd

import (
	"fmt"
	"log"
	"os/exec"
	"path/filepath"

	e "gctl/exec"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	domain  string
	service string
	version string
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run a service locally",
	Run: func(cmd *cobra.Command, args []string) {

		log.Println("running service:")
		log.Printf("- domain: %s\n", domain)
		log.Printf("- service: %s\n", service)
		log.Printf("- version: %s\n", version)

		ctx := cmd.Context()

		// build out execution path
		servicePath := fmt.Sprintf("services/%s/%s/%s", domain, service, version)
		rootPath := viper.GetString("config.root")
		fullPath := filepath.Join(rootPath, servicePath)

		// configure command
		c := exec.Command("go", "run", "main.go")
		c.Dir = fullPath

		// execute command
		err := e.ExecuteCommand(ctx, c)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	runCmd.Flags().StringVarP(&domain, "domain", "d", "core", "domain for service")
	runCmd.Flags().StringVarP(&service, "service", "s", "registry", "service to run")
	runCmd.Flags().StringVarP(&version, "version", "v", "v1", "version of service to run")
}
