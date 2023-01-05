package main

import (
	"fmt"

	"error-handling/test"

	l "github.com/jgkawell/galactus/pkg/logging"
)

/*
Example output (formatted for readability):
{
  "call_stack": [
    "SOME_PATH/galactus/examples/go/error-handling/main.go:46",
    "SOME_PATH/galactus/examples/go/error-handling/main.go:60",
    "SOME_PATH/galactus/examples/go/error-handling/main.go:72",
    "SOME_PATH/galactus/examples/go/error-handling/test/test.go:16"
  ],
  "error": "[main.main]->[main.Function1]->[main.Function2]->[error-handling/test.Function3]->[my custom error message: third party library error message]",
  "function": "main",
  "key0": "blah0",
  "key1": "blah1",
  "key2": "blah2",
  "key3": "blah3",
  "level": "error",
  "msg": "failed to process request",
  "service": "main",
  "time": "2022-12-16T11:56:11-06:00"
}
Error as returned to client: [main.Function1]->[main.Function2]->[error-handling/test.Function3]->[my custom error message: third party library error message]
*/

func main() {
	// initialize logger
	logger := l.CreateLogger("info", "main")

	// add top level fields
	logger = logger.WithField("key0", "blah0")

	// call function and pass logger with added fields
	err := Function1(logger)

	if err != nil {
		// log error at top of call stack ONLY
		// NOTE: fields here match the logger fields at the lowest point in the call stack (Function3)
		logger.WrappedError(err, "failed to process request")

		// Error as returned to calling client
		fmt.Printf("Error as returned to client: %v\n", err)
	}
}

func Function1(logger l.Logger) l.Error {
	// add another field to logger
	logger = logger.WithField("key1", "blah1")
	// call next function and pass logger with added fields
	err := Function2(logger)
	if err != nil {
		// wrap error on return to caller so that all fields, files, and line numbers are preserved
		return logger.WrapError(err)
	}
	return nil
}

func Function2(logger l.Logger) l.Error {
	// add another field to logger
	logger = logger.WithField("key2", "blah2")
	// call next function and pass logger with added fields
	err := test.Function3(logger)
	if err != nil {
		// wrap error on return to caller so that all fields, files, and line numbers are preserved
		return logger.WrapError(err)
	}
	return nil
}
