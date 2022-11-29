package chassis

import (
	"context"
	"net/http"
	"time"

	"github.com/jgkawell/galactus/pkg/chassis/db"
	"github.com/jgkawell/galactus/pkg/chassis/terminator"
	l "github.com/jgkawell/galactus/pkg/logging/v2"

	"github.com/DeanThompson/ginpprof"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	gintrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/gin-gonic/gin"
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

	// tracing middleware for datadog
	b.httpRouter.Use(gintrace.Middleware(b.viper.GetString("traceName")))
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

func (b *mainBuilder) StartHttpServer() {
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

func (b *mainBuilder) StopHttpServer() {
	if b.httpServer != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := b.httpServer.Shutdown(ctx); err != nil {
			b.logger.WithError(err).Error("failed to shut down http server")
		}
		b.logger.Info("http server stopped")
	}
}

func (b *mainBuilder) healthHandler(ctx *gin.Context) {
	logger := b.logger.WithHTTPContext(ctx)
	logger.Trace("liveness handler called")
	if b.noSqlClient != nil {
		if err := db.PingNoSqlClient(b.noSqlClient); err != nil {
			msg := "failed to ping mongo client"
			logger.WithError(err).Error(msg)
			ctx.JSON(http.StatusFailedDependency, gin.H{"message": msg})
		} else {
			msg := "Health check: succeeded in pinging mongo client"
			logger.Trace(msg)
			ctx.JSON(http.StatusOK, gin.H{})
		}
	}

	if b.sqlClient != nil {
		if err := db.PingSqlClient(logger, b.sqlClient); err != nil {
			msg := "failed to ping sql client"
			logger.WithError(err).Error(msg)
			ctx.JSON(http.StatusFailedDependency, gin.H{"message": msg})
		} else {
			msg := "Health check: succeeded in pinging sql client"
			logger.Trace(msg)
			ctx.JSON(http.StatusOK, gin.H{})
		}
	}

	if b.wellnessCheckConfig != nil {
		b.wellnessCheckConfig.Check(ctx)
	}
}

func (b *mainBuilder) readinessHandler(ctx *gin.Context) {
	logger := b.logger.WithHTTPContext(ctx)
	logger.Trace("readiness handler called")
	if b.noSqlClient != nil {
		if err := db.PingNoSqlClient(b.noSqlClient); err != nil {
			msg := "failed to ping mongo client"
			logger.WithError(err).Error(msg)
			ctx.JSON(http.StatusFailedDependency, gin.H{"message": msg})
		} else {
			msg := "Readiness check: succeeded in pinging mongo client"
			logger.Trace(msg)
			ctx.JSON(http.StatusOK, gin.H{})
		}
	}

	if b.sqlClient != nil {
		if err := db.PingSqlClient(logger, b.sqlClient); err != nil {
			msg := "failed to ping sql client"
			logger.WithError(err).Error(msg)
			ctx.JSON(http.StatusFailedDependency, gin.H{"message": msg})
		} else {
			msg := "Health check: succeeded in pinging sql client"
			logger.Trace(msg)
			ctx.JSON(http.StatusOK, gin.H{})
		}
	}

	if b.readinessCheckConfig != nil {
		b.readinessCheckConfig.Check(ctx)
	}
}

// serviceStatusNotification is responsible for handling the POST requests from the shawarma sidecar when
// the state of the assigned endpoints/services change
func (b *mainBuilder) serviceStatusNotification(ctx *gin.Context) {
	logger := b.logger.WithHTTPContext(ctx)
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
	logger = logger.WithFields(l.Fields{
		"status":      state.Status,
		"k8s_service": state.Service,
	})
	logger.Info("parsed new status")

	switch state.Status {
	case statusActive:
		logger.Info("service is active")
		// start new consumers based on the service name
		b.initializeBrokerListeners(state.Service)
	case statusInactive:
		logger.Info("service is inactive")
		// stop existing consumer channels based on the service name
		b.cancelBrokerListeners(state.Service)
	default:
		logger.Error("unknown status")
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	ctx.Status(http.StatusOK)
}
