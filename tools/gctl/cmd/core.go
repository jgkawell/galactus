package cmd

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"path/filepath"

	e "gctl/exec"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// coreCmd represents the core command
var coreCmd = &cobra.Command{
	Use:   "core",
	Short: "Run all core galactus services",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := cmd.Context()

		// start registry first
		go func() {
			err := runService(ctx, "core", "registry", "v1")
			if err != nil {
				log.Fatal(err)
			}
		}()

		// block until registry is up and healthy
		healthy := false
		for !healthy {
			resp, _ := http.Get("http://localhost:35000/health")
			if resp != nil && resp.StatusCode == http.StatusOK {
				healthy = true
			}
		}

		// start commandhandler
		go func() {
			err := runService(ctx, "core", "commandhandler", "v1")
			if err != nil {
				log.Fatal(err)
			}
		}()

		// start eventstore
		go func() {
			err := runService(ctx, "core", "eventstore", "v1")
			if err != nil {
				log.Fatal(err)
			}
		}()

		// start notifier
		go func() {
			err := runService(ctx, "core", "notifier", "v1")
			if err != nil {
				log.Fatal(err)
			}
		}()

		// wait until user interrupt (ctrl + c)
		<-ctx.Done()
	},
}

func init() {
	runCmd.AddCommand(coreCmd)
}

func runService(ctx context.Context, domain, service, version string) error {

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

	return nil
}
