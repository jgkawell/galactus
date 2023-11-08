package chassis

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/jgkawell/galactus/pkg/chassis/terminator"
	l "github.com/jgkawell/galactus/pkg/logging"

	"github.com/DeanThompson/ginpprof"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// these constants define the status values that are POSTed by the shawarma sidecar
const (
	statusActive   = "active"
	statusInactive = "inactive"
)

// serviceState is the struct used by the shawarma sidecar when POSTing the status of the application
type serviceState struct {
	Status  string `json:"status"`
	Service string `json:"service"`
}

// createHttpServer creates the http server
func (b *mainBuilder) createHttpServer() {
	if !b.isDevMode {
		gin.SetMode(gin.ReleaseMode)
	}

	b.httpRouter = gin.New()

	// in dev mode add pprof
	if b.isDevMode {
		ginpprof.Wrap(b.httpRouter)
	}

	// tracing middleware
	b.httpRouter.Use(l.GinMiddleware([]string{"/health", "/readiness", "/metrics", "/applicationstate"}))

	// k8s health and readiness checks
	b.httpRouter.GET("/health", b.healthHandler)
	b.httpRouter.GET("/readiness", b.readinessHandler)

	// prometheus metrics
	b.httpRouter.GET("/metrics", func(c *gin.Context) { promhttp.Handler().ServeHTTP(c.Writer, c.Request) })

	// this route is used by the shawarma sidecar to check the status of the application
	b.httpRouter.POST("/applicationstate", b.serviceStatusNotification)

	// register the logger endpoints for getting and setting log levels
	l.RegisterHTTPEndpointsWithGin(b.httpRouter)
}

func (b *mainBuilder) startHttpServer() {
	b.logger.WithField("port", b.httpPort).Info("starting http server")

	b.httpServer = &http.Server{Addr: "localhost:" + b.httpPort}
	b.httpServer.Handler = b.httpRouter

	if err := b.httpServer.ListenAndServe(); err != nil {
		if err == http.ErrServerClosed {
			b.logger.Info("http server closed")
		} else {
			b.logger.WithError(err).Error("http server failed")
		}
	}
	terminator.TerminateApplication()
}

func (b *mainBuilder) stopHttpServer() {
	if b.httpServer != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := b.httpServer.Shutdown(ctx); err != nil {
			b.logger.WithError(err).Error("failed to shut down http server")
		}
		b.logger.Info("http server stopped")
	}
}

func (b *mainBuilder) healthHandler(c *gin.Context) {
	logger := b.logger.WithContext(c)
	logger.Trace("liveness handler called")

	for _, db := range b.databases {
		if err := db.Ping(c.Request.Context()); err != nil {
			msg := "failed to ping database"
			logger.WithError(err).Error(msg)
			c.JSON(http.StatusFailedDependency, gin.H{"message": msg})
			return
		}
	}

	if b.wellnessCheckConfig != nil {
		b.wellnessCheckConfig.Check(c)
	}

	logger.Trace("liveness check succeeded")
	c.JSON(http.StatusOK, gin.H{})
}

func (b *mainBuilder) readinessHandler(c *gin.Context) {
	logger := b.logger.WithContext(c)
	logger.Trace("readiness handler called")

	for _, db := range b.databases {
		if err := db.Ping(c.Request.Context()); err != nil {
			msg := "failed to ping database"
			logger.WithError(err).Error(msg)
			c.JSON(http.StatusFailedDependency, gin.H{"message": msg})
			return
		}
	}

	if b.readinessCheckConfig != nil {
		b.readinessCheckConfig.Check(c)
	}

	logger.Trace("readiness check succeeded")
	c.JSON(http.StatusOK, gin.H{})
}

// serviceStatusNotification is responsible for handling the POST requests from the shawarma sidecar when
// the state of the assigned endpoints/services change
func (b *mainBuilder) serviceStatusNotification(ctx *gin.Context) {
	logger := b.logger.WithContext(ctx)
	logger.Info("received status notification")

	// parse out the status from the request
	var state serviceState
	err := ctx.BindJSON(&state)
	if err != nil {
		logger.WithError(err).Error("failed to bind JSON for status notification")
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	// k8s_service = the name of the k8s endpoint/service (e.g. "users" or "users-preview")
	// TODO: how does this handle canary deployments?
	logger = logger.WithFields(l.Fields{
		"status":      state.Status,
		"k8s_service": state.Service,
	})
	logger.Info("parsed new status")

	switch state.Status {
	case statusActive:
		logger.Info("service is active")
		b.subscribeConsumers(preview(state.Service))
	case statusInactive:
		logger.Info("service is inactive")
		// stop existing consumer channels based on the service name
		b.unsubscribeConsumers(preview(state.Service))
	default:
		logger.Error("unknown status")
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	ctx.Status(http.StatusOK)
}

// HELPERS

func preview(service string) bool {
	return strings.Contains(service, "preview")
}
