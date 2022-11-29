package run

import (
	"path/filepath"

	"gctl/docker"
	"gctl/output"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// TODO: have a command to remove generated Docker images?

func Init(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()

	rootPath := viper.GetString("config.root")

	// build custom rabbitmq image
	output.Println("Building custom RabbitMQ Docker image...")
	fullPath := filepath.Join(rootPath, "third_party", "rabbitmq")
	err := docker.BuildImage(ctx, fullPath, "rabbitmq:galactus")
	if err != nil {
		output.Error(err)
		return err
	}

	output.Println("Finished")
	return nil
}
