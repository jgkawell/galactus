package context

import (
	"context"
	"errors"

	agpb "github.com/circadence-official/galactus/api/gen/go/core/aggregates/v1"
	l "github.com/circadence-official/galactus/pkg/logging/v2"

	"google.golang.org/grpc/metadata"
)

type ExecutionContext interface {
	GetLogger() l.Logger
	GetContext() context.Context
	GetTransactionID() (string)
	GetContextWithTransactionID() (context.Context)
}

type ExecutionContextWithoutTransactionID interface {
	GetLogger() l.Logger
	GetContext() context.Context
}

type executionContext struct {
	context       context.Context
	logger        l.Logger
	transactionID string
}

const (
	TransactionIDLoggerFieldKey   = "transaction_id"
	TransactionIDNotSet           = "transaction_id not set"
	TransactionIDNotSetInMetadata = "transaction_id not found in metadata of context"
	MetadataNotFoundInCTX         = "metadata object not found in context"
)

// NewExecutionContext creates a new execution context with the given transaction id.
// This function will return an error if the transaction id is empty.
func NewExecutionContext(ctx context.Context, logger l.Logger, transactionID string) (ExecutionContext, l.Error) {
	if transactionID == "" {
		return nil, logger.WrapError(errors.New(TransactionIDNotSet))
	}

	logger = logger.WithField(TransactionIDLoggerFieldKey, transactionID)
	return &executionContext{
		context:       ctx,
		logger:        logger,
		transactionID: transactionID,
	}, nil
}

// NewExecutionContextWithoutTransactionID creates a new execution context WITHOUT a transaction id
// NOTE: you should only call this rarely. most everything should have a transaction id.
func NewExecutionContextWithoutTransactionID(ctx context.Context, logger l.Logger) ExecutionContextWithoutTransactionID {
	return &executionContext{
		context:       ctx,
		logger:        logger,
		transactionID: "",
	}
}

// NewExecutionContextFromEvent creates a new execution context from an event and pulls the transaction id from the event.
// This function will return an error if the transaction id is empty.
func NewExecutionContextFromEvent(ctx context.Context, logger l.Logger, event *agpb.Event) (ExecutionContext, l.Error) {
	transactionID := event.GetTransactionId()
	if transactionID == "" {
		return nil, logger.WrapError(errors.New(TransactionIDNotSet))
	}

	logger = logger.WithField(TransactionIDLoggerFieldKey, transactionID)
	return &executionContext{
		context:       ctx,
		logger:        logger,
		transactionID: transactionID,
	}, nil
}

// NewExecutionContextFromMetadata creates a new execution context and pulls the transaction id from the metadata of the gRPC context.
// This function will return an error if the transaction id is empty.
// NOTE: Only call this when handling a gRPC request.
func NewExecutionContextFromContextWithMetadata(ctx context.Context, logger l.Logger) (ExecutionContext, l.Error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, logger.WrapError(errors.New(MetadataNotFoundInCTX))
	}

	idList := md["transaction_id"]
	if len(idList) == 0 {
		return nil, logger.WrapError(errors.New(TransactionIDNotSetInMetadata))
	}
	transID := idList[0]

	logger = logger.WithField(TransactionIDLoggerFieldKey, transID)
	return &executionContext{
		context:       ctx,
		logger:        logger,
		transactionID: transID,
	}, nil
}

func (c *executionContext) GetLogger() l.Logger {
	return c.logger
}

func (c *executionContext) GetContext() context.Context {
	return c.context
}

// GetTransactionID returns the transaction id of the context.
func (c *executionContext) GetTransactionID() (string) {
	return c.transactionID
}

// GetExecutionContextWithTransactionId returns a new context with transaction_id set in the metadata of the context.
func (c *executionContext) GetContextWithTransactionID() (context.Context) {
	return metadata.NewOutgoingContext(c.GetContext(), metadata.Pairs("transaction_Id", c.GetTransactionID()))
}
