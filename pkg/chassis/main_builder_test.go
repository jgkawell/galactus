package chassis

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"net/http/httptest"
	"reflect"
	"syscall"
	"testing"

	"github.com/circadence-official/galactus/pkg/azkeyvault"
	cf "github.com/circadence-official/galactus/pkg/chassis/clientfactory"
	"github.com/circadence-official/galactus/pkg/chassis/db"
	ec "github.com/circadence-official/galactus/pkg/chassis/env"
	"github.com/circadence-official/galactus/pkg/chassis/messagebus"
	"github.com/circadence-official/galactus/pkg/chassis/terminator"
	l "github.com/circadence-official/galactus/pkg/logging/v2"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	v "github.com/spf13/viper"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/undefinedlabs/go-mpatch"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

var mbTest = MainBuilderConfig{
	InitializeConfig: &InitializeConfig{LogFile: "", BaseDirectory: ".", OnInitialize: func(b MainBuilder) {}},
	MessageBusConfig: &MessageBusConfig{},
	DaoLayerConfig:   &DaoLayerConfig{CreateDaoLayer: func(b MainBuilder) {}},
	NoSqlConfig:      &NoSqlConfig{},
	SqlConfig:        &SqlConfig{},
	KeyVaultConfig: &KeyVaultConfig{
		RequireKeyVault:               func(b MainBuilder) bool { return false },
		KeyVaultResourceGroupVariable: "",
		GetKeyVaultResourceGroup:      func() string { return "test" },
		KeyVaultNameVariable:          "",
		GetKeyVaultName:               func() string { return "test" },
		KeyVaultOverridesVariable:     "",
	},
	GatewayLayerConfig: &GatewayLayerConfig{
		CreateInternalClients: func(b MainBuilder, clientFactory cf.ClientFactory) []*grpc.ClientConn { return nil },
	},
	CacheConfig:        &CacheConfig{},
	ServiceLayerConfig: &ServiceLayerConfig{CreateServiceLayer: func(b MainBuilder) {}},
	HandlerLayerConfig: &HandlerLayerConfig{
		CreateRestHandlers: func(b MainBuilder) {},
		CreateRpcHandlers:  func(b MainBuilder) {},
		RpcOptions:         nil,
		CreateBrokerConfig: func(b MainBuilder) ConsumerConfig { return ConsumerConfig{} },
	},
}

type mbMock struct {
	mock.Mock
}

func (mb *mbMock) Connect(c messagebus.MessageBus, ctx context.Context, connections ...string) error {
	args := mb.Called(ctx, connections)
	return args.Error(0)
}

func (mb *mbMock) ReadEnvironmentConfigurationsMock(logger l.Logger, dir string) error {
	args := mb.Called(logger, dir)
	return args.Error(0)
}

func (mb *mbMock) NewMessageBusMock(svc, namespace string,
	logger l.Logger, mgmtURL string) (messagebus.MessageBus, error) {
	args := mb.Called(svc, namespace, logger, mgmtURL)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(messagebus.MessageBus), args.Error(1)
}

func (mb *mbMock) CreateClientMock(logger l.Logger, dbAddress string) (*mongo.Client, error) {
	args := mb.Called(logger, dbAddress)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*mongo.Client), args.Error(1)
}

func (mb *mbMock) CreateSQLClientMock(logger l.Logger, user, secret, host, port, name, schema string, mode bool) (*gorm.DB, error) {
	args := mb.Called(logger, user, secret, host, port, name, schema, mode)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*gorm.DB), args.Error(1)
}

