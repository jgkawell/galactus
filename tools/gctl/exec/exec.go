package exec

import (
	"bufio"
	"context"
	"fmt"
	"os/exec"
)

func ExecuteCommand(ctx context.Context, cmd *exec.Cmd) (err error) {
	// create a pipe for the output
	cmdReader, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}

	// scanner for output
	scanner := bufio.NewScanner(cmdReader)
	go func() {
		for scanner.Scan() {
			fmt.Println(scanner.Text())
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
