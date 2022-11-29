package chassis

import (
	// standard libraries
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"net/url"
	"os"
	"os/signal"
	"path"
	"strings"
	"syscall"
	"time"

	// chassis packages
	runtime "github.com/banzaicloud/logrus-runtime-formatter"
	"github.com/jgkawell/galactus/pkg/chassis/broker"
	cf "github.com/jgkawell/galactus/pkg/chassis/clientfactory"
	ct "github.com/jgkawell/galactus/pkg/chassis/context"
	"github.com/jgkawell/galactus/pkg/chassis/db"
	ec "github.com/jgkawell/galactus/pkg/chassis/env"
	"github.com/jgkawell/galactus/pkg/chassis/events"
	messagebus "github.com/jgkawell/galactus/pkg/chassis/messagebus"
	"github.com/jgkawell/galactus/pkg/chassis/terminator"

	// other galactus modules
	azkeyvault "github.com/jgkawell/galactus/pkg/azkeyvault"
	l "github.com/jgkawell/galactus/pkg/logging/v2"

	// galactus api packages
	agpb "github.com/jgkawell/galactus/api/gen/go/core/aggregates/v1"
	rgpb "github.com/jgkawell/galactus/api/gen/go/core/registry/v1"
	evpb "github.com/jgkawell/galactus/api/gen/go/generic/events/v1"

	// third party libraries
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// MainBuilder is an interface that exposes the functionality for using the chassis module
type MainBuilder interface {

	// CONTROL FUNCTIONS

	// Run starts the execution of the application.
	Run()
	// Close releases all assets. Call via defer.
	Close()
	// Initialize and open GORM db session based on dialector.
	InitializeGORM(dbAddress string) (*gorm.DB, error)
	// StartHttpServer launches the http server. This will consume the calling thread.
	StartHttpServer()
	// Stop HttpServer shuts the http server down.
	StopHttpServer()
	// StartRpcServer launches the rpc server. This will consume the calling thread.
	StartRpcServer()
	// StopRpcServer shuts the rpc server down.
	StopRpcServer()

	// GETTER FUNCTIONS

	// IsDevMode specifies if this application is running in development mode (true) or production mode (false)
	IsDevMode() bool
	// GetCacheClient gets the cache implementation client.
	GetCacheClient() *redis.Client
	// GetConfig exposes the viper configuration for the application.
	GetConfig() *viper.Viper
	// GetLogger exposes the logger the application is using.
	GetLogger() l.Logger
	// GetMongoClient exposes the mongo client.
	GetMongoClient() *mongo.Client
	// GetHttpRouter exposes the gin http router.
	GetHttpRouter() *gin.Engine
	// GetBroker exposes the message bus.
	GetBroker() messagebus.MessageBus
	// GetRegistryClient exposes the registry client.
	GetRegistryClient() rgpb.RegistryClient
	// GetRpcServer exposes the grpc server.
	GetRpcServer() *grpc.Server
	// GetAzureKeyVaultClient exposes the azure key vault client.
	GetAzureKeyVaultClient() azkeyvault.KeyVaultClient
	// GetSqlClient exposes the sql client for services to use
	GetSqlClient() *gorm.DB
	// GetEventManager... TODO
	GetEventManager() events.EventManager
}

// DaoLayerConfig defines the function for initializing the DAO layer if any custom setup is required
type DaoLayerConfig struct {
	// CreateDaoLayer creates the dao layer using the mongo client.
	CreateDaoLayer func(b MainBuilder)
}

// NoSqlConfig defines the connection confiuration for the NoSQL database (only Mongo is currently supported)
type NoSqlConfig struct {
	// DbAddressVariable is the config variable containing the database's address.
	// NOTE: This address is reserved for mongo db only.
	DbAddressVariable string
}

// SqlConfig defines the connection configuration to the SQL database (only Postgres is currently supported)
type SqlConfig struct {
	// Sql configuration variables that are pulled either from a secret store, or deployment configration
	SqlDbHost   string
	SqlDbPort   string
	SqlDbName   string
	SqlDbUser   string
	SqlDbSecret string
	SqlDbSchema string
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

// HandlerLayerConfig specifies how the handler layer will be configured.
type HandlerLayerConfig struct {
	// HttpPortVariable will only be used if DoNotRegisterService is true. It specifies the config key to read the port from.
	// ONLY USE THIS FOR THE REGISTRY SERVICE ITSELF OR IF YOU REALLY KNOW WHAT YOU'RE DOING.
	HttpPortVariable string
	// CreateRestHandlers creates the http restful interface using the gin-engine router.
	CreateRestHandlers func(b MainBuilder)
	// RpcPortVariable will only be used if DoNotRegisterService is true. It specifies the config key to read the port from.
	// ONLY USE THIS FOR THE REGISTRY SERVICE ITSELF OR IF YOU REALLY KNOW WHAT YOU'RE DOING.
	RpcPortVariable string
	// CreateRpcHandlers creates the rpc interface using the grpc-server.
	CreateRpcHandlers func(b MainBuilder)
	// RpcOptions is a slice of optional grpc.ServerOption structs to set on the gRPC server.
	RpcOptions []grpc.ServerOption
	// CreateBrokerConfig defines a function that creates the broker configuration. This needs to
	// be a function called from main.go so that it has both the service interfaces (e.g. Handle())
	// and access to mainbuilder attributes (e.g. the logger).
	CreateBrokerConfig func(b MainBuilder) ConsumerConfig
}

// ConsumerConfig defines the configuration for consumers on the messagebus layer (e.g. RabbitMQ)
type ConsumerConfig struct {
	Configs []HandlerConfig
}

// HandlerConfig defines the configuration for a consumer handler (processing messages off of the messagebus)
type HandlerConfig struct {
	// AggregateType is mapped to the AMQP exchange name
	AggregateType evpb.AggregateType
	// EventType is mapped to the AMQP routing key of the queue
	EventType *evpb.EventType
	// EventCode ...
	EventCode interface{}
	// ConsumerKind is the type of consumer to create (e.g. queue, topic)
	ConsumerKind messagebus.ExchangeKind
	// Handler is the callback function to execute when a message is received
	Handler messagebus.ClientHandler
}

// CheckConfig specifies a configuration for use in readiness and wellness checks.
type CheckConfig struct {
	// Check is the check to perform.
	Check func(ctx *gin.Context)
}

// KeyVaultConfig defines connection configuration for the keyvault module
type KeyVaultConfig struct {
	// RequireKeyVault should return true if the key vault client must be initialized successfully and false otherwise.
	RequireKeyVault func(b MainBuilder) bool
	// KeyVaultResourceGroupVariable is the config variable for the key vault resource group.
	// One of KeyVaultResourceGroupVariable and GetKeyVaultResourceGroup must be specified.
	KeyVaultResourceGroupVariable string
	// GetKeyVaultResourceGroup retrieves the key vault resource group name.
	// One of KeyVaultResourceGroup and GetKeyVaultResourceGroup must be specified.
	GetKeyVaultResourceGroup func() string
	// KeyVaultNameVariable is the config variable for the key vault name.
	// One of KeyVaultNameVariable and GetKeyVaultName must be specified.
	KeyVaultNameVariable string
	// GetKeyVaultResourceGroup retrieves the key vault name.
	GetKeyVaultName func() string
	// KeyVaultOverridesVariable is the config variable for the key vault overrides mapping.
	KeyVaultOverridesVariable string
}

// MessageBusConfig defines the service options for the messagebus module
type MessageBusConfig struct {
	// MessageBusOptions specifies additional settings when creating the message bus.
	// MessageBusOptions []messagebus.ServiceOption
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
	// OnInitialize called when main build is being initialized.
	// This will be executed after config and key vault have been called but prior to dao, service, and handler layers.
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
	// CreateEventStoreClient specifies if the service needs a persistent client to the eventstore service to create events
	CreateEventStoreClient bool
	// DaoLayerConfig is the dao layer configuration.
	DaoLayerConfig *DaoLayerConfig
	// NoSqlConfig is the `NoSQL` db configuatrion
	NoSqlConfig *NoSqlConfig
	// SqlConfig is the `SQL` db configuatrion
	SqlConfig *SqlConfig
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
	// KeyVaultConfig is the key vault configuration.
	KeyVaultConfig *KeyVaultConfig
	// InitializeConfig is the configuration for initialize.
	InitializeConfig *InitializeConfig
	// MessageBusConfig is the configuration for connecting to the message bus.
	MessageBusConfig *MessageBusConfig
	// OnRun is called when MainBuilder.Run is called. DO NOT consume the calling thread.
	OnRun func(b MainBuilder)
	// OnStop is called when the application is closing.
	OnStop func(b MainBuilder)
}

type mainBuilder struct {
	logger l.Logger

	// basic configuration
	applicationName        string
	isDevMode              bool
	createEventStoreClient bool
	httpPort               string
	rpcPort                string

	// http/grpc
	httpServer     *http.Server
	httpRouter     *gin.Engine
	rpcServer      *grpc.Server
	rpcConnections []*grpc.ClientConn

	// clients/servers
	azureKeyVaultClient azkeyvault.KeyVaultClient
	bus                 messagebus.MessageBus
	cacheClient         *redis.Client
	noSqlClient         *mongo.Client
	sqlClient           *gorm.DB
	registryClient      rgpb.RegistryClient
	eventManager        *events.Manager

	// functions
	onRun  func(b MainBuilder)
	onStop func(b MainBuilder)

	// configs
	consumerConfig       ConsumerConfig
	brokerConfigs        []broker.BrokerDefinition
	readinessCheckConfig *CheckConfig
	wellnessCheckConfig  *CheckConfig
	viper                *viper.Viper
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

	b := &mainBuilder{
		applicationName:        mbc.ApplicationName,
		logger:                 logger,
		onRun:                  mbc.OnRun,
		onStop:                 mbc.OnStop,
		readinessCheckConfig:   mbc.ReadinessCheckConfig,
		wellnessCheckConfig:    mbc.WellnessCheckConfig,
		createEventStoreClient: mbc.CreateEventStoreClient,
		eventManager:           &events.Manager{},
	}

	// Get variables from local.yaml/values.yaml and environment variables for use by viper
	baseDir := "."
	if mbc.InitializeConfig != nil {
		if mbc.InitializeConfig.BaseDirectory != "" {
			baseDir = mbc.InitializeConfig.BaseDirectory
		}
	}

	// Setup viper
	err := ec.ReadEnvironmentConfigurations(logger, baseDir)
	if err != nil {
		logger.WithError(err).Fatal("failed to read environment configurations")
	}

	v := viper.GetViper()
	if v.IsSet("configMap") {
		v = v.Sub("configMap")
		v.AutomaticEnv()
	}
	b.viper = v
	b.viper.SetDefault("namespace", "local")
	b.viper.SetDefault("version", "0.0.0")

	// Add env and versions to logs
	logger.AddGlobalField("env", b.viper.GetString("namespace"))
	logger.AddGlobalField("version", b.viper.GetString("version"))

	// Determine what mode the application is running in
	devModeKey := mbc.IsDevModeVariable
	if devModeKey == "" {
		devModeKey = "isDevMode"
	}
	b.isDevMode = b.viper.GetBool(devModeKey)
	if b.isDevMode {
		b.logger.Info("Currently running in dev mode")
		logrus.SetFormatter(&runtime.Formatter{
			ChildFormatter: &logrus.TextFormatter{
				ForceColors: true,
			},
		})
	} else {
		b.logger.Info("Currently running in prod mode")
	}

	// Connect to key vault (only when remote/not dev mode)
	if mbc.KeyVaultConfig != nil && !b.isDevMode {
		b.loadKeyVault(mbc.KeyVaultConfig, baseDir)
	}

	// Call initialize if needed
	if mbc.InitializeConfig != nil && mbc.InitializeConfig.OnInitialize != nil {
		mbc.InitializeConfig.OnInitialize(b)
	}

	// If the service needs to be registered, register the service
	if !mbc.DoNotRegisterService {
		clientFactory := cf.NewClientFactory(b.logger)
		registryClient, registryConnection, err := clientFactory.CreateRegistryClient(b.GetLogger(), b.viper.GetString("registryServiceAddress"))
		b.registryClient = registryClient
		if err != nil {
			b.logger.WithError(err).Fatal("failed to create registry client")
		}
		defer clientFactory.CloseConnection(b.GetLogger(), registryConnection)
	}

	// Setup the message bus
	if mbc.MessageBusConfig != nil {
		var (
			user     = b.viper.GetString("MessageBusUser")
			password = b.viper.GetString("MessageBusPassword")
			ips      = b.viper.GetString("MessageBusIPs")
		)
		b.bus = messagebus.NewMessageBus(
			b.GetLogger(),
			user,
			password,
			ips,
		)

		if err = b.bus.Connect(context.Background(), b.GetLogger()); err != nil {
			logger.WithError(err).Fatal("failed to connect to message bus")
		}
	}

	// Setup the DAO layer
	if mbc.DaoLayerConfig != nil {
		// establish a connection to the NoSQL database if configured
		if mbc.NoSqlConfig != nil {
			dbAddress := b.viper.GetString(mbc.NoSqlConfig.DbAddressVariable)
			b.noSqlClient, err = db.CreateNoSqlClient(logger, dbAddress)
			if err != nil {
				logger.WithField("db_address", dbAddress).WithError(err).Fatal("failed to create database client")
			}
			u, err := url.Parse(dbAddress)
			if err != nil {
				logger.WithField("db_address", dbAddress).WithError(err).Fatal("failed to parse database address")
			}
			if u.User != nil {
				// remove the password so it doesn't get logged
				u.User = url.User(u.User.Username())
			}

			logger.Info("Connected to " + u.String())
		}

		// establish a connection to the SQL database if configured
		if mbc.SqlConfig != nil {
			b.sqlClient, err = db.CreateSqlClient(
				logger,
				b.viper.GetString(mbc.SqlConfig.SqlDbUser),
				b.viper.GetString(mbc.SqlConfig.SqlDbSecret),
				b.viper.GetString(mbc.SqlConfig.SqlDbHost),
				b.viper.GetString(mbc.SqlConfig.SqlDbPort),
				b.viper.GetString(mbc.SqlConfig.SqlDbName),
				b.viper.GetString(mbc.SqlConfig.SqlDbSchema),
				b.isDevMode,
			)
			if err != nil {
				logger.WithError(err).Fatal("failed to create sql database client")
			}

			logger.Info("connected to sql db successfully")
		}

		// establish a connection to the Redis cache if configured
		if mbc.CacheConfig != nil {
			options := &redis.Options{
				Addr:      b.GetConfig().GetString(mbc.CacheConfig.CacheAddress),
				Password:  b.GetConfig().GetString(mbc.CacheConfig.CacheSecret),
				TLSConfig: &tls.Config{MinVersion: tls.VersionTLS12},
			}
			b.cacheClient = redis.NewClient(options)
			go func() {
				ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
				defer cancel()
				err := b.cacheClient.Ping(ctx).Err()
				if err != nil {
					logger.WithField("redis_address", options.Addr).
						WithError(err).
						Panic("could not reach the cache")
				}
			}()
		}

		mbc.DaoLayerConfig.CreateDaoLayer(b)
	}

	// Setup the Gateway layer
	if mbc.GatewayLayerConfig != nil {
		if mbc.GatewayLayerConfig.CreateInternalClients != nil {
			clientFactory := cf.NewClientFactory(logger)
			b.rpcConnections = mbc.GatewayLayerConfig.CreateInternalClients(b, clientFactory)
		}
	}

	// Setup the Service layer
	if mbc.ServiceLayerConfig != nil {
		mbc.ServiceLayerConfig.CreateServiceLayer(b)
	}

	// Setup the broker configuration
	if mbc.HandlerLayerConfig.CreateBrokerConfig != nil {
		b.consumerConfig = mbc.HandlerLayerConfig.CreateBrokerConfig(b)
		for _, c := range b.consumerConfig.Configs {
			err := events.ValidateEventCode(c.EventCode)
			if err != nil {
				b.logger.WithField("event_code", c.EventCode).WithError(err).Fatal("invalid event code in consumer config")
			}
		}
	}

	// register the service with the service registry
	if !mbc.DoNotRegisterService {
		// set protocols based on the chassis configuration
		protocols := []*rgpb.ProtocolRequest{}
		// NOTE: all services have an http handler that handles health checks. this handler can optionally be extended for other uses.
		protocols = append(protocols, &rgpb.ProtocolRequest{
			Order: 0,
			Kind:  agpb.ProtocolKind_PROTOCOL_KIND_HTTP,
		})
		// add rpc handlers only if requested
		if mbc.HandlerLayerConfig.CreateRpcHandlers != nil {
			protocols = append(protocols, &rgpb.ProtocolRequest{
				Order: 1,
				Kind:  agpb.ProtocolKind_PROTOCOL_KIND_GRPC,
			})
		}

		// pull together producers, and consumers from chassis config
		consumers := []*rgpb.ConsumerRequest{}
		for i, c := range b.consumerConfig.Configs {

			// map the messagebus kind to the consumer kind
			var protoKind agpb.ConsumerKind
			switch c.ConsumerKind {
			case messagebus.ExchangeKindTopic:
				protoKind = agpb.ConsumerKind_CONSUMER_KIND_TOPIC
			case messagebus.ExchangeKindDirect:
				protoKind = agpb.ConsumerKind_CONSUMER_KIND_QUEUE
			default:
				b.logger.WithField("kind", c.ConsumerKind).Fatal("unsupported messagebus kind")
			}

			consumers = append(consumers, &rgpb.ConsumerRequest{
				Order:         int32(i),
				Kind:          protoKind,
				AggregateType: events.AggregateTypeAsString(c.AggregateType),
				EventType:     events.EventTypeAsString(c.EventType),
				EventCode:     events.EventCodeAsString(c.EventCode),
			})
		}

		var err error
		ctx, err := ct.NewExecutionContext(context.Background(), b.logger, uuid.NewString())
		if err != nil {
			b.logger.WithError(err).Fatal("failed to create execution context for call to registry")
			return nil
		}
		registerResponse, err := b.registryClient.Register(ctx.GetContext(), &rgpb.RegisterRequest{
			Name:        mbc.ApplicationName,
			Domain:      "", // TODO: how do we manage this?
			Version:     b.viper.GetString("version"),
			Description: b.viper.GetString("description"),
			Protocols:   protocols,
			Consumers:   consumers,
		})
		if err != nil {
			b.logger.WithError(err).Fatal("failed to register the service")
		}

		// set server ports based on registration values
		for _, p := range registerResponse.GetProtocols() {
			if p.GetKind() == agpb.ProtocolKind_PROTOCOL_KIND_GRPC {
				b.rpcPort = fmt.Sprint(p.GetPort())
			}
			if p.GetKind() == agpb.ProtocolKind_PROTOCOL_KIND_HTTP {
				b.httpPort = fmt.Sprint(p.GetPort())
			}
		}

		// set consumer configs based on registration values
		b.brokerConfigs = make([]broker.BrokerDefinition, len(registerResponse.GetConsumers()))
		for _, c := range registerResponse.GetConsumers() {
			b.brokerConfigs[c.GetOrder()] = broker.BrokerDefinition{
				RoutingKey: c.GetRoutingKey(),
				Exchange:   c.GetExchange(),
				QueueName:  c.GetQueueName(),
				Handler:    b.consumerConfig.Configs[c.GetOrder()].Handler,
			}
		}
	}

	// if the service is not supposed to be registered, set ports from config
	if mbc.DoNotRegisterService {
		b.rpcPort = b.viper.GetString(mbc.HandlerLayerConfig.RpcPortVariable)
		b.httpPort = b.viper.GetString(mbc.HandlerLayerConfig.HttpPortVariable)
	}

	// Setup the Handler layer
	if mbc.HandlerLayerConfig != nil {
		// ALWAYS create the http server since it hosts healthchecks, logging endpoints, etc.
		b.createHttpServer()

		// Setup the HTTP/Restful handler only if it's configured
		if mbc.HandlerLayerConfig.CreateRestHandlers != nil {
			mbc.HandlerLayerConfig.CreateRestHandlers(b)
		}

		// Setup the RPC server and handler only if it's configured
		if mbc.HandlerLayerConfig.CreateRpcHandlers != nil {
			b.createRpcServer(mbc.HandlerLayerConfig.RpcOptions...)
			mbc.HandlerLayerConfig.CreateRpcHandlers(b)
		}
	}

	return b
}

// Run runs the microservice applications using the mainbuilder configuration.
// This means starting all the servers and connections and then listening for an exit condition.
func (b *mainBuilder) Run() {
	ctx, _ := ct.NewExecutionContext(context.Background(), b.GetLogger(), uuid.NewString())
	// start datadog tracer (if this value isn't set viper returns false. that way we default to starting the tracer)
	if !b.GetConfig().GetBool("disableTracer") {
		b.GetLogger().Info("starting tracer")
		tracer.Start()
		defer tracer.Stop()
	} else {
		// only warn if not running locally
		if !b.GetConfig().GetBool("isDevMode") {
			b.GetLogger().Warn("tracer is disabled. this should only be done when absolutely necessary (ie. memory leak)")
		}
	}

	if b.onRun != nil {
		b.onRun(b)
	}

	if b.httpRouter != nil {
		go b.StartHttpServer()
	}

	if b.rpcServer != nil {
		go b.StartRpcServer()
	}

	// initialize eventstore client if requested
	if b.createEventStoreClient {
		clientFactory := cf.NewClientFactory(b.logger)
		// TODO: is there a better way to persist clients so we don't have to create them multiple times?
		registryClient, conn, stdErr := clientFactory.CreateRegistryClient(b.GetLogger(), b.viper.GetString("registryServiceAddress"))
		defer conn.Close()
		if stdErr != nil {
			b.logger.WithError(stdErr).Fatal("failed to initialize registry client")
		}
		b.registryClient = registryClient

		resp, stdErr := b.registryClient.Connection(ctx.GetContext(), &rgpb.ConnectionRequest{
			Name:    "eventstore",
			Version: "0.0.0",
			Type:    agpb.ProtocolKind_PROTOCOL_KIND_GRPC,
		})
		if stdErr != nil {
			b.logger.WithError(stdErr).Fatal("failed to request connection info from registry service during initialization of eventstore client")
		}
		client, conn, stdErr := clientFactory.CreateEventStoreClient(b.GetLogger(), resp.GetAddress())
		defer conn.Close()
		if stdErr != nil {
			b.logger.WithError(stdErr).Fatal("failed to initialize eventstore client")
		}
		b.eventManager.Client = client
	}

	// if running locally, initialize broker listeners without waiting for Shawarma call
	if b.isDevMode {
		b.initializeBrokerListeners(b.applicationName)
	}

	signal.Notify(terminator.ApplicationChannel, os.Interrupt, syscall.SIGTERM)
	<-terminator.ApplicationChannel
	if b.onStop != nil {
		b.onStop(b)
	}
}

// Close closes all current connections and stops all active servers
func (b *mainBuilder) Close() {
	if b.onStop != nil {
		b.onStop(b)
	}
	if b.noSqlClient != nil {
		db.DisconnectNoSqlClient(b.logger, b.noSqlClient)
	}
	if b.sqlClient != nil {
		db.DisconnectSqlClient(b.logger, b.sqlClient)
	}
	for _, conn := range b.rpcConnections {
		cf.CloseConnection(b.logger, conn)
	}
	b.StopHttpServer()
	b.StopRpcServer()
	b.logger.Info("service stopped")
}

// InitializeGORM sets up the GORM SQL DB connection
func (b *mainBuilder) InitializeGORM(dbAddress string) (*gorm.DB, error) {
	return gorm.Open(postgres.Open(dbAddress), &gorm.Config{FullSaveAssociations: true})
}

// PRIVATE METHODS

// initializeBrokerListeners registers the topic and queue listeners with the messagebus according to the service configuration and preview/no preview status.
// This function should only be called when the shawarma sidecar makes a request to the HTTP handler saying that one of the service endpoints has been ASSIGNED to this pod.
// This way we guarantee that the service will only read messages that it is supposed to based off of it's Argo rollout status.
func (b *mainBuilder) initializeBrokerListeners(serviceName string) {
	// setup queues
	if len(b.brokerConfigs) > 0 {
		for _, bd := range b.brokerConfigs {
			bd.RoutingKey = modifyRoutingKey(bd.RoutingKey, serviceName)
			broker.RegisterConsumer(b.logger, b.bus, bd)
		}
	}
}

// cancelBrokerListeners cancels (disconnects) the topic and queue listeners with the messagebus according to the service configuration and preview/no preview status.
// This function should only be called when the shawarma sidecar makes a request to the HTTP handler saying that one of the service endpoints has been UNASSIGNED from this pod.
// This way we guarantee that the service will only read messages that it is supposed to based off of it's Argo rollout status.
func (b *mainBuilder) cancelBrokerListeners(serviceName string) {
	// cancel queues
	if len(b.brokerConfigs) > 0 {
		for _, bd := range b.brokerConfigs {
			bd.RoutingKey = modifyRoutingKey(bd.RoutingKey, serviceName)
			b.bus.CancelConsumerChannelsBySuffix(bd.RoutingKey, false)
		}
	}
}

// modifyRoutingKey takes in a routing key and the service name (shape: SERVICE or SERVICE-preview)
// and appends the `-preview` suffix if the service is in preview mode. This controls argo blue/green
// routing for messagebus traffic in the same way as Istio for network traffic.
func modifyRoutingKey(routingKey, serviceName string) string {
	if strings.Contains(serviceName, "preview") {
		routingKey = fmt.Sprintf("%s-%s", routingKey, "preview")
	}
	return routingKey
}

// loadKeyVault loads the keyvault configuration and overrides any defined values
// in the values configuration file with their keyvault values.
func (b *mainBuilder) loadKeyVault(config *KeyVaultConfig, baseDir string) {
	var (
		err           error
		resourceGroup string
		keyVaultName  string
	)

	if config.GetKeyVaultResourceGroup != nil {
		resourceGroup = config.GetKeyVaultResourceGroup()
	} else {
		resourceGroup = b.viper.GetString(config.KeyVaultResourceGroupVariable)
	}

	if config.GetKeyVaultName != nil {
		keyVaultName = config.GetKeyVaultName()
	} else {
		keyVaultName = b.viper.GetString(config.KeyVaultNameVariable)
	}

	keyVaultRequired := true
	if config.RequireKeyVault != nil {
		keyVaultRequired = config.RequireKeyVault(b)
	}

	if len(resourceGroup) == 0 {
		msg := "GetKeyVaultResourceGroup or KeyVaultResourceGroup is required"
		if keyVaultRequired {
			b.logger.Fatal(msg)
		}
		b.logger.Warn(msg)
		return
	}
	if len(keyVaultName) == 0 {
		msg := "GetKeyVaultName or KeyVaultName is required"
		if keyVaultRequired {
			b.logger.Fatal(msg)
		}
		b.logger.Warn(msg)
		return
	}

	subFile := "/etc/kubernetes/azure.json"
	if b.azureKeyVaultClient, err = azkeyvault.NewClientConfigPath(b.logger, subFile, resourceGroup, keyVaultName); err != nil {
		subFile = path.Join(baseDir, "azure.json")
		if b.azureKeyVaultClient, err = azkeyvault.NewClientConfigPath(b.logger, subFile, resourceGroup, keyVaultName); err != nil {
			if keyVaultRequired {
				b.logger.WithError(err).Fatal("failed to create key vault client")
			} else {
				b.logger.WithError(err).Warn("failed to create key vault client")
			}
		}
	}

	if config.KeyVaultOverridesVariable != "" {
		overrides := b.viper.GetStringMapString(config.KeyVaultOverridesVariable)
		kvc := b.GetAzureKeyVaultClient()
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		for configKey, keyVaultKey := range overrides {
			keyVaultValue, err := kvc.GetKeyVaultSecret(ctx, b.GetLogger(), keyVaultKey)
			if err != nil {
				b.GetLogger().WithError(err).Fatal("failed to get key vault value")
			}
			b.viper.Set(configKey, keyVaultValue)
		}
	}
}

// GETTER FUNCTIONS
// --- Please keep in alphabetical order ---

func (b *mainBuilder) GetAzureKeyVaultClient() azkeyvault.KeyVaultClient {
	return b.azureKeyVaultClient
}

func (b *mainBuilder) GetBroker() messagebus.MessageBus {
	return b.bus
}

func (b *mainBuilder) GetCacheClient() *redis.Client {
	return b.cacheClient
}

func (b *mainBuilder) GetConfig() *viper.Viper {
	return b.viper
}

func (b *mainBuilder) GetHttpRouter() *gin.Engine {
	return b.httpRouter
}

func (b *mainBuilder) GetLogger() l.Logger {
	return b.logger
}

func (b *mainBuilder) GetMongoClient() *mongo.Client {
	return b.noSqlClient
}

func (b *mainBuilder) GetRegistryClient() rgpb.RegistryClient {
	return b.registryClient
}

func (b *mainBuilder) GetRpcServer() *grpc.Server {
	return b.rpcServer
}

func (b *mainBuilder) GetSqlClient() *gorm.DB {
	return b.sqlClient
}

func (b *mainBuilder) IsDevMode() bool {
	return b.isDevMode
}

func (b *mainBuilder) GetEventManager() events.EventManager {
	return b.eventManager
}
