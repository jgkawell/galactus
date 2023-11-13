package dev

import (
	"os/exec"

	e "gctl/exec"
	"gctl/output"

	"github.com/spf13/cobra"
)

func Report(cmd *cobra.Command, args []string) (err error) {
	ctx := cmd.Context()

	// run gofmt
	output.Println("Running gofmt (improperly formatted files will be listed)...")
	err = e.ExecuteCommand(ctx, "gofmt", output.Cyan, exec.Command("gofmt", "-l", "."))
	if err != nil {
		return err
	}

	// run go vet
	output.Println("Running go vet...")
	// squash error because go vet returns non-zero exit code when it finds issues
	_ = e.ExecuteCommand(ctx, "go vet", output.Cyan, exec.Command("go", "vet", "./..."))

	// run gocyclo
	output.Println("Running gocyclo (functions with cyclomatic complexity > 15 will be listed)...")
	// squash error because gocyclo returns non-zero exit code when it finds issues
	_ = e.ExecuteCommand(ctx, "gocyclo", output.Cyan, exec.Command("gocyclo", "-over", "15", "."))

	// run ineffassign
	output.Println("Running ineffassign...")
	// squash error because ineffassign returns non-zero exit code when it finds issues
	_ = e.ExecuteCommand(ctx, "ineffassign", output.Cyan, exec.Command("ineffassign", "./..."))

	// run misspell
	output.Println("Running misspell...")
	err = e.ExecuteCommand(ctx, "misspell", output.Cyan, exec.Command("misspell", "."))
	if err != nil {
		return err
	}

	return err
}
