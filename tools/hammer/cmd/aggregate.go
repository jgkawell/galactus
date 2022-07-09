package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var (
	aggregateServiceName   string
	aggregateOutputPath    string
	aggregateTemplatePath  string
	aggregateCommonPath    string
)

func init() {
	flags := generateAggregateServiceTemplateCmd.Flags()

	flags.StringVarP(&aggregateServiceName, "aggregate_name", "a", "", "name of the aggregate this will service")
	flags.StringVarP(&aggregateOutputPath, "output_path", "o", "generated", "location to place the new consumer dir, and all of it's new files")
	flags.StringVarP(&aggregateTemplatePath, "template_path", "t", "templates/aggregate", "service type templates for hammer to use`")
	flags.StringVarP(&aggregateCommonPath, "common_path", "c", "templates/common", "common templates for hammer to use")

	_ = generateAggregateServiceTemplateCmd.MarkFlagRequired("aggregate_name")
}

var generateAggregateServiceTemplateCmd = &cobra.Command{
	Use:   "aggregate",
	Short: "given the name of the new aggregate service, generate all the files that are needed to create a new aggregate service",
	Example: `hammer aggregate \
	--aggregate_name="Football" \
	--output_path="../../internal" \
	--template_path="templates/aggregate" \
	--common_path="templates/common"`,
	RunE: func(cmd *cobra.Command, args []string) error {
		aggregate := &TemplateValues{
			Name:        aggregateServiceName,
			Aggregate:   aggregateServiceName,
			OutputPath:  aggregateOutputPath,
			Directories: []string{"handler", "service"},
		}

		basePath := filepath.Join(aggregateOutputPath, strings.ToLower(aggregateServiceName))
		fmt.Printf("creating aggregate service at: %s\n", basePath)

		if err := os.MkdirAll(basePath, os.ModePerm); err != nil {
			panic(err)
		}

		// write common
		if err := WriteTemplate(basePath, aggregateCommonPath, aggregate); err != nil {
			panic(err)
		}

		// write process specific stuff
		if err := WriteTemplate(basePath, aggregateTemplatePath, aggregate); err != nil {
			panic(err)
		}
		return nil
	},
}
