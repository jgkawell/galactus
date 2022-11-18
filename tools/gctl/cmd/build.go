package cmd

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/client"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// buildCmd represents the build command
var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("build called")

		var (
			err error
		)

		ctx := cmd.Context()

		// build out execution path
		rootPath := viper.GetString("config.root")
		apiPath := filepath.Join(rootPath, "api")

		log.Println("running: clean")

		// remove buf.lock
		os.Remove(filepath.Join(apiPath, "buf.lock"))
		// clean docs
		log.Println("cleaning docs...")
		os.Remove(filepath.Join(apiPath, "gen", "docs", "docs.md"))
		os.Remove(filepath.Join(apiPath, "gen", "docs", "index.html"))
		// clean go
		log.Println("cleaning go...")
		err = deleteAllSubDirectories(filepath.Join(apiPath, "gen", "go"))
		if err != nil {
			log.Fatal(err)
		}
		// clean openapiv2
		log.Println("cleaning openapiv2...")
		err = deleteAllSubDirectories(filepath.Join(apiPath, "gen", "openapiv2"))
		if err != nil {
			log.Fatal(err)
		}
		// clean web
		log.Println("cleaning web...")
		err = deleteAllSubDirectories(filepath.Join(apiPath, "gen", "web"))
		if err != nil {
			log.Fatal(err)
		}

		// run builder
		log.Println("running builder...")

		cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
		if err != nil {
			log.Fatal(err)
		}

		hostConfig := container.HostConfig{
			Mounts: []mount.Mount{
				{
					Type:   mount.TypeBind,
					Source: apiPath,
					Target: "/workspace",
				},
			},
		}

		// TODO: consolidate below Docker steps into helper function

		// mod update
		containerConfig := &container.Config{
			Image:      "proto-builder:v3",
			WorkingDir: "/workspace",
			Cmd:        []string{"mod", "update"},
		}
		resp, err := cli.ContainerCreate(ctx, containerConfig, &hostConfig, nil, nil, "proto-builder")
		if err != nil {
			log.Fatal(err)
		}
		err = cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{})
		if err != nil {
			log.Fatal(err)
		}
		_, err = waitForContainer(ctx, cli, resp.ID)
		if err != nil {
			log.Fatal(err)
		}
		err = cli.ContainerRemove(ctx, resp.ID, types.ContainerRemoveOptions{})
		if err != nil {
			log.Fatal(err)
		}
		// TODO: tail logs from container

		// generate
		containerConfig = &container.Config{
			Image:      "proto-builder:v3",
			WorkingDir: "/workspace",
			Cmd:        []string{"generate"},
		}
		resp, err = cli.ContainerCreate(ctx, containerConfig, &hostConfig, nil, nil, "proto-builder")
		if err != nil {
			log.Fatal(err)
		}
		err = cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{})
		if err != nil {
			log.Fatal(err)
		}
		_, err = waitForContainer(ctx, cli, resp.ID)
		if err != nil {
			log.Fatal(err)
		}
		err = cli.ContainerRemove(ctx, resp.ID, types.ContainerRemoveOptions{})
		if err != nil {
			log.Fatal(err)
		}
		// TODO: tail logs from container

		// TODO: move the below to a helper and pass a function in for processing?

		// extra steps for go
		// TODO: adding this breaks the web cleanup (probably because we're reusing the waitgroup)
		// for w := 1; w <= runtime.NumCPU(); w++ {
		// 	go loopFilesWorker(goFileCleanup)
		// }
		// //Start the recursion
		// LoopDirsFiles(filepath.Join(apiPath, "gen", "go"))
		// wg.Wait()

		// extra steps for web

		for w := 1; w <= runtime.NumCPU(); w++ {
			go loopFilesWorker(webFileCleanup)
		}
		//Start the recursion
		LoopDirsFiles(filepath.Join(apiPath, "gen", "web"))
		wg.Wait()

		log.Println("finished")
	},
}

func webFileCleanup(filePath string) {
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

func goFileCleanup(filePath string) {
	if strings.HasSuffix(filePath, ".pb.gorm.go") {
		old := "_ google.golang.org/protobuf/types/known/timestamppb"
		new := "google.golang.org/protobuf/types/known/timestamppb"
		sed(old, new, filePath)
	}
}

func init() {
	apiCmd.AddCommand(buildCmd)
}

func deleteAllSubDirectories(path string) error {
	dir, err := ioutil.ReadDir(path)
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

func waitForContainer(ctx context.Context, cli *client.Client, id string) (state int64, err error) {
	resultC, errC := cli.ContainerWait(ctx, id, "")
	select {
	case err := <-errC:
		return 0, err
	case result := <-resultC:
		return result.StatusCode, nil
	}
}

var (
    wg   sync.WaitGroup
    jobs chan string = make(chan string)
)

func loopFilesWorker(fileProcessor func(filePath string)) error {
    for path := range jobs {
        files, err := ioutil.ReadDir(path)
        if err != nil {
            wg.Done()
            return err
        }

        for _, file := range files {
            if !file.IsDir() {
                fileProcessor(filepath.Join(path, file.Name()))
            }
        }
        wg.Done()
    }
    return nil
}

func LoopDirsFiles(path string) error {
    files, err := ioutil.ReadDir(path)
    if err != nil {
        return err
    }
    //Add this path as a job to the workers
    //You must call it in a go routine, since if every worker is busy, then you have to wait for the channel to be free.
    go func() {
		wg.Add(1)
        jobs <- path
    }()
    for _, file := range files {
        if file.IsDir() {
            //Recursively go further in the tree
            LoopDirsFiles(filepath.Join(path, file.Name()))
        }
    }
    return nil
}

func sed(old, new, filePath string) error {
	fileData, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	fileString := string(fileData)
	fileString = strings.ReplaceAll(fileString, old, new)
	fileData = []byte(fileString)

	err = ioutil.WriteFile(filePath, fileData, 0o600)
	if err != nil {
		return err
	}

	return nil
}
