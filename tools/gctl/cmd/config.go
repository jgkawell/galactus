package cmd

import (
	"gctl/output"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var configCommand = &cobra.Command{
	Use:   "config",
	Short: "Initialize the gctl configuration",
	Long: `Initialize the gctl configuration. Be sure to run
this from the root of the galactus repository`,
	RunE: func(cmd *cobra.Command, args []string) error {
		cwd, err := os.Getwd()
		if err != nil {
			output.Error(err)
			return err
		}
		output.Println("Setting the root directory to: %s", cwd)
		viper.Set("config.root", cwd)
		viper.WriteConfig()
		return nil
	},
}

func init() {
	// add parent
	rootCmd.AddCommand(configCommand)
}
