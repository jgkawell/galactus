package secrets

import (
	"context"

	"github.com/jgkawell/galactus/pkg/chassis/env"
)

type Client interface {
	// Initialize will initialize the client with the given configuration
	Initialize(ctx context.Context, config env.Reader) (err error)
	// Get retrieves a secret from the vault for the given key
	Get(ctx context.Context, key string) (value string, err error)
	// Set will add or update a secret value for the given key
	Set(ctx context.Context, key string, value string) (err error)
	// Delete will remove a secret from the vault for the given key
	Delete(ctx context.Context, key string) (err error)
}
