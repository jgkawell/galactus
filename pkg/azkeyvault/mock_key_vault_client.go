// Code generated by mockery v2.9.4. DO NOT EDIT.

package azkeyvault

import (
	context "context"

	logging "github.com/circadence-official/galactus/pkg/logging/v2"
	mock "github.com/stretchr/testify/mock"
)

// MockKeyVaultClient is an autogenerated mock type for the KeyVaultClient type
type MockKeyVaultClient struct {
	mock.Mock
}

// DeleteKeyVaultSecret provides a mock function with given fields: ctx, logger, secret
func (_m *MockKeyVaultClient) DeleteKeyVaultSecret(ctx context.Context, logger logging.Logger, secret string) logging.Error {
	ret := _m.Called(ctx, logger, secret)

	var r0 logging.Error
	if rf, ok := ret.Get(0).(func(context.Context, logging.Logger, string) logging.Error); ok {
		r0 = rf(ctx, logger, secret)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(logging.Error)
		}
	}

	return r0
}

// GetKeyVaultSecret provides a mock function with given fields: ctx, logger, secret
func (_m *MockKeyVaultClient) GetKeyVaultSecret(ctx context.Context, logger logging.Logger, secret string) (string, logging.Error) {
	ret := _m.Called(ctx, logger, secret)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, logging.Logger, string) string); ok {
		r0 = rf(ctx, logger, secret)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 logging.Error
	if rf, ok := ret.Get(1).(func(context.Context, logging.Logger, string) logging.Error); ok {
		r1 = rf(ctx, logger, secret)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(logging.Error)
		}
	}

	return r0, r1
}

// SetKeyVaultSecret provides a mock function with given fields: ctx, logger, secret, value
func (_m *MockKeyVaultClient) SetKeyVaultSecret(ctx context.Context, logger logging.Logger, secret string, value string) logging.Error {
	ret := _m.Called(ctx, logger, secret, value)

	var r0 logging.Error
	if rf, ok := ret.Get(0).(func(context.Context, logging.Logger, string, string) logging.Error); ok {
		r0 = rf(ctx, logger, secret, value)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(logging.Error)
		}
	}

	return r0
}
