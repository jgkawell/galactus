package handler

import (
	"context"

	s "notifier/service"

	ntpb "github.com/circadence-official/galactus/api/gen/go/core/notifier/v1"
	l "github.com/circadence-official/galactus/pkg/logging/v2"
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
	h.logger.WithField("connection_request", req).Debug("establish a connection")

	conn, err := h.service.Connect(context.Background(), h.logger, req, stream)
	if err != nil {
		h.logger.WithError(err).Error("failed to connect")
	}

	return <-conn.Error
}
