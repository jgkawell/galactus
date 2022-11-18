package cmd

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/archive"
	"github.com/docker/docker/pkg/jsonmessage"
	"github.com/moby/term"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("init called")

		ctx := cmd.Context()

		// build out execution path
		rootPath := viper.GetString("config.root")
		fullPath := filepath.Join(rootPath, "api")

		cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
		if err != nil {
			log.Fatal(err)
		}

		buildContext, err := getDockerContext(fullPath)
		if err != nil {
			log.Fatal(err)
		}
		opt := types.ImageBuildOptions{
			Tags: []string{"proto-builder:v3"},
		}
		resp, err := cli.ImageBuild(ctx, buildContext, opt)
		if err != nil {
			log.Fatal(err)
		}

		id, isTerm := term.GetFdInfo(os.Stdout)
		_ = jsonmessage.DisplayJSONMessagesStream(resp.Body, os.Stdout, id, isTerm, nil)

		log.Println("finished")
	},
}

func init() {
	apiCmd.AddCommand(initCmd)
}

func getDockerContext(filePath string) (io.Reader, error) {
    ctx, err := archive.TarWithOptions(filePath, &archive.TarOptions{})
	if err != nil {
		return nil, err
	}
    return ctx, nil
}

