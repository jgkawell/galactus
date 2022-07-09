package env

import (
	"os"
	"testing"

	l "github.com/circadence-official/galactus/pkg/logging/v2"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

var (
	expectedLocalConfigurations                  = map[string]interface{}{"namespace": "TILLER_NAMESPACE", "service": map[string]interface{}{"name": "env_test", "type": "test", "version": "v1"}}
	expectedEnvironmentConfiguration             = map[string]interface{}{"__executecmdhandler__": true, "namespace": "TILLER_NAMESPACE", "service": map[string]interface{}{"name": "env_test", "type": "test", "version": "v1"}}
	expectedEnvironmentConfigurationWithServeArg = map[string]interface{}{"__executecmdhandler__": false, "namespace": "TILLER_NAMESPACE", "service": map[string]interface{}{"name": "env_test", "type": "test", "version": "v1"}}
	expectedEnvironmentConfigurationWithPairArg  = map[string]interface{}{"key": "value", "__executecmdhandler__": true, "namespace": "TILLER_NAMESPACE", "service": map[string]interface{}{"name": "env_test", "type": "test", "version": "v1"}}
)

func TestExecuteCmdHandler(t *testing.T) {
	viper.Reset()
	defer viper.Reset()

	viper.Set(executeCmdHandlerKey, true)
	assert.True(t, ExecuteCmdHandler())
	viper.Set(executeCmdHandlerKey, false)
	assert.False(t, ExecuteCmdHandler())
}

func TestReadLocalConfigurationSuccess(t *testing.T) {
	viper.Reset()
	defer viper.Reset()

	logger, _ := l.CreateNullLogger()
	err := readConfiguration(logger, "./testData")
	assert.NoError(t, err)
	assert.Equal(t, expectedLocalConfigurations, viper.GetViper().AllSettings())
}

func TestReadLocalConfigurationFailed(t *testing.T) {
	viper.Reset()
	defer viper.Reset()

	logger, _ := l.CreateNullLogger()
	err := readConfiguration(logger, "")
	assert.Error(t, err)
}

func TestReadEnvironmentConfigurationsSuccess(t *testing.T) {
	viper.Reset()
	defer viper.Reset()

	logger, _ := l.CreateNullLogger()
	os.Args = make([]string, 0)
	err := ReadEnvironmentConfigurations(logger, "./testData")
	assert.NoError(t, err)
	assert.Equal(t, expectedEnvironmentConfiguration, viper.GetViper().AllSettings())
}

func TestReadEnvironmentConfigurationsWithServeArg(t *testing.T) {
	viper.Reset()
	defer viper.Reset()

	logger, _ := l.CreateNullLogger()
	os.Args = make([]string, 0)
	os.Args = append(os.Args, "serve")
	err := ReadEnvironmentConfigurations(logger, "./testData")
	assert.NoError(t, err)
	assert.Equal(t, expectedEnvironmentConfigurationWithServeArg, viper.GetViper().AllSettings())
}

func TestReadEnvironmentConfigurationsWithPairArg(t *testing.T) {
	viper.Reset()
	defer viper.Reset()

	logger, _ := l.CreateNullLogger()
	os.Clearenv()
	os.Args = make([]string, 0)
	os.Args = append(os.Args, "key=value")
	err := ReadEnvironmentConfigurations(logger, "./testData")
	assert.NoError(t, err)
	assert.Equal(t, expectedEnvironmentConfigurationWithPairArg, viper.GetViper().AllSettings())
}

func TestReadEnvironmentConfigurationsFailed(t *testing.T) {
	viper.Reset()
	defer viper.Reset()

	logger, _ := l.CreateNullLogger()
	err := ReadEnvironmentConfigurations(logger, "")
	assert.Error(t, err)
}
