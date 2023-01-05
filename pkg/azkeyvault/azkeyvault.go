// Package azkeyvault provided functionality to access the Azure KeyVault and Manage Azure Tokens
package azkeyvault

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"time"

	m "github.com/jgkawell/galactus/pkg/azkeyvault/model"
	s "github.com/jgkawell/galactus/pkg/azkeyvault/service"
	t "github.com/jgkawell/galactus/pkg/azkeyvault/token"
	l "github.com/jgkawell/galactus/pkg/logging"
)

// KeyVaultClient key vault client interface
type KeyVaultClient interface {
	// GetKeyVaultSecret retrieves secret value from rg/vault/secret.
	GetKeyVaultSecret(ctx context.Context, logger l.Logger, secret string) (string, l.Error)
	// SetKeyVaultSecret Add/update secret.
	SetKeyVaultSecret(ctx context.Context, logger l.Logger, secret string, value string) l.Error
	// Delete keyvault secret
	DeleteKeyVaultSecret(ctx context.Context, logger l.Logger, secret string) l.Error
}

type keyVaultClient struct {
	ks            s.KeyVaultService
	logger        l.Logger
	resourceGroup string
	keyVault      string
}

// newKeyVaultClient key vault client
func newKeyVaultClient(ts t.Service, config m.AzureAuthConfig, logger l.Logger, resourceGroup, keyVault string) (KeyVaultClient, l.Error) {
	mgmtToken, err := ts.GetManagementToken(
		config.Cloud,
		config.TenantID,
		config.AADClientSecret,
		config.AADClientID)
	if err != nil {
		return nil, logger.WrapError(l.NewError(err, "failed to get management token"))
	}
	kvToken, err := ts.GetKeyvaultToken(
		config.Cloud,
		config.TenantID,
		config.AADClientSecret,
		config.AADClientID)
	if err != nil {
		return nil, logger.WrapError(l.NewError(err, "failed to get key vault token"))
	}

	// use tokens to initialize keyvault service
	ks := s.NewService(config, mgmtToken, kvToken, logger)

	return &keyVaultClient{
		ks:            ks,
		logger:        logger.WithField("struct", "KeyVaultClient"),
		resourceGroup: resourceGroup,
		keyVault:      keyVault,
	}, nil
}

// NewClientConfigStruct - takes configuration struct instead of reading a file in the system.
func NewClientConfigStruct(logger l.Logger, config m.AzureAuthConfig, resourceGroup, kvName string) (KeyVaultClient, l.Error) {
	logger = logger.WithField("cloud", config.Cloud)
	logger.Info("initializing keyvault client via config struct")
	ts := t.NewTokenService(config)
	return newKeyVaultClient(ts, config, logger, resourceGroup, kvName)
}

// NewClientConfigPath - takes a configuration file path instead of reading a file in the system.
func NewClientConfigPath(logger l.Logger, configPath, resourceGroup, keyVault string) (KeyVaultClient, l.Error) {
	logger = logger.WithField("config_path", configPath)
	logger.Info("initializing keyvault client via config path")

	// Parse Config Map
	var config m.AzureAuthConfig
	nbytes, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, logger.WrapError(err)
	}

	// Unmarshal configmap into struct
	err = json.Unmarshal(nbytes, &config)
	if err != nil {
		return nil, logger.WrapError(err)
	}

	// initialize token service and use it to get tokens
	ts := t.NewTokenService(config)
	return newKeyVaultClient(ts, config, logger, resourceGroup, keyVault)
}

// GetKeyVaultSecret lookup secret value from rg/vault/secret
func (c *keyVaultClient) GetKeyVaultSecret(ctx context.Context, logger l.Logger, secret string) (string, l.Error) {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	c.logger.WithFields(l.Fields{
		"resource_group": c.resourceGroup,
		"key_vault":      c.keyVault,
		"secret":        secret,
	}).Debug("getting keyvault secret")

	return c.ks.GetSecret(ctx, logger, c.resourceGroup, c.keyVault, secret)
}

// SetKeyVaultSecret Add/update secret.
func (c *keyVaultClient) SetKeyVaultSecret(ctx context.Context, logger l.Logger, secret, value string) l.Error {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	logger.WithFields(l.Fields{
		"resource_group": c.resourceGroup,
		"key_vault":      c.keyVault,
		"secret":        secret,
	}).Info("setting keyvault secret")
	return c.ks.SetSecret(ctx, logger, c.resourceGroup, c.keyVault, secret, value)
}

// DeleteKeyVaultSecret remove secret from key vault
func (c *keyVaultClient) DeleteKeyVaultSecret(ctx context.Context, logger l.Logger, secret string) l.Error {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	logger.WithFields(l.Fields{
		"resource_group": c.resourceGroup,
		"key_vault":      c.keyVault,
		"secret":        secret,
	}).Info("deleting keyvault secret")
	return c.ks.DeleteSecret(ctx, logger, c.resourceGroup, c.keyVault, secret)
}