func Test_NewMainBuilder(t *testing.T) {
	var (
		mb        mbMock
		busMock   mbMock
		mongoMock mbMock
		sqlMock   mbMock
	)
	h, err := mpatch.PatchMethod(ec.ReadEnvironmentConfigurations, mb.ReadEnvironmentConfigurationsMock)
	require.NoError(t, err)
	defer h.Unpatch()

	h2, err := mpatch.PatchMethod(messagebus.NewMessageBus, busMock.NewMessageBusMock)
	require.NoError(t, err)
	defer h2.Unpatch()

	h3, err := mpatch.PatchMethod(db.CreateNoSqlClient, mongoMock.CreateClientMock)
	require.NoError(t, err)
	defer h3.Unpatch()

	h4, err := mpatch.PatchMethod(db.CreateSqlClient, sqlMock.CreateSQLClientMock)
	require.NoError(t, err)
	defer h4.Unpatch()

	testCases := []struct {
		testName, appName string
		mainbuilder       MainBuilderConfig
		readEnvErr, mBusErr,
		connectErr, mongoErr, sqlErr error
		mbIps, namespace, dir, dbVar string
		panic                        bool
	}{
		{testName: "Everything is OK", mainbuilder: mbTest, namespace: "#1", mbIps: "localhost:5670", appName: "#1",
			dir: "./env1", dbVar: "db"},

		{testName: "panic on ReadEnvironmentConfigurations", mainbuilder: mbTest, dir: "./env2", appName: "#2",
			namespace:  "#2",
			readEnvErr: errors.New("panic"), panic: true},

		{testName: "panic on NewMessageBus", mainbuilder: mbTest, dir: "./env3", appName: "#3", namespace: "#3",
			mBusErr: errors.New("panic"), panic: true, mbIps: "localhost:5671"},

		{testName: "panic on connect to mBus", mainbuilder: mbTest, dir: "./env4", appName: "#4", namespace: "#4",
			connectErr: errors.New("panic"), panic: true, mbIps: "localhost:5672"},

		{testName: "panic on MongoClient creation", mainbuilder: mbTest, dir: "./env5", appName: "#5", namespace: "#5",
			mongoErr: errors.New("panic"), panic: true, mbIps: "localhost:5673"},

		{testName: "panic on SqlClient creation", mainbuilder: mbTest, dir: "./env6", appName: "#6", namespace: "#6",
			sqlErr: errors.New("panic"), panic: true, mbIps: "localhost:5673"},
	}
	for _, tt := range testCases {
		t.Run(tt.testName, func(t *testing.T) {
			connectionAMQP := fmt.Sprintf("amqp://guest:guest@%v", tt.mbIps)
			connectionHTTP := fmt.Sprintf("http://guest:guest@%v", tt.mbIps)
			tt.mainbuilder.InitializeConfig.LogFile = ""
			tt.mainbuilder.ApplicationName = tt.appName
			tt.mainbuilder.InitializeConfig.BaseDirectory = tt.dir
			tt.mainbuilder.NoSqlConfig.DbAddressVariable = tt.dbVar
			v.Set("MessageBusUser", "guest")
			v.Set("MessageBusPassword", "guest")
			v.Set("MessageBusIPs", tt.mbIps)
			v.Set("namespace", tt.namespace)
			v.Set(tt.dbVar, connectionHTTP)
			v.Set("isDevMode", "true")

			b := &messagebus.MockMessageBus{}
			// mock error on empty user string
			sqlMock.On("CreateSQLClientMock", mock.Anything, "", mock.Anything,
				mock.Anything, mock.Anything, mock.Anything, mock.Anything, true).Return(&gorm.DB{}, tt.sqlErr).Once()

			mongoMock.On("CreateClientMock", mock.Anything, connectionHTTP).Return(&mongo.Client{}, tt.mongoErr).Once()

			b.On("Connect", context.TODO(), connectionAMQP).Return(tt.connectErr).Once()

			busMock.On("NewMessageBusMock", tt.appName, tt.namespace, mock.Anything, connectionHTTP,
				mock.Anything).Return(b, tt.mBusErr).Once()

			mb.On("ReadEnvironmentConfigurationsMock", mock.Anything, tt.dir).Return(tt.readEnvErr).Once()
			if tt.panic != false {
				require.Panics(t, func() {
					NewMainBuilder(&tt.mainbuilder)
				})
			} else {
				require.NotPanics(t, func() {
					NewMainBuilder(&tt.mainbuilder)
				})
			}
		})
		v.Reset()
	}
}

