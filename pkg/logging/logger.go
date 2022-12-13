package logging

import (
	"context"
	"runtime"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

const ddTraceIDKey = "dd.trace_id"
const ddSpanIDKey = "dd.span_id"

// Logger is an interface wrapping around logrus to provide
// context based logging for tracing
type Logger interface {
	// AddGlobalField adds a key:value pair to every log written with this logger
	// This is helpful for service-wide values
	AddGlobalField(key string, val string)
	// SetLevel sets the logging level for the logger
	SetLevel(level logrus.Level)
	// GetEntry returns a copy of the logrus Entry used in the logger
	// NOTE: Only use the base logrus.Entry when absolutely necessary.
	// Logging should really be done through the CustomLogger wrapper,
	// NOT through logrus directly.
	GetEntry() *logrus.Entry
	// WrapError wraps a standard Go error (OR `logging.Error`) into a custom error with the
	// additional context of current logger fields and call stack information.
	WrapError(error) Error

	// WithHTTPContext injects tracing IDs into logs from context
	// Since it returns type CustomLogger you can continue dot operations
	// after the call. For example:
	//
	//  logger.WithHTTPContext(ctx).WithField("animal", "dog").Warn("Invalid species")
	WithHTTPContext(ctx *gin.Context) *customLogger
	// WithRPCContext injects tracing IDs into logs from context
	// Since it returns type CustomLogger you can continue dot operations
	// after the call. For example:
	//
	//     logger.WithRPCContext(ctx).WithField("animal", "dog").Warn("Invalid species")
	//
	// This function should be called as the first step in all Handler functions. For example:
	//
	//     func (h handlerImpl) GetContent(ctx context.Context, req *pb.GetContentRequest) (response *pb.GetContentResponse, err error) {
	//       logger := h.logger.WithRPCContext(ctx)
	WithRPCContext(ctx context.Context) *customLogger
	// WithError - Add an error as single field (using the key defined in ErrorKey) to the Entry.
	WithError(err error) *customLogger
	// WithContext - Add a context to the Entry.
	WithContext(ctx context.Context) *customLogger
	// WithField - Add a single field to the Entry.
	WithField(key string, value interface{}) *customLogger
	// WithFields - Add a map of fields to the Entry.
	WithFields(fields Fields) *customLogger
	// WithTime - Overrides the time of the Entry.
	WithTime(t time.Time) *customLogger

	// Trace - Definition:
	// "Seriously, WTF is going on here?!?!
	// I need to log every single statement I execute to find this @#$@ing memory corruption bug before I go insane"
	Trace(msg string)
	// Debug - Definition:
	// Off by default, able to be turned on for debugging specific unexpected problems.
	// This is where you might log detailed information about key method parameters or
	// other information that is useful for finding likely problems in specific 'problematic' areas of the code.
	Debug(msg string)
	// Info - Definition:
	// Normal logging that's part of the normal operation of the app;
	// diagnostic stuff so you can go back and say 'how often did this broad-level operation happen?',
	// or 'how did the user's data get into this state?'
	Info(msg string)
	// Warn - Definition:
	// something that's concerning but not causing the operation to abort;
	// # of connections in the DB pool getting low, an unusual-but-expected timeout in an operation, etc.
	// Think of 'WARN' as something that's useful in aggregate; e.g. grep, group,
	// and count them to get a picture of what's affecting the system health
	Warn(msg string)
	// Error - Definition:
	// something that the app's doing that it shouldn't.
	// This isn't a user error ('invalid search query');
	// it's an assertion failure, network problem, etc etc.,
	// probably one that is going to abort the current operation
	Error(msg string)
	// WrappedError - Definition:
	// this is a convenience method that calls Error() but makes sure to wrap the error a final time
	// so that all current call context is included in the error. This has the same output as:
	//   logger.WithFields(logger.WrapError(err).Fields()).WithError(logger.WrapError(err)).Error("failed to process request")
	// but instead has a much simpler oneliner of:
	//   logger.WrappedError(err, "failed to process request")
	WrappedError(Error, string) *customLogger
	// Fatal - Definition:
	// the app (or at the very least a thread) is about to die horribly.
	// This is where the info explaining why that's happening goes.
	Fatal(msg string)
	// Panic - Definition:
	// Be careful about calling this vs Fatal:
	// - For Fatal level, the log message goes to the configured log output, while panic is only going to write to stderr.
	// - Panic will print a stack trace, which may not be relevant to the error at all.
	// - Defers will be executed when a program panics, but calling os.Exit exits immediately, and deferred functions can't be run.
	// In general, only use panic for programming errors, where the stack trace is important to the context of the error.
	// If the message isn't targeted at the programmer, you're simply hiding the message in superfluous data.
	Panic(msg string)
}

// customLogger wraps around logrus with added functionality
type customLogger struct {
	entry *logrus.Entry
}

// Fields type, used to pass to WithFields()
type Fields map[string]interface{}

// newCustomLogger creates a Logger with defined logrus Entry
func newCustomLogger(e *logrus.Entry) customLogger {
	return customLogger{entry: e}
}

// AddGlobalField adds a key:value pair to every log written with this logger
// This is helpful for service-wide values
func (l *customLogger) AddGlobalField(key string, val string) {
	l.WithField(key, val).Info("Adding field to global logger")
	l.entry = l.entry.WithField(key, val)
}

// SetLevel sets the logging level for the logger
func (l *customLogger) SetLevel(level logrus.Level) {
	l.entry.Logger.SetLevel(level)
}

// GetEntry returns a copy of the logrus Entry used in the logger
// NOTE: Only use the base logrus.Entry when absolutely necessary.
// Logging should really be done through the CustomLogger wrapper,
// NOT through logrus directly.
func (l *customLogger) GetEntry() *logrus.Entry {
	return l.entry
}

// WrapError wraps a standard Go error (OR `logging.Error`) into a custom error with the
// additional context of current logger fields and call stack information.
func (l *customLogger) WrapError(err error) Error {
	wrappedError := wrap(err, Fields(l.entry.Data))
	if wrappedError == nil {
		l.entry.Error("failed to wrap error")
		// return custom error but without any wrapping
		return customError{
			cause: err,
		}
	}
	return wrappedError
}

// WithError - Add an error as single field (using the key defined in ErrorKey) to the Entry.
func (l *customLogger) WithError(err error) *customLogger {
	newLogger := newCustomLogger(l.entry.WithError(err))
	return &newLogger
}

// WithContext - Add a context to the Entry.
func (l *customLogger) WithContext(ctx context.Context) *customLogger {
	newLogger := newCustomLogger(l.entry.WithContext(ctx))
	return &newLogger
}

// WithField - Add a single field to the Entry.
func (l *customLogger) WithField(key string, value interface{}) *customLogger {
	newLogger := newCustomLogger(l.entry.WithField(key, value))
	return &newLogger
}

// WithFields - Add a map of fields to the Entry.
func (l *customLogger) WithFields(fields Fields) *customLogger {
	// Copy custom fields into logrus fields
	logrusFields := logrus.Fields{}
	for index, element := range fields {
		logrusFields[index] = element
	}
	// Create new logger with given fields
	newLogger := newCustomLogger(l.entry.WithFields(logrusFields))
	return &newLogger
}

// WithTime - Overrides the time of the Entry.
func (l *customLogger) WithTime(t time.Time) *customLogger {
	newLogger := newCustomLogger(l.entry.WithTime(t))
	return &newLogger
}

// WithHTTPContext injects tracing IDs into logs from context
// Since it returns type CustomLogger you can continue dot operations
// after the call. For example:
//
//  logger.WithHTTPContext(ctx).WithField("animal", "dog").Warn("Invalid species")
func (l *customLogger) WithHTTPContext(ctx *gin.Context) *customLogger {
	span, found := tracer.SpanFromContext(ctx.Request.Context())
	if found {
		return l.WithFields(Fields{
			ddTraceIDKey: span.Context().TraceID(),
			ddSpanIDKey:  span.Context().SpanID(),
		})
	}
	l.withLogContext().entry.Debug("Failed to find find span from HTTP context for logger")
	newLogger := newCustomLogger(l.entry)
	return &newLogger
}

// WithRPCContext injects tracing IDs into logs from context
// Since it returns type CustomLogger you can continue dot operations
// after the call. For example:
//
//     logger.WithRPCContext(ctx).WithField("animal", "dog").Warn("Invalid species")
//
// This function should be called as the first step in all Handler functions. For example:
//
//     func (h handlerImpl) GetContent(ctx context.Context, req *pb.GetContentRequest) (response *pb.GetContentResponse, err error) {
//       logger := h.logger.WithRPCContext(ctx)
func (l *customLogger) WithRPCContext(ctx context.Context) *customLogger {
	span, found := tracer.SpanFromContext(ctx)
	if found {
		return l.WithFields(Fields{
			ddTraceIDKey: span.Context().TraceID(),
			ddSpanIDKey:  span.Context().SpanID(),
		})
	}
	l.withLogContext().entry.Debug("Failed to find find span from RPC context for logger")
	newLogger := newCustomLogger(l.entry)
	return &newLogger
}

// withLogContext makes sure that the external logger module uses
// the function name of the caller instead of CustomLogger (ex. Debug())
func (l *customLogger) withLogContext() *customLogger {
	pc, _, _, ok := runtime.Caller(2)
	if !ok {
		return nil
	}
	fn := runtime.FuncForPC(pc).Name()
	functionName := fn[strings.LastIndex(fn, ".")+1:]
	return l.WithField("function", functionName)
}

// Entry Print family functions

// Trace - Definition:
// "Seriously, WTF is going on here?!?!
// I need to log every single statement I execute to find this @#$@ing memory corruption bug before I go insane"
func (l *customLogger) Trace(msg string) {
	l.withLogContext().entry.Trace(msg)
}

// Debug - Definition:
// Off by default, able to be turned on for debugging specific unexpected problems.
// This is where you might log detailed information about key method parameters or
// other information that is useful for finding likely problems in specific 'problematic' areas of the code.
func (l *customLogger) Debug(msg string) {
	l.withLogContext().entry.Debug(msg)
}

// Info - Definition:
// Normal logging that's part of the normal operation of the app;
// diagnostic stuff so you can go back and say 'how often did this broad-level operation happen?',
// or 'how did the user's data get into this state?'
func (l *customLogger) Info(msg string) {
	l.withLogContext().entry.Info(msg)
}

// Warn - Definition:
// something that's concerning but not causing the operation to abort;
// # of connections in the DB pool getting low, an unusual-but-expected timeout in an operation, etc.
// Think of 'WARN' as something that's useful in aggregate; e.g. grep, group,
// and count them to get a picture of what's affecting the system health
func (l *customLogger) Warn(msg string) {
	l.withLogContext().entry.Warn(msg)
}

// Error - Definition:
// something that the app's doing that it shouldn't.
// This isn't a user error ('invalid search query');
// it's an assertion failure, network problem, etc etc.,
// probably one that is going to abort the current operation
func (l *customLogger) Error(msg string) {
	l.withLogContext().entry.Error(msg)
}

// WrappedError - Definition:
// this is a convenience method that calls Error() but makes sure to wrap the error a final time
// so that all current call context is included in the error. This has the same output as:
//   logger.WithFields(logger.WrapError(err).Fields()).WithError(logger.WrapError(err)).Error("failed to process request")
// but instead has a much simpler oneliner of:
//   logger.WrappedError(err, "failed to process request")
func (l *customLogger) WrappedError(err Error, msg string) *customLogger {
	err = wrap(err, Fields(l.entry.Data))
	l.WithFields(err.Fields()).WithError(err).withLogContext().entry.Error(msg)
	return l
}

// Fatal - Definition:
// the app (or at the very least a thread) is about to die horribly.
// This is where the info explaining why that's happening goes.
func (l *customLogger) Fatal(msg string) {
	l.withLogContext().entry.Fatal(msg)
}

// Panic - Definition:
// Be careful about calling this vs Fatal:
// - For Fatal level, the log message goes to the configured log output, while panic is only going to write to stderr.
// - Panic will print a stack trace, which may not be relevant to the error at all.
// - Defers will be executed when a program panics, but calling os.Exit exits immediately, and deferred functions can't be run.
// In general, only use panic for programming errors, where the stack trace is important to the context of the error.
// If the message isn't targeted at the programmer, you're simply hiding the message in superfluous data.
func (l *customLogger) Panic(msg string) {
	l.withLogContext().entry.Panic(msg)
}
