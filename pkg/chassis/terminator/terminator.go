// terminator package is broken out so there are no circular dependencies between different parts of MainBuilder.
package terminator

import (
	"os"
	"syscall"
)

// ApplicationChannel is an application-level channel mainly used to signal termination of the application.
var ApplicationChannel = make(chan os.Signal)

// TerminateApplication will kill the application when called.
func TerminateApplication() {
	ApplicationChannel <- syscall.SIGTERM
}