func Test_Getters(t *testing.T) {
	logger, _ := l.CreateNullLogger()
	b := mainBuilder{
		cacheClient:         &redis.Client{},
		logger:              logger,
		bus:                 &messagebus.MockMessageBus{},
		isDevMode:           false,
		noSqlClient:         &mongo.Client{},
		httpPort:            "8080",
		httpServer:          &http.Server{},
		httpRouter:          &gin.Engine{},
		rpcPort:             "8091",
		rpcServer:           &grpc.Server{},
		rpcConnections:      []*grpc.ClientConn{},
		sqlClient:           &gorm.DB{},
		azureKeyVaultClient: &azkeyvault.MockKeyVaultClient{},
		viper:               v.New(),
	}
	require.NotNil(t, b.GetConfig())
	require.NotNil(t, b.GetCacheClient())
	require.NotNil(t, b.GetLogger())
	require.NotNil(t, b.GetMongoClient())
	require.NotNil(t, b.GetHttpRouter())
	require.NotNil(t, b.GetBroker())
	require.NotNil(t, b.GetRpcServer())
	require.NotNil(t, b.GetAzureKeyVaultClient())
	require.NotNil(t, b.GetAzureKeyVaultClient())
	require.NotNil(t, b.GetSqlClient())
	require.NotNil(t, b.IsDevMode())
}

func Test_loadKeyVault(t *testing.T) {
	testCases := []struct {
		testName, rg, name, over string
		required, panic          bool
		getName, getRg           func() string
		rgVar, nameVar, overVar  string
	}{
		{testName: "panic, rg = 0, required", panic: true, required: true, nameVar: "#1", overVar: "#1",
			getName: func() string { return "#1" }, name: "#1", over: "#1"},

		{testName: "warning rg = 0, not required", panic: false, required: false, nameVar: "#2", overVar: "#2",
			getName: func() string { return "#2" }, name: "#2", over: "#2"},

		{testName: "panic, kv = 0, required", panic: true, required: true, rgVar: "#3", overVar: "#3",
			getRg: func() string { return "#3" }, rg: "#3", over: "#3"},

		{testName: "warning kv = 0, not required", panic: false, required: false, rgVar: "#4", overVar: "#4",
			getRg: func() string { return "#4" }, rg: "#4", over: "#4"},

		{testName: "warning create kv client", panic: false, required: false, nameVar: "5", rgVar: "#5", overVar: "#5",
			name: "#5", rg: "#5", over: "#5"},

		{testName: "warning create kv client", panic: true, required: true, nameVar: "6", rgVar: "#6", overVar: "#6",
			name: "#6", rg: "#6", over: "#6"},
	}
	for _, tt := range testCases {
		logger, hook := l.CreateNullLogger()
		t.Run(tt.testName, func(t *testing.T) {
			viper := v.New()
			viper.Set(tt.rgVar, tt.rg)
			viper.Set(tt.nameVar, tt.name)
			viper.Set(tt.overVar, tt.over)
			b := mainBuilder{logger: logger, azureKeyVaultClient: nil, viper: viper}
			config := &KeyVaultConfig{
				RequireKeyVault:               func(b MainBuilder) bool { return tt.required },
				KeyVaultResourceGroupVariable: tt.rgVar,
				GetKeyVaultResourceGroup:      tt.getRg,
				KeyVaultNameVariable:          tt.nameVar,
				GetKeyVaultName:               tt.getName,
				KeyVaultOverridesVariable:     tt.overVar,
			}
			if tt.panic {
				require.Panics(t, func() {
					b.loadKeyVault(config, ".")
				})
			} else {
				b.loadKeyVault(config, ".")
				if tt.getName == nil {
					if tt.name == "#5" {
						require.EqualValues(t, "failed to create key vault client", hook.LastEntry().Message)
					} else {
						require.EqualValues(t, "GetKeyVaultName or KeyVaultName is required", hook.LastEntry().Message)
					}
				} else {
					require.EqualValues(t, "GetKeyVaultResourceGroup or KeyVaultResourceGroup is required", hook.LastEntry().Message)
				}

			}
			v.Reset()
		})
	}
}

func (mb *mbMock) ServeMock(srv *grpc.Server, lis net.Listener) error {
	args := mb.Called(lis)
	return args.Error(0)
}

