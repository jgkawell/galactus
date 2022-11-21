/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// cleanCmd represents the clean command
var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("clean called")

		var (
			err error
		)


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
	},
}

func init() {
	apiCmd.AddCommand(cleanCmd)
}
