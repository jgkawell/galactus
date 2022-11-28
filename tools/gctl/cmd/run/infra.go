package run

import (
	"context"
	"gctl/docker"
	"gctl/output"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/go-connections/nat"
	"github.com/spf13/cobra"
)

const (
	mongoImage    = "mongo:4.2.21"
	postgresImage = "postgres:14.4"
	brokerImage   = "broker"
	hasuraImage   = "hasura/graphql-engine:v2.15.2"
)

func Infra(cmd *cobra.Command, args []string) (err error) {
	ctx := cmd.Context()

	// postgres
	config := &container.Config{
		Image: postgresImage,
		Env: []string{
			"POSTGRES_PASSWORD=admin",
			"POSTGRES_USER=admin",
			"POSTGRES_DB=dev",
		},
	}
	hostConfig := &container.HostConfig{
		PortBindings: nat.PortMap{
			"5432/tcp": []nat.PortBinding{
				{HostPort: "5434"},
			},
		},
	}
	id, err := docker.StartContainer(ctx, "galactus-sql-db", config, hostConfig, true, true)
	if err != nil {
		output.Error(err)
		return err
	}

	<-ctx.Done()

	ctx = context.Background()
	err = docker.StopContainer(ctx, id)
	if err != nil {
		return err
	}
	err = docker.RemoveContainer(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
