package azkeyvault

import (
	"errors"
	"testing"

	m "github.com/circadence-official/galactus/pkg/azkeyvault/model"
	ts "github.com/circadence-official/galactus/pkg/azkeyvault/token"
	l "github.com/circadence-official/galactus/pkg/logging/v2"

	"github.com/stretchr/testify/assert"
)

func TestNewKeyVaultClient(t *testing.T) {
	var tests = []struct {
		name     string
		testFile string
		hasError bool
		logmsg   string
	}{
		{
			name:     "Success",
			testFile: "testdata/sampleAuthConfig.json",
			hasError: false,
		},
		{
			name:     "Can't Read input file",
			testFile: "testdata/nonexistantAuthConfig.json",
			hasError: true,
		},
		{
			name:     "Can't Parse input file",
			testFile: "testdata/malformedAuthConfig.json",
			hasError: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			kvx, err := NewClientConfigPath(l.CreateLogger("info", "azkeyvault_test"), tt.testFile, "rg", "name")
			if tt.hasError {
				assert.NotNil(t, err)
				assert.Nil(t, kvx)

			} else {
				assert.Nil(t, err)
				assert.NotNil(t, kvx)
			}
		})
	}
}

func TestNewKeyVaultClient_Private(t *testing.T) {
	var authConf = m.AzureAuthConfig{
		Cloud:           "AZUREPUBLICCLOUD",
		TenantID:        "Sampletenantid124",
		AADClientID:     "sampleaadclientid123",
		AADClientSecret: "sampleaadclientsecret345",
		SubscriptionID:  "samplesubscription234",
	}
	var tests = []struct {
		name         string
		mTokenError  error
		kvTokenError error
	}{
		{
			name:         "Success",
			mTokenError:  nil,
			kvTokenError: nil,
		},
		{
			name:         "Can't Get management Token",
			mTokenError:  errors.New("managemet token error"),
			kvTokenError: nil,
		},
		{
			name:         "Can't get Key Vault Token",
			mTokenError:  nil,
			kvTokenError: errors.New("keyvault token error"),
		},
	}
	for _, tt := range tests {
		mockLogger, _ := l.CreateNullLogger()
		t.Run(tt.name, func(t *testing.T) {
			mockTs := &ts.MockService{}
			mockTs.On("GetManagementToken",
				authConf.Cloud,
				authConf.TenantID,
				authConf.AADClientSecret,
				authConf.AADClientID).Return(nil, tt.mTokenError)
			mockTs.On("GetKeyvaultToken",
				authConf.Cloud,
				authConf.TenantID,
				authConf.AADClientSecret,
				authConf.AADClientID).Return(nil, tt.kvTokenError)
			kvx, err := newKeyVaultClient(mockTs, authConf, mockLogger.WithField("test", "azkeyvault"), "rg", "name")
			if tt.kvTokenError != nil {
				assert.Nil(t, kvx)
				assert.Equal(t, "failed to get key vault token: "+tt.kvTokenError.Error(), err.Error())
			} else if tt.mTokenError != nil {
				assert.Nil(t, kvx)
				assert.Equal(t, "failed to get management token: "+tt.mTokenError.Error(), err.Error())
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
