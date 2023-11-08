package chassis

import (
	// standard libraries
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"
	"time"

	// chassis packages
	cf "github.com/jgkawell/galactus/pkg/chassis/clientfactory"
	ct "github.com/jgkawell/galactus/pkg/chassis/context"
	"github.com/jgkawell/galactus/pkg/chassis/database"
	ec "github.com/jgkawell/galactus/pkg/chassis/env"
	"github.com/jgkawell/galactus/pkg/chassis/events"
	messagebus "github.com/jgkawell/galactus/pkg/chassis/messagebus"
	"github.com/jgkawell/galactus/pkg/chassis/secrets"
	"github.com/jgkawell/galactus/pkg/chassis/terminator"

	// other galactus modules
	rgpb "github.com/jgkawell/galactus/api/gen/go/core/registry/v1"
	l "github.com/jgkawell/galactus/pkg/logging"

	// third party libraries

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/uptrace/uptrace-go/uptrace"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// MainBuilder is an interface that exposes the functionality for using the chassis module
type MainBuilder interface {

	// CONTROL FUNCTIONS

	// Run starts the execution of the application.
	Run()
	// Close releases all assets. Call via defer.
	Close()

	// PRIVATE FUNCTIONS

	// startHttpServer launches the http server. This will consume the calling thread.
	startHttpServer()
	// Stop HttpServer shuts the http server down.
	stopHttpServer()
	// startRpcServer launches the rpc server. This will consume the calling thread.
	startRpcServer()
	// stopRpcServer shuts the rpc server down.
	stopRpcServer()

	// GETTER FUNCTIONS

	// IsDevMode specifies if this application is running in development mode (true) or production mode (false)
	IsDevMode() bool
	// GetCacheClient gets the cache implementation client.
	GetCacheClient() *redis.Client
	// GetConfig exposes the viper configuration for the application.
	GetConfig() *viper.Viper
	// GetLogger exposes the logger the application is using.
	GetLogger() l.Logger
	// GetHttpRouter exposes the gin http router.
	GetHttpRouter() *gin.Engine
	// GetRegistryClient exposes the registry client.
	GetRegistryClient() rgpb.RegistryClient
	// GetRpcServer exposes the grpc server.
	GetRpcServer() *grpc.Server
	// GetSecretsClient exposes the secret vault client.
	GetSecretsClient() secrets.Client
	// GetEventManager... TODO
	GetEventManager() events.EventManager
	GetTelemetry() ct.Telemetry
}

// EventConfig defines the configuration for events
type EventConfig struct{}

// DaoLayerConfig defines the function for initializing the DAO layer if any custom setup is required
type DaoLayerConfig struct {
	// CreateDaoLayer creates the dao layer
	CreateDaoLayer func(b MainBuilder)
}

// DatabaseConfig defines the connection configuration to all databases
type DatabaseConfig struct {
	Databases []database.Client
}

// CacheConfig defines the configuration for the cache layer (only Redis is currently supported)
type CacheConfig struct {
	// CacheAddress host of the cache instance
	CacheAddress string
	// CacheSecret the secret/password for the cache instance
	CacheSecret string
}

// GatewayLayerConfig specifies how the gateway layer will be configured.
type GatewayLayerConfig struct {
	// CreateInternalClients creates the gateway layer using the client factory.
	CreateInternalClients func(b MainBuilder, clientFactory cf.ClientFactory) []*grpc.ClientConn
}

// ServiceLayerConfig specifies how the service layer will be configured.
type ServiceLayerConfig struct {
	// CreateServiceLayer creates the service layer using the client factory.
	CreateServiceLayer func(b MainBuilder)
}

type GrpcHandlers struct {
	Desc    grpc.ServiceDesc
	Handler interface{}
}

// HandlerLayerConfig specifies how the handler layer will be configured.
type HandlerLayerConfig struct {
	// HttpPortConfigKey will only be used if DoNotRegisterService is true. It specifies the config key to read the port from.
	// ONLY USE THIS FOR THE REGISTRY SERVICE ITSELF OR IF YOU REALLY KNOW WHAT YOU'RE DOING.
	HttpPortConfigKey string
	// CreateRestHandlers creates the http restful interface using the gin-engine router. This needs to
	// be a function called from main.go so that it has both the service interfaces (
	// and access to mainbuilder attributes (e.g. the logger).
	CreateRestHandlers func(b MainBuilder)
	// GrpcPortConfigKey will only be used if DoNotRegisterService is true. It specifies the config key to read the port from.
	// ONLY USE THIS FOR THE REGISTRY SERVICE ITSELF OR IF YOU REALLY KNOW WHAT YOU'RE DOING.
	GrpcPortConfigKey string
	// CreateRpcHandlers creates the rpc interface using the grpc server. This needs to
	// be a function called from main.go so that it has both the service interfaces
	// and access to mainbuilder attributes (e.g. the logger).
	CreateRpcHandlers func(b MainBuilder) []GrpcHandlers
	// RpcOptions is a slice of optional grpc.ServerOption structs to set on the gRPC server.
	RpcOptions []grpc.ServerOption
	// CreateConsumers defines a function that creates the consumer interfaces on the messagebus.
	// This needs to be a function called from main.go so that it has both the service interfaces
	// and access to mainbuilder attributes (e.g. the logger).
	CreateConsumers func(b MainBuilder) []ConsumerConfig
}

// ConsumerConfig defines the configuration for a consumer handler (processing messages off of the messagebus)
type ConsumerConfig struct {
	// Event is the event type to subscribe to
	Event protoreflect.ProtoMessage
	// Consumer is the callback function to execute when a message is received
	Consumer messagebus.Consumer
	// Duplicate indicates whether or not to duplicate messages to all replicas of the service. By default, this is false and messages are
	// load-balanced to a single replica of the service. If Duplicate is set to true, messages will instead be duplicated to each replica
	// of the service. This is useful in special cases where you need to ensure that all replicas of a service receive a message.
	Duplicate bool
	// IgnoreType states whether the subscribing consumer should ignore the event type when receiving messages. If true, the consumer
	// will receive all events that match Event.Source. If false, the consumer will only receive events that match both Event.Source and Event.Type.
	// Essentially this is a way to subscribe to all events of a certain aggregate type instead of a specific event type.
	IgnoreType bool
}

// CheckConfig specifies a configuration for use in readiness and wellness checks.
type CheckConfig struct {
	// Check is the check to perform.
	Check func(ctx *gin.Context)
}

// SecretsConfig defines connection configuration for the secrets vault to be used by the application.
type SecretsConfig struct {
	Client secrets.Client
	// Required should return true if the secrets client must be initialized successfully and false otherwise.
	// A common usecase for this is to only require the secrets client if the application is running in production mode.
	// For that usecase, you can use the following:
	//
	//  func(b chassis.MainBuilder) bool { return !b.IsDevMode() }
	Required func(b MainBuilder) bool
}

// MessageBusConfig defines the service options for the messagebus module
type MessageBusConfig struct {
	Buses []messagebus.Client
}

// InitializeConfig defines the beginning configuration of the application such
// as the root directory, logfile, etc.
type InitializeConfig struct {
	// BaseDirectory specifies the base directory from which to run the application.
	// This should only be used in local debugging.
	BaseDirectory string
	// LogFile specifies the file to print logs into.
	// This should only be used in local debugging.
	LogFile string
	// OnInitialize is called after the config and secrets have been initialized but prior to dao, service, and handler layers.
	OnInitialize func(b MainBuilder)
}

// MainBuilderConfig specifies each configuration for MainBuilder.
type MainBuilderConfig struct {
	// ApplicationName is the name of the application.
	ApplicationName string
	// Bool config variable that specifies whether the application is running in dev mode.
	// Defaults to "isDevMode" if not specified.
	IsDevModeVariable string
	// DoNotRegisterService specifies the service to register with the service registry.
	// Defaults to false if not specified. If you set this you must also set the server ports for both HTTP and RPC in the HandlerLayerConfig.
	//
	// ONLY SET THIS TO TRUE FOR THE REGISTRY SERVICE ITSELF OR IF YOU REALLY KNOW WHAT YOU'RE DOING.
	DoNotRegisterService bool
	// EventConfig is the event configuration.
	EventConfig *EventConfig
	// DaoLayerConfig is the dao layer configuration.
	DaoLayerConfig *DaoLayerConfig
	// DatabaseConfig is the `SQL` db configuatrion
	DatabaseConfig *DatabaseConfig
	// CacheConfig is the `Cache` configuatrion
	CacheConfig *CacheConfig
	// GatewayLayerConfig is the gateway layer configuration.
	GatewayLayerConfig *GatewayLayerConfig
	// ServiceLayerConfig is the service layer configuration.
	ServiceLayerConfig *ServiceLayerConfig
	// HandlerLayerConfig is the handler layer configuration.
	HandlerLayerConfig *HandlerLayerConfig
	// ReadinessCheckConfig is the readiness check configuration.
	ReadinessCheckConfig *CheckConfig
	// WellnessCheckConfig is the wellness check configuration.
	WellnessCheckConfig *CheckConfig
	// SecretsConfig is the secrets vault configuration.
	SecretsConfig *SecretsConfig
	// InitializeConfig is the configuration for initialize.
	InitializeConfig *InitializeConfig
	// MessageBusConfig is the configuration for connecting to the message bus.
	MessageBusConfig *MessageBusConfig
	// OnRun is called when MainBuilder.Run is called. DO NOT consume the calling thread.
	// It will be called after the configuration and logger have been initialized but before anything
	// else has been initialized.
	OnRun func(b MainBuilder)
	// OnStop is called when the application is closing.
	OnStop func(b MainBuilder)
}

type mainBuilder struct {
	logger l.Logger

	// basic configuration
	applicationName     string
	applicationDomain   string
	applicationVersion  string
	isDevMode           bool
	createEventerClient bool
	httpPort            string
	rpcPort             string

	// http/grpc
	httpServer     *http.Server
	httpRouter     *gin.Engine
	rpcServer      *grpc.Server
	rpcConnections []*grpc.ClientConn

	// clients/servers
	secretsClient  secrets.Client
	cacheClient    *redis.Client
	registryClient rgpb.RegistryClient
	eventManager   events.EventManager
	databases      []database.Client
	buses          []messagebus.Client
	tracer         trace.Tracer

	// functions
	onRun  func(b MainBuilder)
	onStop func(b MainBuilder)

	// configs
	consumerConfig       []ConsumerConfig
	readinessCheckConfig *CheckConfig
	wellnessCheckConfig  *CheckConfig
	appConfig            *viper.Viper
}

// PUBLIC METHODS

// NewMainBuilder initializes the whole microservice application. Pass in the configuration for each layer and
// then call Run() to start the application.
func NewMainBuilder(mbc *MainBuilderConfig) MainBuilder {
	// Create the base logger for the service
	logger := l.CreateLogger("info", mbc.ApplicationName)
	if mbc.InitializeConfig != nil && mbc.InitializeConfig.LogFile != "" {
		f, err := os.OpenFile(mbc.InitializeConfig.LogFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0755)
		if err != nil {
			logger.WithError(err).WithField("file", mbc.InitializeConfig.LogFile).Fatal("failed to initialize logger with file")
		}
		logrus.SetOutput(f)
	}

	createEventerClient := false
	if mbc.EventConfig != nil {
		createEventerClient = true
	}
	b := &mainBuilder{
		applicationName:      mbc.ApplicationName,
		logger:               logger,
		onRun:                mbc.OnRun,
		onStop:               mbc.OnStop,
		readinessCheckConfig: mbc.ReadinessCheckConfig,
		wellnessCheckConfig:  mbc.WellnessCheckConfig,
		createEventerClient:  createEventerClient,
	}

	// Get variables from local.yaml/values.yaml and environment variables for use by viper
	baseDir := "."
	if mbc.InitializeConfig != nil {
		if mbc.InitializeConfig.BaseDirectory != "" {
			baseDir = mbc.InitializeConfig.BaseDirectory
		}
	}

	// Setup application configuration
	b.setupConfig(mbc.InitializeConfig)

	// start tracer (if this value isn't set viper returns false. that way we default to starting the tracer)
	if !b.GetConfig().GetBool("disableTracer") {
		b.GetLogger().Info("starting tracer")
		uptrace.ConfigureOpentelemetry(
			uptrace.WithDSN(b.GetConfig().GetString("uptrace.dsn")),
			uptrace.WithServiceName(b.applicationName),
			uptrace.WithServiceVersion(b.applicationVersion),
		)
		b.tracer = otel.Tracer(b.applicationName)
	} else {
		// only warn if not running locally
		if !b.GetConfig().GetBool("isDevMode") {
			b.GetLogger().Warn("tracer is disabled. this should only be done when absolutely necessary (ie. memory leak)")
		}
	}

	// Add env and versions to logs
	logger.AddGlobalField("env", b.appConfig.GetString("namespace"))
	logger.AddGlobalField("version", b.appConfig.GetString("version"))

	// Determine what mode the application is running in
	b.setupMode(mbc.IsDevModeVariable)

	// Connect to secrets vault
	if mbc.SecretsConfig != nil {
		b.setupSecrets(mbc.SecretsConfig, baseDir)
	}

	// Call initialize if needed
	if mbc.InitializeConfig != nil && mbc.InitializeConfig.OnInitialize != nil {
		mbc.InitializeConfig.OnInitialize(b)
	}

	// Setup the message bus
	if mbc.MessageBusConfig != nil {
		b.setupMessageBus(mbc.MessageBusConfig)
	}

	// Setup the DAO layer
	if mbc.DaoLayerConfig != nil {
		// establish a connection to the SQL database if configured
		if mbc.DatabaseConfig != nil {
			b.setupDatabases(mbc.DatabaseConfig)
		}

		// establish a connection to the Redis cache if configured
		if mbc.CacheConfig != nil {
			b.setupCache(mbc.CacheConfig)
		}

		mbc.DaoLayerConfig.CreateDaoLayer(b)
	}

	// Setup the Gateway layer
	if mbc.GatewayLayerConfig != nil {
		if mbc.GatewayLayerConfig.CreateInternalClients != nil {
			clientFactory := cf.NewClientFactory()
			b.rpcConnections = mbc.GatewayLayerConfig.CreateInternalClients(b, clientFactory)
		}
	}

	// Setup the Service layer
	if mbc.ServiceLayerConfig != nil {
		mbc.ServiceLayerConfig.CreateServiceLayer(b)
	}

	// Register the service with the service registry
	if !mbc.DoNotRegisterService {
		b.setupRegistration(mbc.HandlerLayerConfig, mbc.ApplicationName)
	}

	// If the service is not supposed to be registered, set ports from config
	if mbc.DoNotRegisterService {
		b.rpcPort = b.appConfig.GetString(mbc.HandlerLayerConfig.GrpcPortConfigKey)
		b.httpPort = b.appConfig.GetString(mbc.HandlerLayerConfig.HttpPortConfigKey)
	}

	// Setup the Handler layer
	if mbc.HandlerLayerConfig != nil {
		b.setupHandlers(mbc.HandlerLayerConfig)
	}

	return b
}

// Run runs the microservice applications using the mainbuilder configuration.
// This means starting all the servers and connections and then listening for an exit condition.
func (b *mainBuilder) Run() {
	ctx := ct.New(context.Background(), b.GetTelemetry())
	if !b.GetConfig().GetBool("disableTracer") {
		defer uptrace.Shutdown(ctx)
	}

	// initialize eventer client if requested
	if b.createEventerClient {
		manager, err := events.NewEventManager(ctx, b.appConfig.GetString("registryServiceAddress"))
		if err != nil {
			b.GetLogger().WithError(err).Panic("failed to initialize event manager")
		}
		b.eventManager = manager
		defer b.eventManager.Close()
	}

	if b.onRun != nil {
		b.onRun(b)
	}

	if b.httpRouter != nil {
		go b.startHttpServer()
	}

	if b.rpcServer != nil {
		go b.startRpcServer()
	}

	// if running locally, initialize broker listeners without waiting for Shawarma call
	if b.isDevMode {
		b.subscribeConsumers(false)
	}

	signal.Notify(terminator.ApplicationChannel, os.Interrupt, syscall.SIGTERM)
	<-terminator.ApplicationChannel
	if b.onStop != nil {
		b.onStop(b)
	}
}

// Close closes all current connections and stops all active servers
func (b *mainBuilder) Close() {
	ctx := context.Background()

	// call service provided stop function
	if b.onStop != nil {
		b.onStop(b)
	}

	// close database connections
	for _, db := range b.databases {
		db.Disconnect(ctx)
	}

	// close all messagebus connections
	for _, bus := range b.buses {
		err := bus.Shutdown(false)
		if err != nil {
			b.logger.WithError(err).Error("failed to gracefully shutdown messagebus: forcing shutdown")
			bus.Shutdown(true)
		}
	}

	// close all rpc connections
	for _, conn := range b.rpcConnections {
		cf.CloseConnection(b.logger, conn)
	}

	// close servers
	b.stopHttpServer()
	b.stopRpcServer()

	b.logger.Info("service stopped")
}

// SETUP FUNCTIONS

func (b *mainBuilder) setupConfig(config *InitializeConfig) {

	// Get variables from local.yaml/values.yaml and environment variables for use by viper
	baseDir := "."
	if config != nil {
		if config.BaseDirectory != "" {
			baseDir = config.BaseDirectory
		}
	}

	// Setup viper
	err := ec.ReadEnvironmentConfigurations(b.logger, baseDir)
	if err != nil {
		b.logger.WithError(err).Fatal("failed to read environment configurations")
	}

	v := viper.GetViper()
	if v.IsSet("configMap") {
		v = v.Sub("configMap")
		v.AutomaticEnv()
	}
	b.appConfig = v

	b.applicationDomain = b.appConfig.GetString("domain")
	b.applicationVersion = b.appConfig.GetString("version")
}

func (b *mainBuilder) setupMode(isDevModeVariable string) {
	if isDevModeVariable == "" {
		isDevModeVariable = "isDevMode"
	}
	b.isDevMode = b.appConfig.GetBool(isDevModeVariable)
	if b.isDevMode {
		b.logger.Info("Currently running in dev mode")
		logrus.SetFormatter(&l.Formatter{
			ChildFormatter: &logrus.TextFormatter{
				ForceColors: true,
			},
		})
	} else {
		b.logger.Info("Currently running in prod mode")
	}
}

// setupSecrets loads the secrets vault configuration and overrides any defined values
// in the values configuration file with their secret values.
func (b *mainBuilder) setupSecrets(config *SecretsConfig, baseDir string) {
	ctx := context.Background()
	secretsVaultRequired := true
	if config.Required != nil {
		secretsVaultRequired = config.Required(b)
	}

	err := config.Client.Initialize(ctx, b.GetConfig())
	if err != nil {
		if secretsVaultRequired {
			b.logger.WithError(err).Fatal("failed to create required vault client")
		} else {
			b.logger.WithError(err).Warn("failed to create key vault client but not required")
			return
		}
	}
	b.secretsClient = config.Client

	type SecretOverride struct {
		ConfigKey string
		SecretKey string
	}
	overrides := []SecretOverride{}
	err = b.GetConfig().UnmarshalKey("secretsOverrides", &overrides)
	if err != nil {
		b.logger.WithError(err).Fatal("failed to read secrets overrides")
	}
	for _, o := range overrides {
		value, err := b.GetSecretsClient().Get(ctx, o.SecretKey)
		if err != nil || value == "" {
			b.logger.WithField("secret_key", o.SecretKey).WithError(err).Fatal("failed to read secret from vault")
		}
		b.GetConfig().Set(o.ConfigKey, value)
	}
}

func (b *mainBuilder) setupMessageBus(config *MessageBusConfig) {
	ctx := context.Background()
	for _, bus := range config.Buses {
		err := bus.Initialize(ctx, b.GetConfig())
		if err != nil {
			b.logger.WithError(err).Fatal("failed to initialize messagebus client")
		}
	}
	b.buses = config.Buses
}

func (b *mainBuilder) setupDatabases(config *DatabaseConfig) {
	ctx := context.Background()
	for _, c := range config.Databases {
		err := c.Initialize(ctx, b.GetConfig())
		if err != nil {
			b.logger.WithError(err).Fatal("failed to initialize database client")
		}
	}
	b.databases = config.Databases
}

func (b *mainBuilder) setupCache(config *CacheConfig) {
	options := &redis.Options{
		Addr:      b.GetConfig().GetString(config.CacheAddress),
		Password:  b.GetConfig().GetString(config.CacheSecret),
		TLSConfig: &tls.Config{MinVersion: tls.VersionTLS12},
	}
	b.cacheClient = redis.NewClient(options)
	go func() {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		err := b.cacheClient.Ping(ctx).Err()
		if err != nil {
			b.logger.WithField("redis_address", options.Addr).WithError(err).Fatal("failed to connect to the cache")
		}
	}()
}

func (b *mainBuilder) setupRegistration(config *HandlerLayerConfig, serviceName string) {
	clientFactory := cf.NewClientFactory()
	registryClient, conn, stdErr := clientFactory.CreateRegistryClient(b.GetLogger(), b.appConfig.GetString("registryServiceAddress"))
	defer conn.Close()
	if stdErr != nil {
		b.logger.WithError(stdErr).Fatal("failed to initialize registry client")
	}
	b.registryClient = registryClient

	// register service with registry
	c := context.Background()
	ctx, span := ct.New(c, b.GetTelemetry()).Span()
	defer span.End()
	response, stdErr := b.registryClient.Register(ctx, &rgpb.RegisterRequest{
		Name:    b.applicationName,
		Domain:  b.applicationDomain,
		Version: b.applicationVersion,
	})
	if stdErr != nil {
		b.logger.WithError(stdErr).Fatal("failed to register service with registry")
	}

	// if no handler layer config, nothing else to register
	if config == nil {
		return
	}

	// register grpc handlers
	if config.CreateRpcHandlers != nil {
		for _, s := range config.CreateRpcHandlers(b) {
			response, stdErr := b.registryClient.RegisterGrpcServer(ctx, &rgpb.RegisterGrpcServerRequest{
				Id:    response.Id,
				Route: s.Desc.ServiceName,
			})
			if stdErr != nil {
				b.logger.WithError(stdErr).Fatal("failed to register grpc handler with registry")
			}
			b.rpcPort = response.Port
		}
	}

	// TODO: register http handlers

	// TODO: register consumers
	// if config.CreateConsumers != nil {
	// 	consumers := make([]*rgpb.ConsumerRequest, len(b.consumerConfig))
	// 	for i, c := range b.consumerConfig {
	// 		e, _ := events.New(c.Event)
	// 		kind := rgpb.ConsumerKind_CONSUMER_KIND_ONE
	// 		if c.Duplicate {
	// 			kind = rgpb.ConsumerKind_CONSUMER_KIND_ALL
	// 		}
	// 		consumers[i] = &rgpb.ConsumerRequest{
	// 			Kind:        kind,
	// 			EventSource: e.Source,
	// 			EventType:   &e.Type,
	// 		}
	// 	}
	// 	_, stdErr = b.registryClient.RegisterConsumers(ctx, &rgpb.RegisterConsumersRequest{
	// 		Id:        response.Id,
	// 		Consumers: consumers,
	// 	})
	// 	if stdErr != nil {
	// 		b.logger.WithError(stdErr).Fatal("failed to register consumer with registry")
	// 	}
	// }
}

func (b *mainBuilder) setupHandlers(config *HandlerLayerConfig) {
	// ALWAYS create the http server since it hosts healthchecks, logging endpoints, etc.
	b.createHttpServer()

	// Setup the HTTP/Restful handlers only if configured
	if config.CreateRestHandlers != nil {
		config.CreateRestHandlers(b)
	}

	// Setup the RPC server and handlers only if configured
	if config.CreateRpcHandlers != nil {
		b.createRpcServer(config.RpcOptions...)
		for _, s := range config.CreateRpcHandlers(b) {
			b.GetRpcServer().RegisterService(&s.Desc, s.Handler)
		}
	}

	// Validate the consumer configuration only if configured
	if config.CreateConsumers != nil {
		for _, c := range b.consumerConfig {
			err := events.Validate(c.Event)
			if err != nil {
				b.logger.WithField("event", c.Event).WithError(err).Fatal("invalid event type in consumer config")
			}
		}
		b.consumerConfig = config.CreateConsumers(b)
	}
}

// PRIVATE METHODS

// subscribeConsumers subscribes all configured consumers to their respective event types on all configured messagebuses.
// This function should only be called when the shawarma sidecar makes a request to the HTTP handler saying that one of the service endpoints has been ASSIGNED to this pod.
// This way we guarantee that the service will only read messages that it is supposed to based off of it's Argo rollout status.
func (b *mainBuilder) subscribeConsumers(preview bool) {
	ctx := context.Background()
	for _, c := range b.consumerConfig {
		// we can ignore the error since validation has already happened
		e, _ := events.New(c.Event)
		params := messagebus.SubscribeParams{
			Event:      e,
			Consumer:   c.Consumer,
			IgnoreType: c.IgnoreType,
		}
		if !c.Duplicate {
			params.Group = fmt.Sprintf("%s.%s.%s", b.applicationDomain, b.applicationName, b.applicationVersion)
		}
		if preview {
			params.Tags = []string{"preview"}
		}
		for _, bus := range b.buses {
			err := bus.Subscribe(ctx, params)
			if err != nil {
				b.logger.WithError(err).WithField("event", e).Fatal("failed to subscribe consumer")
			}
			b.logger.WithFields(l.Fields{"event": e, "tags": params.Tags}).Info("subscribed consumer")
		}
	}
}

// unsubscribeConsumers unsubscribes all configured consumers from their respective event types on all configured messagebuses.
// This function should only be called when the shawarma sidecar makes a request to the HTTP handler saying that one of the service endpoints has been UNASSIGNED from this pod.
// This way we guarantee that the service will only read messages that it is supposed to based off of it's Argo rollout status.
func (b *mainBuilder) unsubscribeConsumers(preview bool) {
	ctx := context.Background()
	for _, c := range b.consumerConfig {
		// we can ignore the error since validation has already happened
		e, _ := events.New(c.Event)
		params := messagebus.UnsubscribeParams{
			Event:      e,
			IgnoreType: c.IgnoreType,
		}
		if preview {
			params.Tags = []string{"preview"}
		}
		for _, bus := range b.buses {
			err := bus.Unsubscribe(ctx, params)
			if err != nil {
				b.logger.WithError(err).WithField("event", e).Fatal("failed to unsubscribe consumer")
			}
			b.logger.WithFields(l.Fields{"event": e, "tags": params.Tags}).Info("unsubscribed consumer")
		}
	}
}

// GETTER FUNCTIONS
// --- Please keep in alphabetical order ---

func (b *mainBuilder) GetSecretsClient() secrets.Client {
	return b.secretsClient
}

func (b *mainBuilder) GetCacheClient() *redis.Client {
	return b.cacheClient
}

func (b *mainBuilder) GetConfig() *viper.Viper {
	return b.appConfig
}

func (b *mainBuilder) GetEventManager() events.EventManager {
	return b.eventManager
}

func (b *mainBuilder) GetHttpRouter() *gin.Engine {
	return b.httpRouter
}

func (b *mainBuilder) GetLogger() l.Logger {
	return b.logger
}

func (b *mainBuilder) GetRegistryClient() rgpb.RegistryClient {
	return b.registryClient
}

func (b *mainBuilder) GetRpcServer() *grpc.Server {
	return b.rpcServer
}

func (b *mainBuilder) GetTelemetry() ct.Telemetry {
	return ct.Telemetry{
		Logger: b.logger,
		Tracer: b.tracer,
	}
}

func (b *mainBuilder) IsDevMode() bool {
	return b.isDevMode
}
