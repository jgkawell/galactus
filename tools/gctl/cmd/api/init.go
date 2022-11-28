package api

import (
	"io"
	"os"
	"path/filepath"

	"gctl/output"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/archive"
	"github.com/docker/docker/pkg/jsonmessage"
	"github.com/moby/term"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// TODO: have a command to remove generated Docker images?

func Init(cmd *cobra.Command, args []string) (err error) {
	output.Println("init called")

	ctx := cmd.Context()

	// build out execution path
	rootPath := viper.GetString("config.root")
	fullPath := filepath.Join(rootPath, "api")

	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		output.Error(err)
		return err
	}

	buildContext, err := getDockerContext(fullPath)
	if err != nil {
		output.Error(err)
		return err
	}
	opt := types.ImageBuildOptions{
		Tags: []string{"proto-builder:v3"},
	}
	resp, err := cli.ImageBuild(ctx, buildContext, opt)
	if err != nil {
		output.Error(err)
		return err
	}

	id, isTerm := term.GetFdInfo(os.Stdout)
	_ = jsonmessage.DisplayJSONMessagesStream(resp.Body, os.Stdout, id, isTerm, nil)

	output.Println("Finished")
	return err
}

func getDockerContext(filePath string) (io.Reader, error) {
	ctx, err := archive.TarWithOptions(filePath, &archive.TarOptions{})
	if err != nil {
		return nil, err
	}
	return ctx, nil
}
