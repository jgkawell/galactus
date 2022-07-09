package terminator

import (
	"syscall"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTerminateApplicationSuccess(t *testing.T) {
	go func() {
		TerminateApplication()
	}()
	assert.Equal(t, syscall.SIGTERM, <-ApplicationChannel)
}

func TestTerminateApplicationFailed(t *testing.T) {
	go func() {
		TerminateApplication()
		ApplicationChannel <- nil
	}()
	assert.Equal(t, syscall.SIGTERM, <-ApplicationChannel)
	assert.NotEqual(t, syscall.SIGTERM, <-ApplicationChannel)
}
