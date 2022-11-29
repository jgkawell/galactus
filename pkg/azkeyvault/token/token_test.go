package token

import (
	"testing"

	m "github.com/jgkawell/galactus/pkg/azkeyvault/model"

	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestParseAzureEnvironment(t *testing.T) {
	var tests = []struct {
		name        string
		cloud       string
		expectedEnv *azure.Environment
		err         bool
	}{
		{"Empty Cloud Name", "", &azure.PublicCloud, false},
		{"Invalid Cloud Name", "InvalidCloud", &azure.Environment{}, true},
		{"Valid Cloud Name", "AZUREPUBLICCLOUD", &azure.PublicCloud, false},
		{"NonPublic Cloud Name", "AZUREUSGOVERNMENTCLOUD", &azure.USGovernmentCloud, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := NewTokenService(m.AzureAuthConfig{})
			env, err := ts.parseAzureEnvironment(tt.cloud)
			assert.Equal(t, env, tt.expectedEnv)
			if tt.err {
				assert.NotNil(t, err)
			}
		})
	}
}

func TestFixEndpoint(t *testing.T) {
	var tests = []struct {
		name     string
		endpoint string
		err      error
	}{
		{"Trailing Slash", "https://www.test.com/endpoint/", nil},
		{"No Trailing Slash", "https://www.test.com/endpoint", nil},
		{"Empty String", "", errors.New("Missing Endpoint")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			newEndpoint, err := fixEndpoint(tt.endpoint)
			if tt.err != nil {
				assert.Equal(t, tt.err.Error(), err.Error())
				assert.Equal(t, "", newEndpoint)
			} else {
				assert.Equal(t, "https://www.test.com/endpoint", newEndpoint)
				assert.Equal(t, nil, err)
			}
		})
	}
}

func TestValidate(t *testing.T) {
	var tests = []struct {
		name            string
		cloud           string
		tenantID        string
		aADClientSecret string
		aADClientID     string
		err             error
	}{
		{"No Error", "test", "test", "test", "test", nil},
		{"Empty Cloud Name", "", "test", "test", "test", errors.New("Invalid cloudName")},
		{"Empty tenantID", "test", "", "test", "test", errors.New("Invalid tenantID")},
		{"Empty aADClientSecret", "test", "test", "", "test", errors.New("Invalid aADClientSecret")},
		{"Empty aADClientID", "test", "test", "test", "", errors.New("Invalid aADClientID")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validate(tt.cloud, tt.tenantID, tt.aADClientSecret, tt.aADClientID)
			if tt.err == nil {
				assert.Nil(t, err)
			} else {
				assert.Equal(t, err.Error(), tt.err.Error())
			}
		})
	}
}
