package run

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os/exec"
	"path/filepath"
	"time"

	e "gctl/exec"
	"gctl/output"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	commandHandlerVersion = "v1"
	eventStoreVersion     = "v1"
	notifierVersion       = "v1"
	registryVersion       = "v1"

	registryHealthEndpoint               = "http://localhost:35000/health"
	registryHealthCheckRetries           = 3
	registryHealthCheckRetryDelaySeconds = 3
)

func Core(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()

	// start registry first
	go runService(ctx, "core", "registry", registryVersion, output.Blue)

	// block until registry is up and healthy
	healthy := false
	retries := 0
	for !healthy && retries < registryHealthCheckRetries {
		resp, _ := http.Get(registryHealthEndpoint)
		if resp != nil && resp.StatusCode == http.StatusOK {
			healthy = true
		}
		time.Sleep(registryHealthCheckRetryDelaySeconds * time.Second)
		retries++
		output.Println("retrying registry health check")
	}
	if !healthy {
		output.Println("registry service never became healthy")
		return errors.New("registry failed to start")
	}

	// start commandhandler
	go runService(ctx, "core", "commandhandler", commandHandlerVersion, output.Cyan)

	// start eventstore
	go runService(ctx, "core", "eventstore", eventStoreVersion, output.Green)

	// start notifier
	go runService(ctx, "core", "notifier", notifierVersion, output.Yellow)

	// wait until user interrupt (ctrl + c)
	<-ctx.Done()
	return nil
}

func runService(ctx context.Context, domain, service, version string, c output.Color) {
	// build out execution path
	servicePath := fmt.Sprintf("services/%s/%s/%s", domain, service, version)
	rootPath := viper.GetString("config.root")
	fullPath := filepath.Join(rootPath, servicePath)

	// configure command
	cmd := exec.Command("go", "run", "main.go")
	cmd.Dir = fullPath

	// execute command
	output.Println("running %s service...", service)
	err := e.ExecuteCommand(ctx, service, c, cmd)
	if err != nil {
		output.Println("%s service failed to run with error: %s", service, err.Error())
	}
}
