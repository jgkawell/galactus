package token

import (
	"fmt"

	m "github.com/jgkawell/galactus/pkg/azkeyvault/model"
	l "github.com/jgkawell/galactus/pkg/logging"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/adal"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/pkg/errors"
)

// Service provides functionality surrounding getting necessary tokens from Azure.
type Service interface {
	GetManagementToken(string, string, string, string) (autorest.Authorizer, error)
	GetKeyvaultToken(string, string, string, string) (autorest.Authorizer, error)
	getServicePrincipalToken(string, *azure.Environment, string, string, string) (*adal.ServicePrincipalToken, error)
	parseAzureEnvironment(string) (*azure.Environment, error)
}

type token struct {
	config m.AzureAuthConfig
}

// NewTokenService creates a service to provide functionality to get tokens from Azure.
func NewTokenService(config m.AzureAuthConfig) Service {
	return &token{
		config: config,
	}
}

// GetManagementToken get keyvalt mgmt token for auth.
func (t *token) GetManagementToken(cloudName string, tenantID string, aADClientSecret string, aADClientID string) (authorizer autorest.Authorizer, err error) {
	err = validate(cloudName, tenantID, aADClientSecret, aADClientID)
	if err != nil {
		return nil, err
	}

	env, err := t.parseAzureEnvironment(cloudName)
	if err != nil {
		return nil, l.NewError(err, "failed to parse Azure environment")
	}

	rmEndPoint := env.ResourceManagerEndpoint
	servicePrincipalToken, err := t.getServicePrincipalToken(tenantID, env, rmEndPoint, aADClientSecret, aADClientID)
	if err != nil {
		return nil, l.NewError(err, "failed to get service principal token")
	}
	authorizer = autorest.NewBearerAuthorizer(servicePrincipalToken)
	return authorizer, nil

}

// GetKeyvaultToken get keyvalt token for auth.
func (t *token) GetKeyvaultToken(cloudName string, tenantID string, aADClientSecret string, aADClientID string) (authorizer autorest.Authorizer, err error) {
	err = validate(cloudName, tenantID, aADClientSecret, aADClientID)
	if err != nil {
		return nil, err
	}

	env, err := t.parseAzureEnvironment(cloudName)
	if err != nil {
		return nil, l.NewError(err, "failed to parse Azure environment")
	}

	kvEndPoint, err := fixEndpoint(env.KeyVaultEndpoint)
	if err != nil {
		return nil, err
	}

	servicePrincipalToken, err := t.getServicePrincipalToken(tenantID, env, kvEndPoint, aADClientSecret, aADClientID)
	if err != nil {
		return nil, l.NewError(err, "failed to get service principal token")
	}
	authorizer = autorest.NewBearerAuthorizer(servicePrincipalToken)
	return authorizer, nil
}

// getServicePrincipalToken creates a new service principal token based on the configuration.
func (t *token) getServicePrincipalToken(tenantID string, env *azure.Environment, resource string, aADClientSecret string, aADClientID string) (*adal.ServicePrincipalToken, error) {
	oauthConfig, err := adal.NewOAuthConfig(env.ActiveDirectoryEndpoint, tenantID)
	if err != nil {
		return nil, l.NewError(err, "failed creating the OAuth config")
	}

	if len(aADClientSecret) > 0 {
		return adal.NewServicePrincipalToken(
			*oauthConfig,
			aADClientID,
			aADClientSecret,
			resource)
	}

	return nil, fmt.Errorf("no credentials provided for AAD application %s", aADClientID)
}

// parseAzureEnvironment returns azure environment by name.
func (t *token) parseAzureEnvironment(cloudName string) (*azure.Environment, error) {
	if cloudName == "" {
		return &azure.PublicCloud, nil
	}
	env, err := azure.EnvironmentFromName(cloudName)
	return &env, errors.Wrapf(err, "failed to get environment from cloudName: %s", cloudName)
}

// fixEndpoint removes trailing '/' characters from a string.
func fixEndpoint(kvEndPoint string) (string, error) {
	if kvEndPoint == "" {
		return "", errors.New("Missing Endpoint")
	} else if '/' == kvEndPoint[len(kvEndPoint)-1] {
		return kvEndPoint[:len(kvEndPoint)-1], nil
	} else {
		return kvEndPoint, nil
	}
}

// validate checks that each parameter is not null.
func validate(cloudName string, tenantID string, aADClientSecret string, aADClientID string) error {
	if cloudName == "" {
		return errors.New("Invalid cloudName")
	} else if tenantID == "" {
		return errors.New("Invalid tenantID")
	} else if aADClientSecret == "" {
		return errors.New("Invalid aADClientSecret")
	} else if aADClientID == "" {
		return errors.New("Invalid aADClientID")
	} else {
		return nil
	}
}
