package run

import (
	"context"
	"time"

	"gctl/docker"
	"gctl/output"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/go-connections/nat"
	"github.com/spf13/cobra"
)

const (
	mongoImage    = "mongo:4.2.21"
	postgresImage = "postgres:14.4"
	rabbitImage   = "rabbitmq:galactus"
	hasuraImage   = "hasura/graphql-engine:v2.15.2"
)

func Infra(cmd *cobra.Command, args []string) (err error) {
	ctx := cmd.Context()

	var config *container.Config
	var hostConfig *container.HostConfig

	// postgres
	config = &container.Config{
		Image: postgresImage,
		Env: []string{
			"POSTGRES_PASSWORD=admin",
			"POSTGRES_USER=admin",
			"POSTGRES_DB=dev",
		},
		ExposedPorts: map[nat.Port]struct{}{
			"5432/tcp": {}},
	}
	hostConfig = &container.HostConfig{
		PortBindings: nat.PortMap{
			"5432/tcp": []nat.PortBinding{
				{
					HostIP:   "0.0.0.0",
					HostPort: "5434",
				},
			},
		},
	}
	id, err := docker.StartContainer(ctx, "galactus-postgres", config, hostConfig, true, true)
	if err != nil {
		output.Error(err)
		return err
	}
	defer stop(context.Background(), id)

	// mongo
	config = &container.Config{
		Image: mongoImage,
		Env:   []string{},
		ExposedPorts: map[nat.Port]struct{}{
			"27017/tcp": {}},
	}
	hostConfig = &container.HostConfig{
		PortBindings: nat.PortMap{
			"27017/tcp": []nat.PortBinding{
				{
					HostIP:   "0.0.0.0",
					HostPort: "27017",
				},
			},
		},
	}
	id, err = docker.StartContainer(ctx, "galactus-mongo", config, hostConfig, true, true)
	if err != nil {
		output.Error(err)
		return err
	}
	defer stop(context.Background(), id)

	// rabbitmq
	config = &container.Config{
		Image: rabbitImage,
		Env:   []string{},
		ExposedPorts: map[nat.Port]struct{}{
			"15672/tcp": {},
			"5672/tcp":  {},
		},
	}
	hostConfig = &container.HostConfig{
		PortBindings: nat.PortMap{
			"15672/tcp": []nat.PortBinding{
				{
					HostIP:   "0.0.0.0",
					HostPort: "15672",
				},
			},
			"5672/tcp": []nat.PortBinding{
				{
					HostIP:   "0.0.0.0",
					HostPort: "5672",
				},
			},
		},
	}
	id, err = docker.StartContainer(ctx, "galactus-rabbitmq", config, hostConfig, true, true)
	if err != nil {
		output.Error(err)
		return err
	}
	defer stop(context.Background(), id)

	time.Sleep(5 * time.Second)

	// hasura
	config = &container.Config{
		Image: hasuraImage,
		Env: []string{
			"HASURA_GRAPHQL_METADATA_DATABASE_URL=postgres://admin:admin@host.docker.internal:5434/dev",
			"HASURA_GRAPHQL_DATABASE_URL=postgres://admin:admin@host.docker.internal:5434/dev",
			"HASURA_GRAPHQL_ENABLE_CONSOLE=true",
			"HASURA_GRAPHQL_ENABLED_LOG_TYPES=startup, http-log, webhook-log, websocket-log, query-log",
			"HASURA_GRAPHQL_ENABLE_TELEMETRY=false",
		},
		ExposedPorts: map[nat.Port]struct{}{
			"8080/tcp": {},
		},
	}
	hostConfig = &container.HostConfig{
		PortBindings: nat.PortMap{
			"8080/tcp": []nat.PortBinding{
				{
					HostIP:   "0.0.0.0",
					HostPort: "8082",
				},
			},
		},
	}
	id, err = docker.StartContainer(ctx, "galactus-hasura", config, hostConfig, true, true)
	if err != nil {
		output.Error(err)
		return err
	}
	defer stop(context.Background(), id)

	<-ctx.Done()

	return nil
}

func stop(ctx context.Context, id string) {
	err := docker.StopContainer(ctx, id)
	if err != nil {
		output.Error(err)
	}
}
