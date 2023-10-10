package database

import (
	context "context"

	"github.com/jgkawell/galactus/pkg/chassis/env"
)

type Client interface {
	Initialize(ctx context.Context, config env.Reader) (err error)
	Ping(ctx context.Context) (err error)
	Disconnect(ctx context.Context) (err error)
}
