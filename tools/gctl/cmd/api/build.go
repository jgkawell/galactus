package api

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"gctl/docker"
	e "gctl/exec"
	"gctl/files"
	"gctl/output"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	protoContainer = "proto-builder"
)

func Build(cmd *cobra.Command, args []string) (err error) {
	ctx := cmd.Context()
	dctl, err := docker.NewDockerController()
	if err != nil {
		return nil
	}

	// build out execution path
	rootPath := viper.GetString("config.root")
	apiPath := filepath.Join(rootPath, "api")

	// clean the api directory
	err = Clean(cmd, args)
	if err != nil {
		// don't print as it already has been
		return err
	}

	// run docker proto-builder image
	output.Println("Building api...")

	// base configuration for docker container runs
	config := &container.Config{
		Image:      protoImage,
		WorkingDir: "/workspace",
	}
	hostConfig := &container.HostConfig{
		Mounts: []mount.Mount{
			{
				Type:   mount.TypeBind,
				Source: apiPath,
				Target: "/workspace",
			},
		},
	}

	// mod update
	config.Cmd = []string{"mod", "update", "protos"}
	err = dctl.RunContainer(ctx, protoContainer, config, hostConfig, true, true)
	if err != nil {
		output.Error(err)
		return err
	}

	// generate
	config.Cmd = []string{"generate", "protos"}
	err = dctl.RunContainer(ctx, protoContainer, config, hostConfig, true, true)
	if err != nil {
		output.Error(err)
		return err
	}

	// if running on linux, we need to chown the files to the current user
	c := exec.Command("uname", "-s")
	o, err := e.ExecuteCommandReturnStdout(ctx, c)
	if err != nil {
		output.Error(err)
		return err
	}
	if strings.TrimSpace(o) == "Linux" {
		c = exec.Command("sudo", "chown", "-R", "1000:1000", apiPath)
		err = e.ExecuteCommand(ctx, "chown", output.Magenta, c)
		if err != nil {
			output.Error(err)
			return err
		}
	}

	// extra steps for go
	processor := files.NewProcessor(goFileCleanup)
	processor.Start(filepath.Join(apiPath, "gen", "go"))
	processor.Wait()
	c = exec.Command("go", "mod", "tidy")
	c.Dir = filepath.Join(apiPath, "gen", "go")
	err = e.ExecuteCommand(ctx, "go", output.Green, c)
	if err != nil {
		output.Error(err)
		return err
	}

	output.Println("Finished")
	return err
}

// goFileCleanup fixes issues with generated golang protobuf files
func goFileCleanup(filePath string) {
	// protoc-gen-gorm messes up this import by "nulling" it out
	if strings.HasSuffix(filePath, ".pb.gorm.go") {
		old := "_ google.golang.org/protobuf/types/known/timestamppb"
		new := "google.golang.org/protobuf/types/known/timestamppb"
		sed(old, new, filePath)
	}
}

// deleteAllSubDirectories removes all subdirectories and their children without affecting the top-level non-directory files
func deleteAllSubDirectories(path string) error {
	dir, err := os.ReadDir(path)
	if err != nil {
		return err
	}
	for _, d := range dir {
		if d.IsDir() {
			os.RemoveAll(filepath.Join(path, d.Name()))
		}
	}
	return nil
}

// sed replicates the basic functionality of the unix/linux sed command
func sed(old, new, filePath string) error {
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}
	fileString := string(fileData)
	fileString = strings.ReplaceAll(fileString, old, new)
	fileData = []byte(fileString)
	err = os.WriteFile(filePath, fileData, 0o600)
	if err != nil {
		return err
	}
	return nil
}
