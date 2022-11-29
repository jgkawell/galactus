package diagrams

import (
	"fmt"
	"path/filepath"

	"gctl/docker"
	"gctl/output"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	diagramsContainer = "diagrams-builder"
)

var (
	Type string
)

func Build(cmd *cobra.Command, args []string) (err error) {
	ctx := cmd.Context()
	dctl, err := docker.NewDockerController()
	if err != nil {
		return nil
	}

	// build out execution path
	rootPath := viper.GetString("config.root")
	diagramsPath := filepath.Join(rootPath, "docs", "diagrams")

	// base configuration for docker container runs
	config := &container.Config{
		Image:      diagramsImage,
		WorkingDir: "/workspace",
	}
	hostConfig := &container.HostConfig{
		Mounts: []mount.Mount{
			{
				Type:   mount.TypeBind,
				Source: diagramsPath,
				Target: "/workspace",
			},
		},
	}

	fileType := fmt.Sprintf("-t%s", Type)
	config.Cmd = []string{"plantuml", fileType, "-duration", "-progress", "-recurse", "-o", "gen/", "./**.puml"}
	err = dctl.RunContainer(ctx, diagramsContainer, config, hostConfig, true, true)
	if err != nil {
		output.Error(err)
		return err
	}

	output.Println("Finished")
	return err
}
