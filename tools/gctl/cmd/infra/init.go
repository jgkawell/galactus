package infra

import (
	"path/filepath"

	"gctl/docker"
	"gctl/output"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func Init(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()
	dctl, err := docker.NewDockerController()
	if err != nil {
		return nil
	}

	rootPath := viper.GetString("config.root")

	// build custom rabbitmq image
	output.Println("Building custom RabbitMQ Docker image...")
	fullPath := filepath.Join(rootPath, "third_party", "rabbitmq")
	err = dctl.BuildImage(ctx, fullPath, rabbitImage)
	if err != nil {
		output.Error(err)
		return err
	}

	// pull needed images
	output.Println("Pulling Mongo Docker image...")
	err = dctl.PullImage(ctx, mongoImage)
	if err != nil {
		output.Error(err)
		return err
	}
	output.Println("Pulling Postgres Docker image...")
	err = dctl.PullImage(ctx, postgresImage)
	if err != nil {
		output.Error(err)
		return err
	}
	output.Println("Pulling Hasura Docker image...")
	err = dctl.PullImage(ctx, hasuraImage)
	if err != nil {
		output.Error(err)
		return err
	}

	output.Println("Finished")
	return nil
}
