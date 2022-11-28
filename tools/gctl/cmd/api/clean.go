package api

import (
	"os"
	"path/filepath"

	"gctl/output"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func Clean(cmd *cobra.Command, args []string) (err error) {
	output.Println("Cleaning `api/`...")

	// build out execution path
	rootPath := viper.GetString("config.root")
	apiPath := filepath.Join(rootPath, "api")

	// remove buf.lock
	os.Remove(filepath.Join(apiPath, "buf.lock"))
	// clean docs
	output.Println("Cleaning docs...")
	os.Remove(filepath.Join(apiPath, "gen", "docs", "docs.md"))
	os.Remove(filepath.Join(apiPath, "gen", "docs", "index.html"))
	// clean go
	output.Println("Cleaning go...")
	err = deleteAllSubDirectories(filepath.Join(apiPath, "gen", "go"))
	if err != nil {
		output.Error(err)
		return err
	}
	// clean openapiv2
	output.Println("Cleaning openapiv2...")
	err = deleteAllSubDirectories(filepath.Join(apiPath, "gen", "openapiv2"))
	if err != nil {
		output.Error(err)
		return err
	}
	// clean web
	output.Println("Cleaning web...")
	err = deleteAllSubDirectories(filepath.Join(apiPath, "gen", "web"))
	if err != nil {
		output.Error(err)
		return err
	}

	return err
}
