package docker

import (
	"bufio"
	"context"
	"strings"
	"time"

	"gctl/output"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
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
		config.Tty = true
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

	if showOutput {
		// retrieve and print logs
		err := printContainerLogs(ctx, cli, containerName)
		if err != nil {
			return err
		}
	}

	// wait for container to complete
	_, err = waitForContainer(ctx, cli, resp.ID)
	if err != nil {
		return err
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

func StartContainer(ctx context.Context, containerName string, config *container.Config, host *container.HostConfig, showOutput, removeContainer bool) (string, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return "", err
	}

	// configure and create container
	if showOutput {
		config.AttachStdout = true
		config.AttachStderr = true
		config.Tty = true
	}
	resp, err := cli.ContainerCreate(ctx, config, host, nil, nil, containerName)
	if err != nil {
		return "", err
	}

	// start container
	err = cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{})
	if err != nil {
		return "", err
	}

	if showOutput {
		// retrieve and print logs
		err := printContainerLogs(ctx, cli, containerName)
		if err != nil {
			return "", err
		}
	}

	return resp.ID, nil
}

func StopContainer(ctx context.Context, id string) error {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return err
	}
	timeout := 30 * time.Second
	info, _ := cli.ContainerInspect(ctx, id)
	output.Println("Stopping container: %s", strings.TrimPrefix(info.Name, "/"))
	return cli.ContainerStop(ctx, id, &timeout)
}

func RemoveContainer(ctx context.Context, id string) error {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return err
	}
	info, _ := cli.ContainerInspect(ctx, id)
	output.Println("Removing container: %s", strings.TrimPrefix(info.Name, "/"))
	return cli.ContainerRemove(ctx, id, types.ContainerRemoveOptions{})
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
		Follow:     true,
	})
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(reader)
	go func() {
		for scanner.Scan() {
			output.PrintlnWithNameAndColor(containerName, scanner.Text(), output.Blue)
		}
	}()

	return nil
}
