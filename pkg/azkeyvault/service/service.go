package service

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httputil"
	"time"

	m "github.com/jgkawell/galactus/pkg/azkeyvault/model"
	l "github.com/jgkawell/galactus/pkg/logging/v2"

	kv "github.com/Azure/azure-sdk-for-go/services/keyvault/2016-10-01/keyvault"
	kvmgmt "github.com/Azure/azure-sdk-for-go/services/keyvault/mgmt/2016-10-01/keyvault"
	"github.com/Azure/go-autorest/autorest"
	lru "github.com/hashicorp/golang-lru"
)

const urlCacheSize = 8
const secretCacheSize = 256

// KeyVaultService provides access to the Azure KeyVault API
type KeyVaultService interface {
	SetSecret(context.Context, l.Logger, string, string, string, string) l.Error
	GetSecret(context.Context, l.Logger, string, string, string) (string, l.Error)
	DeleteSecret(ctx context.Context, logger l.Logger, resourceGroup, keyVault, key string) l.Error
	getVaultURL(context.Context, l.Logger, m.AzureAuthConfig, string, string) (*string, l.Error)
}

type keyVaultService struct {
	config       m.AzureAuthConfig
	kvClient     kv.BaseClient
	vaultsClient kvmgmt.VaultsClient
	urlCache     *lru.Cache
	secretCache  *lru.TwoQueueCache
}

type secretCacheEntry struct {
	value  string
	expiry *time.Time
}

// NewService will create a new KeyVaultService given applicable tokens.
func NewService(config m.AzureAuthConfig, mgmtToken autorest.Authorizer, kvToken autorest.Authorizer, logger l.Logger) KeyVaultService {
	vc := kvmgmt.NewVaultsClient(config.SubscriptionID)
	logger.WithField("vaults_client", vc).Info("created vaults client")
	vc.Authorizer = mgmtToken
	vc.RequestInspector = LogRequest(logger)
	vc.ResponseInspector = LogResponse(logger)

	kvc := kv.New()
	logger.WithField("kv_client", kvc).Info("created kv client")
	kvc.Authorizer = kvToken

	urlCache, err := lru.New(urlCacheSize)
	if err != nil {
		logger.WithError(err).Error("failed to create url cache")
	}
	secretCache, err := lru.New2Q(secretCacheSize)
	if err != nil {
		logger.WithError(err).Error("failed to create secret cache")
	}

	return &keyVaultService{
		config:       config,
		kvClient:     kvc,
		vaultsClient: vc,
		urlCache:     urlCache,
		secretCache:  secretCache,
	}
}

func LogRequest(logger l.Logger) autorest.PrepareDecorator {
	return func(p autorest.Preparer) autorest.Preparer {
		return autorest.PreparerFunc(func(r *http.Request) (*http.Request, error) {
			r, err := p.Prepare(r)
			if err != nil {
				logger.WithError(err).Error("failed to prepare request")
			}
			dump, _ := httputil.DumpRequestOut(r, true)
			logger.Debug(string(dump))
			return r, err
		})
	}
}

func LogResponse(logger l.Logger) autorest.RespondDecorator {
	return func(p autorest.Responder) autorest.Responder {
		return autorest.ResponderFunc(func(r *http.Response) error {
			err := p.Respond(r)
			if err != nil {
				logger.WithError(err).Error("failed to respond")
			}
			dump, _ := httputil.DumpResponse(r, true)
			logger.Debug(string(dump))
			return err
		})
	}
}

// SetSecret on specified Azure KeyVault.
func (s *keyVaultService) SetSecret(ctx context.Context, logger l.Logger, resourceGroup, keyVault, secret, value string) l.Error {
	logger = logger.WithFields(l.Fields{
		"resource_group": resourceGroup,
		"key_vault":      keyVault,
		"secret":         secret,
	})

	vaultURL, customErr := s.getVaultURL(ctx, logger, s.config, resourceGroup, keyVault)
	if customErr != nil {
		return logger.WrapError(customErr)
	}
	logger.WithField("vault_url", *vaultURL).Info("got vault url")

	params := kv.SecretSetParameters{Value: &value}
	svalue, err := s.kvClient.SetSecret(ctx, *vaultURL, secret, params)
	if err != nil {
		return logger.WrapError(err)
	}

	if s.secretCache != nil {
		// insert into LRU cache
		key := secretKey(resourceGroup, keyVault, secret)
		expiry := svalue.Attributes.Expires
		s.secretCache.Add(key, &secretCacheEntry{
			value:  string(*svalue.Value),
			expiry: (*time.Time)(expiry),
		})
	}
	return nil
}