func Test_StartRpcServer(t *testing.T) {
	testCases := []struct {
		testName, port string
		mockErr        error
	}{
		{testName: "grpc port is not set"},

		{testName: "failed to create gprc listener", port: "a", mockErr: errors.New("panic")},
	}
	var srvMock mbMock
	h, err := mpatch.PatchInstanceMethodByName(reflect.TypeOf(grpc.Server{}), "Serve", srvMock.ServeMock)
	require.NoError(t, err)
	for _, tt := range testCases {
		t.Run(tt.testName, func(t *testing.T) {
			logger, _ := l.CreateNullLogger()
			b := mainBuilder{logger: logger, rpcPort: tt.port, rpcServer: &grpc.Server{}}
			srvMock.On("Serve", mock.Anything).Return(tt.mockErr).Once()
			require.Panics(t, func() {
				b.StartRpcServer()
			})
		})
	}
	h.Unpatch()
	t.Run("grpc server have been successfully run", func(t *testing.T) {
		go func() {
			var (
				srv     grpc.Server
				srvMock mbMock
			)
			h, err = mpatch.PatchInstanceMethodByName(reflect.TypeOf(srv), "Serve", srvMock.ServeMock)
			require.NoError(t, err)
			defer h.Unpatch()
			srvMock.On("ServeMock", mock.Anything).Return(nil).Once()
			logger, _ := l.CreateNullLogger()
			b := mainBuilder{logger: logger, rpcPort: "8092", rpcServer: &srv}
			// defer b.StopRpcServer()
			b.StartRpcServer()
		}()
		require.Equal(t, syscall.SIGTERM, <-terminator.ApplicationChannel)

	})
}

func (mb *mbMock) OpenMock(dbDialect gorm.Dialector, opt ...gorm.Option) (*gorm.DB, error) {
	args := mb.Called(dbDialect, opt)
	if args.Get(0) == nil {
		return &gorm.DB{}, args.Error(1)
	}
	dbase := args.Get(0).(gorm.DB)
	return &dbase, args.Error(1)
}

func Test_InitializeGORM(t *testing.T) {
	var gormMock mbMock
	h, err := mpatch.PatchMethod(gorm.Open, gormMock.OpenMock)
	require.NoError(t, err)
	defer h.Unpatch()

	logger, _ := l.CreateNullLogger()
	viper := v.New()
	viper.Set("gormDebug", "true")
	b := mainBuilder{logger: logger, viper: viper}

	gormMock.On("OpenMock", mock.Anything, mock.Anything).Return(gorm.DB{}, nil)
	dbase, err := b.InitializeGORM("localhost:1433")
	require.NotNil(t, dbase)
	require.NoError(t, err)
}

func Test_StopRpcServer(t *testing.T) {
	h, err := mpatch.PatchInstanceMethodByName(reflect.TypeOf(&grpc.Server{}), "Stop", func(cl *grpc.Server) {})
	require.NoError(t, err)
	defer h.Unpatch()

	logger, _ := l.CreateNullLogger()
	b := mainBuilder{logger: logger, rpcPort: "8093", rpcServer: &grpc.Server{}}
	require.NotPanics(t, func() {
		b.StopRpcServer()
	})
}

func (mb *mbMock) ListenAndServeMock(srv *http.Server) error {
	args := mb.Called()
	return args.Error(0)
}

func Test_StartHTTPServer(t *testing.T) {
	var (
		srv     http.Server
		srvMock mbMock
	)
	h, err := mpatch.PatchInstanceMethodByName(reflect.TypeOf(srv), "ListenAndServe", srvMock.ListenAndServeMock)
	require.NoError(t, err)
	defer h.Unpatch()
	go func() {
		logger, _ := l.CreateNullLogger()
		srvMock.On("ListenAndServeMock").Return(nil).Once()
		b := mainBuilder{logger: logger, httpPort: "8080", httpRouter: gin.New()}
		b.StartHttpServer()
	}()
	require.Equal(t, syscall.SIGTERM, <-terminator.ApplicationChannel)
}

