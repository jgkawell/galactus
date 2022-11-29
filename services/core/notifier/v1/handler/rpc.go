package handler

import (
	"context"

	s "notifier/service"

	ntpb "github.com/jgkawell/galactus/api/gen/go/core/notifier/v1"

	ct "github.com/jgkawell/galactus/pkg/chassis/context"
	l "github.com/jgkawell/galactus/pkg/logging/v2"

	"github.com/google/uuid"
)

type notifierHandler struct {
	logger  l.Logger
	service s.Service
}

// NewNotifierHandler - Implements the `ntpb.NotifierServer` interface for server side streaming of notifications
// to a connected client. The `server` is a long running process so it's stored and messages are process in a seperate
// goroutine
func NewNotifierHandler(logger l.Logger, service s.Service) ntpb.NotifierServer {
	return &notifierHandler{logger, service}
}

// Connect - Implements the `Connect` method of the `ntpb.NotifierServer` interface
func (h *notifierHandler) Connect(req *ntpb.ConnectionRequest, stream ntpb.Notifier_ConnectServer) error {
	logger := h.logger.WithFields(l.Fields{
		"actor_id":  req.GetActorId(),
		"client_id": req.GetClientId(),
	})
	logger.Debug("connection requested")

	ctx, err := ct.NewExecutionContext(context.Background(), logger, uuid.NewString())
	if err != nil {
		logger.WrappedError(err, "failed to create execution context. this should never happen here as we are creating a new transaction id ourselves")
	}

	conn, err := h.service.Connect(*ctx, req, stream)
	if err != nil {
		ctx.Logger.WithError(err).Error("failed to connect")
	}

	return <-conn.Error
}
