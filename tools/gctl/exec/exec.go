package exec

import (
	"bufio"
	"context"
	"os/exec"

	"gctl/output"
)

func ExecuteCommand(ctx context.Context, name string, c output.Color, cmd *exec.Cmd) error {
	// create a pipe for the output
	cmdReader, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}
	// create a pipe for the error output
	cmdErrReader, err := cmd.StderrPipe()
	if err != nil {
		return err
	}

	// scanner for output
	scanner := bufio.NewScanner(cmdReader)
	go func() {
		for scanner.Scan() {
			output.PrintlnWithNameAndColor(name, scanner.Text(), c)
		}
	}()

	// scanner for error output
	scannerErr := bufio.NewScanner(cmdErrReader)
	go func() {
		for scannerErr.Scan() {
			output.PrintlnWithNameAndColor(name, scannerErr.Text(), c)
		}
	}()

	// watch for done signal and kill process if received
	go func() {
		<-ctx.Done()
		cmd.Process.Kill()
	}()

	// start the command
	err = cmd.Start()
	if err != nil {
		return err
	}

	// wait for completion
	err = cmd.Wait()
	if err != nil {
		// only error if not closed by user
		if err.Error() != "signal: killed" {
			return err
		}
	}

	return nil
}
