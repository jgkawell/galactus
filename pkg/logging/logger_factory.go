package logging

import (
	"net/http"
	"os"
	"time"

	runtime "github.com/banzaicloud/logrus-runtime-formatter"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/test"
)

var globalLogger *CustomLogger

// CreateLogger creates a service level logger
// level = the log level to instantiate the logger with
//   Possible values for level are "panic", "fatal", "error", "warn", "warning", "info", "debug", and "trace"
// service = the service name to include with all logs
func CreateLogger(level string, service string) *CustomLogger {
	logrus.SetFormatter(&runtime.Formatter{
		ChildFormatter: &logrus.JSONFormatter{
			DisableHTMLEscape: true,
		},
	})
	logrus.SetOutput(os.Stdout)

	// Add service field
	newEntry := logrus.WithField("service", service)

	// Create new logger
	newLogger := NewCustomLogger(newEntry)
	globalLogger = &newLogger

	// Set starting log level and return
	setLogLevel(globalLogger, level)
	globalLogger.Info("Starting")
	return globalLogger
}

// CreateNullLogger creates a logger for testing that wraps the null logger provided by logrus
func CreateNullLogger() (*CustomLogger, *test.Hook) {
	nullLogger, logHook := test.NewNullLogger()
	newLogger := NewCustomLogger(nullLogger.WithField("", ""))
	globalLogger = &newLogger
	return globalLogger, logHook
}

// RegisterHTTPEndpointsWithGin registers the log changing and viewing endpoints with the Gin router
// Possible values for level are "panic", "fatal", "error", "warn", "warning", "info", "debug", and "trace"
// - POST to set log level: /log?level=<LEVEL>
// - GET to retrieve the current log level
func RegisterHTTPEndpointsWithGin(router *gin.Engine) {
	router.POST("/log", func(c *gin.Context) {
		logger := globalLogger.WithHTTPContext(c)
		level := c.Query("level")
		logger.WithField("current_log_level", level).Info("Trying to set logging level")
		if err := setLogLevel(logger, level); err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
		} else {
			c.JSON(http.StatusOK, "log level set to "+level)
		}
	})
	router.GET("/log", func(c *gin.Context) {
		logger := globalLogger.WithHTTPContext(c)
		level := globalLogger.GetEntry().Logger.Level.String()
		logger.WithField("current_log_level", level).Info("Log level")
		c.JSON(http.StatusOK, level)
	})
}

// GinMiddleware injects the custom logger with traces and http data fields
func GinMiddleware(quietRoutes []string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Measure request duration
		start := time.Now()
		ctx.Next()
		duration := time.Now().Sub(start)
		logger := globalLogger.WithHTTPContext(ctx).WithFields(Fields{
			"status_code":    ctx.Writer.Status(),
			"duration":       duration,
			"request_method": ctx.Request.Method,
			"request_uri":    ctx.Request.RequestURI,
		})
		// Log quiet routes at trace to keep excessive logging down
		if contains(quietRoutes, ctx.Request.RequestURI) {
			logger.Trace("")
		} else {
			logger.Info("")
		}
	}
}

// setLogLevel is a helper method for setting the logger's logging level
// Possible values for level are "panic", "fatal", "error", "warn", "warning", "info", "debug", and "trace"
// These values are in decreasing value so if you set the level to "info" you also get "debug" and "trace"
func setLogLevel(logger Logger, level string) error {
	logger.WithField("requested_log_level", level).Info("Attempting to set log level")

	// Get the log level from the given string
	logLevel, err := logrus.ParseLevel(level)
	if err != nil {
		// Unable to parse the given level
		curLevel := globalLogger.GetEntry().Logger.Level.String()
		logger.WithFields(Fields{
			"requested_log_level": level,
			"current_log_level":   curLevel,
		}).Info("Unknown log level requested. Log level will not change.")
		return err
	}

	// Able to parse the level, set the logger's log level
	globalLogger.SetLevel(logLevel)
	logger.WithField("current_log_level", level).Info("Global logging level set")
	return nil
}

// contains checks if a given string is contained within a given array
func contains(array []string, value string) bool {
	for _, element := range array {
		if element == value {
			return true
		}
	}
	return false
}
