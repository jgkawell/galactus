package context

import (
	"context"
	"runtime"
	"strings"
	"time"

	"github.com/jgkawell/galactus/pkg/logging"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

type (
	Context interface {
		context.Context

		Logger() logging.Logger
		Span() (Context, trace.Span)
		SpanByName(name string) (Context, trace.Span)

		WithError(err error) Context
		WithField(key string, value interface{}) Context
		WithFields(fields logging.Fields) Context
	}
	contextImpl struct {
		context context.Context
		logger  logging.Logger
		tracer  trace.Tracer
		span    trace.Span
		fields  logging.Fields
	}
	Telemetry struct {
		Logger logging.Logger
		Tracer trace.Tracer
	}
)

// Context methods

func New(ctx context.Context, telemetry Telemetry) Context {
	return contextImpl{
		context: ctx,
		logger:  telemetry.Logger.WithContext(ctx),
		tracer:  telemetry.Tracer,
		fields:  logging.Fields{},
	}
}

func NewBackground(telemetry Telemetry) Context {
	return New(context.Background(), telemetry)
}

func (c contextImpl) Logger() logging.Logger {
	return c.logger
}

func (c contextImpl) SpanByName(name string) (Context, trace.Span) {
	ctx, span := c.tracer.Start(c, name)
	for key, value := range c.fields {
		span.SetAttributes(
			attribute.String(key, value.(string)),
		)
	}
	c.context = ctx
	c.logger = c.logger.WithContext(ctx)
	c.span = span
	return c, span
}

func (c contextImpl) Span() (Context, trace.Span) {
	return c.SpanByName(getCallerName())
}

// With methods

func (c contextImpl) WithError(err error) Context {
	c.logger = c.logger.WithError(err)
	if c.span != nil {
		c.span.RecordError(err, trace.WithStackTrace(true))
		c.span.SetStatus(codes.Error, err.Error())
	}
	return c
}

func (c contextImpl) WithField(key string, value interface{}) Context {
	c.logger = c.logger.WithField(key, value)
	c.fields[key] = value
	if c.span != nil {
		c.span.SetAttributes(
			attribute.String(key, value.(string)),
		)
	}
	return c
}

func (c contextImpl) WithFields(fields logging.Fields) Context {
	c.logger = c.logger.WithFields(fields)
	if c.span != nil {
		for key, value := range fields {
			c.fields[key] = value
			c.span.SetAttributes(
				attribute.String(key, value.(string)),
			)
		}
	}
	return c
}

// context.Context interface methods

func (c contextImpl) Deadline() (deadline time.Time, ok bool) {
	return c.context.Deadline()
}

func (c contextImpl) Done() <-chan struct{} {
	return c.context.Done()
}

func (c contextImpl) Err() error {
	return c.context.Err()
}

func (c contextImpl) Value(key any) any {
	return c.context.Value(key)
}

// helpers

func getCallerName() string {
	pc, _, _, ok := runtime.Caller(2)
	if !ok {
		return "unknown"
	}
	fn := runtime.FuncForPC(pc).Name()
	return fn[strings.LastIndex(fn, ".")+1:]
}
