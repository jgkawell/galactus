package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var (
	consumerServiceName   string
	consumerAggregateName string
	consumerCommandName   string
	consumerOutputPath    string
	consumerTemplatePath  string
	consumerCommonPath    string
)

func init() {
	flags := generateConsumerTemplateCmd.Flags()
	flags.StringVarP(&consumerServiceName, "service_name", "s", "", "name of the new consumer")
	flags.StringVarP(&consumerAggregateName, "aggregate_name", "a", "", "name of the aggregate the service will act on")
	flags.StringVarP(&consumerCommandName, "command_name", "n", "", "name of the (first) command the service handles")
	flags.StringVarP(&consumerOutputPath, "output_path", "o", "generated", "location to place the new consumer dir, and all of it's new files")
	flags.StringVarP(&consumerTemplatePath, "template_path", "t", "templates/consumer", "service type templates for hammer to use`")
	flags.StringVarP(&consumerCommonPath, "common_path", "c", "templates/common", "common templates for hammer to use")

	_ = generateConsumerTemplateCmd.MarkFlagRequired("service_name")
	_ = generateConsumerTemplateCmd.MarkFlagRequired("aggregate_name")
	_ = generateConsumerTemplateCmd.MarkFlagRequired("command_name")
}

var generateConsumerTemplateCmd = &cobra.Command{
	Use:   "consumer",
	Short: "given the name of the new consumer, generate all the files that are needed to create a new consumer",
	Example: `hammer consumer \
	--service_name="GameOfficiating" \
	--aggregate_name="Football" \
	--command_name="PauseGame" \
	--output_path="../../internal" \
	--template_path="templates/consumer" \
	--common_path="templates/common"`,
	RunE: func(cmd *cobra.Command, args []string) error {
		consumer := &TemplateValues{
			Name:      consumerServiceName,
			Aggregate: consumerAggregateName,
			Command:   consumerCommandName,
		}

		basePath := filepath.Join(consumerOutputPath, strings.ToLower(consumerServiceName))
		fmt.Printf("creating consumer service at: %s\n", basePath)

		if err := os.MkdirAll(basePath, os.ModePerm); err != nil {
			panic(err)
		}

		// write common
		if err := WriteTemplate(basePath, consumerCommonPath, consumer); err != nil {
			panic(err)
		}

		// write process specific stuff
		if err := WriteTemplate(basePath, consumerTemplatePath, consumer); err != nil {
			panic(err)
		}

		return nil
	},
}
