package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"gctl/docker"
	"gctl/files"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// buildCmd represents the build command
var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("build called")

		var err error

		ctx := cmd.Context()

		// build out execution path
		rootPath := viper.GetString("config.root")
		apiPath := filepath.Join(rootPath, "api")

		fmt.Println("running: clean")

		// remove buf.lock
		os.Remove(filepath.Join(apiPath, "buf.lock"))
		// clean docs
		fmt.Println("cleaning docs...")
		os.Remove(filepath.Join(apiPath, "gen", "docs", "docs.md"))
		os.Remove(filepath.Join(apiPath, "gen", "docs", "index.html"))
		// clean go
		fmt.Println("cleaning go...")
		err = deleteAllSubDirectories(filepath.Join(apiPath, "gen", "go"))
		if err != nil {
			log.Fatal(err)
		}
		// clean openapiv2
		fmt.Println("cleaning openapiv2...")
		err = deleteAllSubDirectories(filepath.Join(apiPath, "gen", "openapiv2"))
		if err != nil {
			log.Fatal(err)
		}
		// clean web
		fmt.Println("cleaning web...")
		err = deleteAllSubDirectories(filepath.Join(apiPath, "gen", "web"))
		if err != nil {
			log.Fatal(err)
		}

		// run docker proto-builder image
		fmt.Println("running: build")

		// base configuration for docker container runs
		config := &container.Config{
			Image:      "proto-builder:v3",
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
		config.Cmd = []string{"mod", "update"}
		err = docker.RunContainer(ctx, "proto-builder", config, hostConfig, true, true)
		if err != nil {
			log.Fatal(err)
		}

		// generate
		config.Cmd = []string{"generate"}
		err = docker.RunContainer(ctx, "proto-builder", config, hostConfig, true, true)
		if err != nil {
			log.Fatal(err)
		}

		// extra steps for go
		processor := files.NewProcessor(goFileCleanup)
		processor.Start(filepath.Join(apiPath, "gen", "go"))
		processor.Wait()

		// extra steps for web
		processor = files.NewProcessor(webFileCleanup)
		processor.Start(filepath.Join(apiPath, "gen", "web"))
		processor.Wait()

		fmt.Println("finished")
	},
}

func init() {
	apiCmd.AddCommand(buildCmd)
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

// webFileCleanup fixes issues with generated web protobuf files
func webFileCleanup(filePath string) {
	// these imports aren't actually needed and so are removed
	if strings.HasSuffix(filePath, ".d.ts") {
		old := "import * as validate_validate_pb from '../../../validate/validate_pb';"
		new := ""
		sed(old, new, filePath)
	}
	if strings.HasSuffix(filePath, "_pb.js") {
		old := "var validate_validate_pb = require('../../../validate/validate_pb.js');\ngoog.object.extend(proto, validate_validate_pb);\n"
		new := ""
		sed(old, new, filePath)

		old = "var validate_validate_pb = require('../../../validate/validate_pb.js')\n"
		new = ""
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