func Test_Close(t *testing.T) {
	logger, _ := l.CreateNullLogger()
	h, err := mpatch.PatchInstanceMethodByName(reflect.TypeOf(&grpc.Server{}), "Stop", func(cl *grpc.Server) {})
	require.NoError(t, err)
	defer h.Unpatch()
	b := mainBuilder{
		logger:         logger,
		noSqlClient:    &mongo.Client{},
		httpServer:     &http.Server{},
		rpcServer:      &grpc.Server{},
		onStop:         func(b MainBuilder) {},
		rpcConnections: []*grpc.ClientConn{{}, {}},
	}
	b.Close()
}

func (mb *mbMock) PingSQLClientMock(logger l.Logger, cl *gorm.DB) error {
	args := mb.Called(logger, cl)
	return args.Error(0)
}

func (mb *mbMock) PingMongoClientMock(cl *mongo.Client) error {
	args := mb.Called(cl)
	return args.Error(0)
}

func Test_httpHandlers(t *testing.T) {
	testCases := []struct {
		testName, rout   string
		mongoErr, sqlErr error
		expCode          int
	}{
		{testName: "failed to ping mongo and sql", mongoErr: errors.New("error"), sqlErr: errors.New("error"),
			expCode: 424, rout: "/health"},
		{testName: "everything is OK", expCode: 200, rout: "/health"},

		{testName: "failed to ping mongo and sql", mongoErr: errors.New("error"), sqlErr: errors.New("error"),
			expCode: 424, rout: "/readiness"},
		{testName: "everything is OK", expCode: 200, rout: "/readiness"},
	}
	gin.SetMode(gin.TestMode)
	mockServer := gin.Default()
	mockDb := mbMock{}
	mongoDb := mbMock{}
	logger, _ := l.CreateNullLogger()
	b := mainBuilder{
		logger:               logger,
		noSqlClient:          &mongo.Client{},
		httpPort:             "8080",
		httpServer:           &http.Server{},
		httpRouter:           mockServer,
		sqlClient:            &gorm.DB{},
		readinessCheckConfig: &CheckConfig{Check: func(ctx *gin.Context) {}},
		wellnessCheckConfig:  &CheckConfig{Check: func(ctx *gin.Context) {}},
	}
	b.httpRouter.GET("/health", b.healthHandler)
	b.httpRouter.GET("/readiness", b.readinessHandler)
	l.RegisterHTTPEndpointsWithGin(b.httpRouter)

	h, err := mpatch.PatchMethod(db.PingSqlClient, mockDb.PingSQLClientMock)
	require.NoError(t, err)
	defer h.Unpatch()
	h2, err := mpatch.PatchMethod(db.PingNoSqlClient, mongoDb.PingMongoClientMock)
	require.NoError(t, err)
	defer h2.Unpatch()
	for _, tt := range testCases {
		t.Run(tt.testName, func(t *testing.T) {
			mockDb.On("PingSQLClientMock", mock.Anything, b.sqlClient).Return(tt.sqlErr).Once()
			mongoDb.On("PingMongoClientMock", b.noSqlClient).Return(tt.mongoErr).Once()
			request, _ := http.NewRequest("GET", tt.rout, nil)
			response := httptest.NewRecorder()
			mockServer.ServeHTTP(response, request)
			require.Equal(t, tt.expCode, response.Code, tt.rout)
		})
	}
}

func Test_Run(t *testing.T) {
	viper := v.New()
	viper.Set("datadog.disableTracer", false)
	logger, _ := l.CreateNullLogger()
	var grpcMock mbMock
	h, err := mpatch.PatchInstanceMethodByName(reflect.TypeOf(grpc.Server{}), "Serve", grpcMock.ServeMock)
	require.NoError(t, err)
	defer h.Unpatch()
	grpcMock.On("ServeMock", mock.Anything).Return(nil).Once()
	b := mainBuilder{
		viper:      viper,
		logger:     logger,
		onRun:      func(b MainBuilder) {},
		onStop:     func(b MainBuilder) {},
		httpRouter: gin.New(),
		rpcServer:  &grpc.Server{},
		rpcPort:    "8094",
	}
	go func() {
		b.Run()
	}()
	require.Equal(t, syscall.SIGTERM, <-terminator.ApplicationChannel)
}
