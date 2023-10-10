package infra

import (
	"gctl/docker"
	"gctl/output"

	"github.com/spf13/cobra"
)

func Stop(cmd *cobra.Command, args []string) (err error) {
	ctx := cmd.Context()
	dctl, err := docker.NewDockerController()
	if err != nil {
		return nil
	}

	// stop hasura before any databases
	err = dctl.StopContainerByName(ctx, hasuraContainer)
	if err != nil {
		output.Error(err)
	}
	err = dctl.StopContainerByName(ctx, postgresContainer)
	if err != nil {
		output.Error(err)
	}
	err = dctl.StopContainerByName(ctx, mongoContainer)
	if err != nil {
		output.Error(err)
	}
	err = dctl.StopContainerByName(ctx, rabbitContainer)
	if err != nil {
		output.Error(err)
	}

	return nil
}