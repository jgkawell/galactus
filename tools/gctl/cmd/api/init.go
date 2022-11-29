package api

import (
	"path/filepath"

	"gctl/docker"
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

	output.Println("Finished")
	return nil
}
