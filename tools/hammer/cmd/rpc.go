package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var (
	rpcServiceName   string
	rpcAggregateName string
	rpcCommandName   string
	rpcOutputPath    string
	rpcTemplatePath  string
	rpcCommonPath    string
)

func init() {
	flags := generateRpcTemplateCmd.Flags()
	flags.StringVarP(&rpcServiceName, "service_name", "s", "", "name of the new rpc service")
	flags.StringVarP(&rpcAggregateName, "aggregate_name", "a", "", "name of the aggregate the service will act on")
	flags.StringVarP(&rpcCommandName, "command_name", "n", "", "name of the (first) command the service handles")
	flags.StringVarP(&rpcOutputPath, "output_path", "o", "generated", "location to place the new consumer dir, and all of it's new files")
	flags.StringVarP(&rpcTemplatePath, "template_path", "t", "templates/rpc", "service type templates for hammer to use`")
	flags.StringVarP(&rpcCommonPath, "common_path", "c", "templates/common", "common templates for hammer to use")

	_ = generateRpcTemplateCmd.MarkFlagRequired("consumer_name")
	_ = generateRpcTemplateCmd.MarkFlagRequired("aggregate_name")
	_ = generateRpcTemplateCmd.MarkFlagRequired("command_name")
}

var generateRpcTemplateCmd = &cobra.Command{
	Use:   "rpc",
	Short: "given the name of the new rpc service, generate all the files that are needed to create a new rpc service",
	Example: `hammer rpc \
	--service_name="GameOfficiating" \
	--aggregate_name="Football" \
	--command_name="PauseGame" \
	--output_path="../../internal" \
	--template_path="templates/rpc" \
	--common_path="templates/common"`,
	RunE: func(cmd *cobra.Command, args []string) error {
		consumer := &TemplateValues{
			Name:      rpcServiceName,
			Aggregate: rpcAggregateName,
			Command:   rpcCommandName,
		}

		fmt.Println("creating rpc service")

		basePath := filepath.Join(rpcOutputPath, strings.ToLower(rpcServiceName))
		fmt.Printf("creating rpc service at: %s\n", basePath)

		if err := os.MkdirAll(basePath, os.ModePerm); err != nil {
			panic(err)
		}

		// write common
		if err := WriteTemplate(basePath, rpcCommonPath, consumer); err != nil {
			panic(err)
		}

		// write process specific stuff
		if err := WriteTemplate(basePath, rpcTemplatePath, consumer); err != nil {
			panic(err)
		}
		return nil
	},
}
