package test

import (
	"errors"

	l "github.com/jgkawell/galactus/pkg/logging/v2"
)

func Function3(logger l.Logger) l.Error {
	// add another field to logger
	logger = logger.WithField("key3", "blah3")
	// call third party function (we don't control anything in this function)
	err := thirdPartyFunction("b")
	if err != nil {
		// wrap error on return to caller so that all fields, files, and line numbers are preserved
		return logger.WrapError(l.NewError(err, "my custom error message"))
	}
	return nil
}

// thirdPartyFunction emulates a call to a third party library where you don't control the logging
func thirdPartyFunction(a string) error {
	if a == "a" {
		return nil
	}
	return errors.New("third party library error message")
}
