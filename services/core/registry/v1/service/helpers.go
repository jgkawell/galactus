package service

import (
	"errors"
	"math/rand"
	"strconv"

	pb "github.com/jgkawell/galactus/api/gen/go/core/registry/v1"

	ct "github.com/jgkawell/galactus/pkg/chassis/context"
	l "github.com/jgkawell/galactus/pkg/logging"
)

const localPortConflictRetryLimit = 1000

// registry reserves 35000 (http) and 35001 (grpc) for itself when running locally
const localMinPort = 35002
const localMaxPort = 45000

func (s *service) generatePort(ctx ct.Context, kind pb.ServerKind) (string, l.Error) {
	ctx, span := ctx.Span()
	defer span.End()

	var port int32
	var err l.Error
	if s.isDevMode {
		port, err = s.generateLocalPort(ctx)
		if err != nil {
			return "", ctx.Logger().WrapError(err)
		}
	} else {
		port, err = s.generateRemotePort(ctx, kind)
		if err != nil {
			return "", ctx.Logger().WrapError(err)
		}
	}
	return strconv.Itoa(int(port)), nil
}

/*
generateLocalPort will generate a random port between `localMinPort` and `localMaxPort` making sure it is not already in use.

If an unused port is not found after 1000 attempts, an error is returned.
*/
func (s *service) generateLocalPort(ctx ct.Context) (int32, l.Error) {
	ctx, span := ctx.Span()
	defer span.End()

	for i := 0; i < localPortConflictRetryLimit; i++ {
		randomPort := rand.Intn(localMaxPort-localMinPort) + localMinPort
		var count int64
		err := s.db.Model(&pb.ServerORM{}).Where("port = ?", strconv.Itoa(randomPort)).Count(&count).Error
		if err != nil {
			return 0, ctx.Logger().WrapError(l.NewError(err, "failed to query for port usage while generating random local port"))
		}
		if count == 0 {
			return int32(randomPort), nil
		}
	}
	return 0, ctx.Logger().WrapError(errors.New("failed to generate local port after maximum conflict retries"))
}

/*
generateRemotePort will generate the remote port based on the protocol kind:

	http = 8080
	grpc = 8090

If an invalid protocol kind is given, an error is returned.
*/
func (s *service) generateRemotePort(ctx ct.Context, kind pb.ServerKind) (int32, l.Error) {
	ctx, span := ctx.Span()
	defer span.End()
	switch kind {
	case pb.ServerKind_SERVER_KIND_HTTP:
		return 8080, nil
	case pb.ServerKind_SERVER_KIND_GRPC:
		return 8090, nil
	default:
		return 0, ctx.Logger().WrapError(errors.New("unsupported protocol kind"))
	}
}
