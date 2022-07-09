# Go Logger

## Agreement

Our internal [logging module](../../pkg/logging/README.md) will be used for any Golang application logging.

## Problem

Way back when there used to be no standard for logging within our Go applications. Some developers would use [logrus](https://github.com/sirupsen/logrus) some used the built-in `log` package and none of them had any consistency on log levels or log fields. This inconsistency led to issues with application observability which in turn created massive issues in operations.

## Benefits

In response to this, our own internal `logging` module was created to create a standard way of logging within Go applications. This module wraps the commonly used `logrus` package and both reduces the number of public functions on the interface as well as adds many helper functions like creating null loggers and setting log levels. These improvements to the stock `logrus` package has helped us gain much better observability into the application by making logging consistent and properly configurable.

This package can be used in any Go application whether that be a CLI, microservice, or anything else. It can be expanded to provide further utility and modified to achieve better logging consistency. Any Go application that creates logs (which is all of them), should use our internal `logging` module.
