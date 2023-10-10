package azkeyvault

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/keyvault/armkeyvault"
	"github.com/jgkawell/galactus/pkg/chassis/env"
	"github.com/jgkawell/galactus/pkg/chassis/secrets"
)

type (
	keyVaultClient struct {
		client        *armkeyvault.SecretsClient
		resourceGroup string
		keyVault      string
	}
	// azureAuthConfig represents the configuration for authenticating to Azure
	azureAuthConfig struct {
		// TenantID is the Azure tenant ID
		TenantID string `json:"tenantId"`
		// ClientID is the Azure client ID
		ClientID string `json:"clientId"`
		// ClientSecret is the Azure client secret
		ClientSecret string `json:"clientSecret"`
		// SubscriptionID is the Azure subscription ID
		SubscriptionID string `json:"subscriptionId"`
		// ResourceGroup is the Azure resource group
		ResourceGroup string `json:"resourceGroup"`
		// KeyVaultName is the Azure key vault name
		KeyVaultName string `json:"keyVaultName"`
	}
)

func New() secrets.Client {
	return &keyVaultClient{}
}

func (c *keyVaultClient) Initialize(ctx context.Context, config env.Reader) error {
	// get the path to the azure config file
	azureJsonPath := config.GetString("azureJsonPath")
	if azureJsonPath == "" {
		return fmt.Errorf("azureJsonPath config value required")
	}

	// read the config file
	var azureConfig azureAuthConfig
	nbytes, err := os.ReadFile(azureJsonPath)
	if err != nil {
		return fmt.Errorf("failed to read azure config file: %w", err)
	}

	// unmarshal file contents into  struct
	err = json.Unmarshal(nbytes, &azureConfig)
	if err != nil {
		return fmt.Errorf("failed to unmarshal azure config: %w", err)
	}

	// create credential
	cred, err := azidentity.NewClientSecretCredential(azureConfig.TenantID, azureConfig.ClientID, azureConfig.ClientSecret, nil)
	if err != nil {
		return fmt.Errorf("failed to create credential: %w", err)
	}

	// create client factory
	clientFactory, err := armkeyvault.NewClientFactory(azureConfig.SubscriptionID, cred, nil)
	if err != nil {
		return fmt.Errorf("failed to create client factory: %w", err)
	}

	// set config values
	c.client = clientFactory.NewSecretsClient()
	c.resourceGroup = azureConfig.ResourceGroup
	c.keyVault = azureConfig.KeyVaultName

	return nil
}

func (c *keyVaultClient) Get(ctx context.Context, key string) (string, error) {
	return "fake_secret", nil
}

func (c *keyVaultClient) Set(ctx context.Context, key string, value string) error {
	return nil
}

func (c *keyVaultClient) Delete(ctx context.Context, key string) error {
	return nil
}
