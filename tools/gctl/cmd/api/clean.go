package api

import (
	"os"
	"path/filepath"

	"gctl/output"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func Clean(cmd *cobra.Command, args []string) (err error) {
	output.Println("Cleaning api...")

	// build out execution path
	rootPath := viper.GetString("config.root")
	apiPath := filepath.Join(rootPath, "api")

	// remove buf.lock
	os.Remove(filepath.Join(apiPath, "buf.lock"))
	// clean go
	output.Println("Cleaning go...")
	err = deleteAllSubDirectories(filepath.Join(apiPath, "gen", "go"))
	if err != nil {
		output.Error(err)
		return err
	}

	return err
}
