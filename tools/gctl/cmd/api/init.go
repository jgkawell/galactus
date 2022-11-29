package api

import (
	"path/filepath"

	"gctl/docker"
	"gctl/output"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// TODO: have a command to remove generated Docker images?

func Init(cmd *cobra.Command, args []string) error {
	output.Println("init called")

	ctx := cmd.Context()

	// build out execution path
	rootPath := viper.GetString("config.root")
	fullPath := filepath.Join(rootPath, "api")

	err := docker.BuildImage(ctx, fullPath, "proto-builder:v3")
	if err != nil {
		output.Error(err)
		return err
	}

	output.Println("Finished")
	return nil
}
