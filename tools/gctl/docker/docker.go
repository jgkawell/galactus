package docker

import (
	"context"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
)

func RunContainer(ctx context.Context, containerName string, config *container.Config, host *container.HostConfig, showOutput, removeContainer bool) error {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return err
	}

	// configure and create container
	if showOutput {
		config.AttachStdout = true
		config.AttachStderr = true
	}
	resp, err := cli.ContainerCreate(ctx, config, host, nil, nil, containerName)
	if err != nil {
		return err
	}

	// start container
	err = cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{})
	if err != nil {
		return err
	}

	// wait for container to complete
	_, err = waitForContainer(ctx, cli, resp.ID)
	if err != nil {
		return err
	}

	if showOutput {
		// retrieve and print logs
		err := printContainerLogs(ctx, cli, containerName)
		if err != nil {
			return err
		}
	}

	if removeContainer {
		// remove completed container
		err = cli.ContainerRemove(ctx, resp.ID, types.ContainerRemoveOptions{})
		if err != nil {
			return err
		}
	}
	return nil
}

func waitForContainer(ctx context.Context, cli *client.Client, id string) (state int64, err error) {
	resultC, errC := cli.ContainerWait(ctx, id, "")
	select {
	case err := <-errC:
		return 0, err
	case result := <-resultC:
		return result.StatusCode, nil
	}
}

func printContainerLogs(ctx context.Context, cli *client.Client, containerName string) error {
	reader, err := cli.ContainerLogs(ctx, containerName, types.ContainerLogsOptions{
		ShowStdout: true,
		ShowStderr: true,
	})
	if err != nil {
		return err
	}
	defer reader.Close()
	_, err = stdcopy.StdCopy(os.Stdout, os.Stderr, reader)
	if err != nil {
		return err
	}
	return nil
}
