package diagrams

import (
	"os"
	"path/filepath"

	"gctl/output"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func Clean(cmd *cobra.Command, args []string) error {
	// build out execution path
	rootPath := viper.GetString("config.root")
	fullPath := filepath.Join(rootPath, "docs", "diagrams")

	cleanGeneratedFiles(fullPath)

	output.Println("Finished")
	return nil
}

// cleanGeneratedFiles recursively deletes all directories name "gen" within a given path
func cleanGeneratedFiles(filePath string) error {
	files, err := os.ReadDir(filePath)
	if err != nil {
		return err
	}
	for _, file := range files {
		if file.IsDir() {
			if file.Name() == "gen" {
				err = os.RemoveAll(filepath.Join(filePath, file.Name()))
				if err != nil {
					return err
				}
				continue
			}
			// Recursively go further in the tree
			err = cleanGeneratedFiles(filepath.Join(filePath, file.Name()))
			if err != nil {
				return err
			}
		}
	}
	return nil
}
