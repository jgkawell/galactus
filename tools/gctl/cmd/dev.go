package cmd

import (
	"gctl/cmd/dev"

	"github.com/spf13/cobra"
)

var reportCmd = &cobra.Command{
	Use:   "report",
	Short: "Uses Go tooling to generate reports on a go module",
	Long: `Uses Go tooling to generate reports on a go module.
It uses the following commands: gofmt, go vet, gocyclo, ineffassign, and misspell

You must have these tools installed before running:
  - gofmt: installed with go
  - go vet: installed with go
  - gocyclo: go install github.com/fzipp/gocyclo/cmd/gocyclo@latest
  - ineffassign: go install github.com/gordonklaus/ineffassign@latest
  - misspell: go install github.com/client9/misspell/cmd/misspell@latest
`,
	RunE: dev.Report,
}

func init() {
	rootCmd.AddCommand(reportCmd)
}
