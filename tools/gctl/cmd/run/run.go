package run

import (
	"fmt"
	"gctl/output"
	"os/exec"
	"path/filepath"

	e "gctl/exec"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	Domain  string
	Service string
	Version string
)

func Run(cmd *cobra.Command, args []string) error {

	output.Println("running service:")
	output.Println("- domain: %s", Domain)
	output.Println("- service: %s", Service)
	output.Println("- version: %s", Version)

	ctx := cmd.Context()

	// build out execution path
	servicePath := fmt.Sprintf("services/%s/%s/%s", Domain, Service, Version)
	rootPath := viper.GetString("config.root")
	fullPath := filepath.Join(rootPath, servicePath)

	// configure command
	c := exec.Command("go", "run", "main.go")
	c.Dir = fullPath

	// execute command
	err := e.ExecuteCommand(ctx, Service, output.Blue, c)
	if err != nil {
		output.Error(err)
	}

	return err
}
