package cmd

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"errors"

	"github.com/adrg/xdg"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// NOTE: Don't try to share CLI arguments between commands here as defaults will get overwritten by other commands at runtime: https://github.com/spf13/cobra/issues/1047
var (
	configFile        string
	defaultConfigFile string
)

const (
	bin_name = "hammer"
)

func init() {
	var err error

	defaultConfigFile, err = xdg.ConfigFile(fmt.Sprintf("%s/config.toml", bin_name))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	_, err = os.Stat(defaultConfigFile)
	// assume the error is b/c the file does not exist
	if err != nil {
		// create config file
		f, err := os.Create(defaultConfigFile)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if f.Name() == "" {
			err = errors.New("failed to configure hammer")
			fmt.Println(err)
			os.Exit(1)
		}
	}

	cobra.OnInitialize(initConfig)

	RootCmd.AddCommand(generateConsumerTemplateCmd)
	RootCmd.AddCommand(generateRpcTemplateCmd)
	RootCmd.AddCommand(generateAggregateServiceTemplateCmd)
}

var RootCmd = &cobra.Command{
	Use:              "hammer",
	Short:            "A tool to generate a base template for galactus microservices",
	TraverseChildren: true,
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Usage()
	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func initConfig() {
	if configFile != "" {
		viper.SetConfigFile(configFile)
	} else {
		viper.SetConfigFile(defaultConfigFile)
	}
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if viper.GetBool("dev-mode") {
		fmt.Println("Executing in dev mode...")
	}
}

type TemplateValues struct {
	Name        string
	Aggregate   string
	Command     string
	OutputPath  string
	Directories []string
}

func (t TemplateValues) GetTitle() string {
	return strings.Title(t.Name)
}

func (t TemplateValues) GetLower() string {
	return strings.ToLower(t.Name)
}

func (t TemplateValues) GetUpper() string {
	return strings.ToUpper(t.Name)
}

func (t TemplateValues) GetAggregateTitle() string {
	return strings.Title(t.Aggregate)
}

func (t TemplateValues) GetAggregateLower() string {
	return strings.ToLower(t.Aggregate)
}

func (t TemplateValues) GetAggregateUpper() string {
	return strings.ToUpper(t.Aggregate)
}

func (t TemplateValues) GetCommandTitle() string {
	return strings.Title(t.Command)
}

func (t TemplateValues) GetCommandLower() string {
	return strings.ToLower(t.Command)
}

func (t TemplateValues) GetCommandUpper() string {
	return strings.ToUpper(t.Command)
}

func WriteTemplate(basePath, path string, templateValues *TemplateValues) error {
	err := filepath.WalkDir(path, func(s string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// am I a dir, if not
		if !d.IsDir() {
			// parse template
			tmp, err := template.ParseFiles(s)
			if err != nil {
				panic(err)
			}

			// remove template path, and append to template slice
			s = strings.ReplaceAll(s, path, "")
			// clean `.tmpl` from s
			s = strings.ReplaceAll(s, ".tmpl", "")

			// ensure parent directory is created before writing to
			// the file
			path := filepath.Join(basePath, s)
			dirName := filepath.Dir(path)
			if _, err = os.Stat(dirName); err != nil {
				err := os.MkdirAll(dirName, os.ModePerm)
				if err != nil {
					panic(err)
				}
			}

			file, err := os.Create(path)
			if err != nil {
				panic(err)
			}

			// execute template
			err = tmp.Execute(file, templateValues)
			if err != nil {
				panic(err)
			}

		}

		return nil
	})

	return err
}
