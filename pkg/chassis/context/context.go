package context

import (
	"context"
	"errors"

	agpb "github.com/jgkawell/galactus/api/gen/go/core/aggregates/v1"

	l "github.com/jgkawell/galactus/pkg/logging/v2"

	"google.golang.org/grpc/metadata"
)

type ExecutionContext struct {
	context       context.Context
	Logger        l.Logger
	transactionID string
}

type ExecutionContextWithoutTransactionID struct {
	context       context.Context
	Logger        l.Logger
}

const (
	TransactionIDLoggerFieldKey   = "transaction_id"
	TransactionIDNotSet           = "transaction_id not set"
	TransactionIDNotSetInMetadata = "transaction_id not found in metadata of context"
	MetadataNotFoundInCTX         = "metadata object not found in context"
)

// NewExecutionContext creates a new execution context with the given transaction id.
// This function will return an error if the transaction id is empty.
func NewExecutionContext(ctx context.Context, logger l.Logger, transactionID string) (*ExecutionContext, l.Error) {
	if transactionID == "" {
		return nil, logger.WrapError(errors.New(TransactionIDNotSet))
	}

	logger = logger.WithField(TransactionIDLoggerFieldKey, transactionID)
	return &ExecutionContext{
		context:       ctx,
		Logger:        logger,
		transactionID: transactionID,
	}, nil
}

// NewExecutionContextWithoutTransactionID creates a new execution context WITHOUT a transaction id
// NOTE: you should only call this rarely. most everything should have a transaction id.
func NewExecutionContextWithoutTransactionID(ctx context.Context, logger l.Logger) *ExecutionContextWithoutTransactionID {
	return &ExecutionContextWithoutTransactionID{
		context:       ctx,
		Logger:        logger,
	}
}

// NewExecutionContextFromEvent creates a new execution context from an event and pulls the transaction id from the event.
// This function will return an error if the transaction id is empty.
func NewExecutionContextFromEvent(ctx context.Context, logger l.Logger, event *agpb.Event) (*ExecutionContext, l.Error) {
	transactionID := event.GetTransactionId()
	if transactionID == "" {
		return nil, logger.WrapError(errors.New(TransactionIDNotSet))
	}

	logger = logger.WithField(TransactionIDLoggerFieldKey, transactionID)
	return &ExecutionContext{
		context:       ctx,
		Logger:        logger,
		transactionID: transactionID,
	}, nil
}

// NewExecutionContextFromMetadata creates a new execution context and pulls the transaction id from the metadata of the gRPC context.
// This function will return an error if the transaction id is empty.
// NOTE: Only call this when handling a gRPC request.
func NewExecutionContextFromContextWithMetadata(ctx context.Context, logger l.Logger) (*ExecutionContext, l.Error) {
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
	return &ExecutionContext{
		context:       ctx,
		Logger:        logger,
		transactionID: transID,
	}, nil
}

func (c *ExecutionContext) GetContext() context.Context {
	return metadata.NewOutgoingContext(c.context, metadata.Pairs("transaction_Id", c.GetTransactionID()))
}

func (c *ExecutionContextWithoutTransactionID) GetContext() context.Context {
	return c.context
}

func (c *ExecutionContext) GetTransactionID() (string) {
	return c.transactionID
}