func secretKey(resourceGroup, keyVault, secret string) string {
	return fmt.Sprintf("%s-%s-%s", resourceGroup, keyVault, secret)
}

// GetSecret from specified Azure KeyVault.
func (s *keyVaultService) GetSecret(ctx context.Context, logger l.Logger, resourceGroup, keyVault, secret string) (string, l.Error) {

	logger = logger.WithFields(l.Fields{
		"resource_group": resourceGroup,
		"key_vault":      keyVault,
		"secret":         secret,
	})
	key := secretKey(resourceGroup, keyVault, secret)

	if s.secretCache != nil {
		// check LRU cache for secret and return if not expired
		if value, ok := s.secretCache.Get(key); ok {
			cachedSecret := value.(*secretCacheEntry)
			if cachedSecret.expiry == nil || cachedSecret.expiry.After(time.Now()) {
				logger.Debug("return cached secret")
				return cachedSecret.value, nil
			}
		}
	}

	vaultURL, err := s.getVaultURL(ctx, logger, s.config, resourceGroup, keyVault)
	if err != nil {
		return secret, logger.WrapError(err)
	}
	logger.WithField("vault_url", *vaultURL).Info("got vault url")

	svalue, customErr := s.kvClient.GetSecret(ctx, *vaultURL, secret, "")
	if customErr != nil {
		return secret, logger.WrapError(customErr)
	}

	if s.secretCache != nil {
		expiry := svalue.Attributes.Expires
		s.secretCache.Add(key, &secretCacheEntry{
			value:  string(*svalue.Value),
			expiry: (*time.Time)(expiry),
		})
	}

	return string(*svalue.Value), nil
}

func urlKey(resourceGroup, vaultName string) string {
	return fmt.Sprintf("%s-%s", resourceGroup, vaultName)
}

// getVaultURL gets url of vault based on cloud, rg, and vault name
func (s *keyVaultService) getVaultURL(ctx context.Context, logger l.Logger, config m.AzureAuthConfig, resourceGroup string, vaultName string) (*string, l.Error) {
	logger = logger.WithFields(l.Fields{
		"subscription_id": config.SubscriptionID,
		"vault_name":      vaultName,
		"resource_group":  resourceGroup,
	})
	logger.Debug("getting Vault URL")

	key := urlKey(resourceGroup, vaultName)
	if s.urlCache != nil {
		logger.Debug("checking for cached URL")
		if value, ok := s.urlCache.Get(key); ok {
			logger.WithField("url", value).Debug("returning cached URL")
			return value.(*string), nil
		}
		logger.Debug("cached URL does not exist")
	}

	logger.Debug("retrieving key vault URL from service")
	vault, err := s.vaultsClient.Get(ctx, resourceGroup, vaultName)
	if err != nil {
		return nil, logger.WrapError(err)
	}

	if s.urlCache != nil {
		logger.Debug("adding URL to local cache")
		s.urlCache.Add(key, vault.Properties.VaultURI)
	}
	return vault.Properties.VaultURI, nil
}

func (s *keyVaultService) DeleteSecret(ctx context.Context, logger l.Logger, resourceGroup, keyVault, secret string) l.Error {
	logger = logger.WithFields(l.Fields{
		"resourceGroup": resourceGroup,
		"keyVault":      keyVault,
		"secret":        secret,
	})
	vaultURL, customErr := s.getVaultURL(ctx, logger, s.config, resourceGroup, keyVault)
	if customErr != nil {
		return logger.WrapError(customErr)
	}
	logger.WithField("vault_url", *vaultURL).Info("got vault url")

	_, err := s.kvClient.DeleteSecret(ctx, *vaultURL, secret)
	if err != nil {
		return logger.WrapError(err)
	}
	if s.secretCache != nil {
		// remove from LRU cache
		key := secretKey(resourceGroup, keyVault, secret)
		s.secretCache.Remove(key)
	}
	return nil
}
