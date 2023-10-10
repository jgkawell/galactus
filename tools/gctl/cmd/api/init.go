package api

import (
	"errors"
	"os"
	"os/exec"
	"path/filepath"

	"gctl/docker"
	e "gctl/exec"
	"gctl/output"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	protoImage = "proto:galactus"
)

func Init(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()
	dctl, err := docker.NewDockerController()
	if err != nil {
		return nil
	}

	// build out execution path
	rootPath := viper.GetString("config.root")
	fullPath := filepath.Join(rootPath, "api")

	err = dctl.BuildImage(ctx, fullPath, protoImage)
	if err != nil {
		output.Error(err)
		return err
	}

	output.Println("Initializing generated code directories...")
	// make go directory
	err = os.MkdirAll(filepath.Join(fullPath, "gen", "go"), os.ModePerm)
	if err != nil {
		output.Error(err)
		return err
	}
	// initialize go mod only if go.mod doesn't already exist
	_, err = os.Stat(filepath.Join(fullPath, "gen", "go", "go.mod"))
	if errors.Is(err, os.ErrNotExist) {
		c := exec.Command("go", "mod", "init", "github.com/jgkawell/galactus/api/gen/go")
		c.Dir = filepath.Join(fullPath, "gen", "go")
		err = e.ExecuteCommand(ctx, "go", output.Magenta, c)
		if err != nil {
			output.Error(err)
			return err
		}
	} else if err != nil {
		// if there's an error other than the file not existing, return it
		output.Error(err)
		return err
	}

	output.Println("Finished")
	return nil
}
