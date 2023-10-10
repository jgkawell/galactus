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

// registry reserves 3500 (http) and 3501 (grpc) for itself when running locally
const localMinPort = 3502
const localMaxPort = 4500

func (s *service) generatePort(ctx ct.ExecutionContext, kind pb.ServerKind) (string, l.Error) {
	var port int32
	var err l.Error
	if s.isDevMode {
		port, err = s.generateLocalPort(ctx.Logger)
		if err != nil {
			return "", ctx.Logger.WrapError(err)
		}
	} else {
		port, err = s.generateRemotePort(ctx.Logger, kind)
		if err != nil {
			return "", ctx.Logger.WrapError(err)
		}
	}
	return strconv.Itoa(int(port)), nil
}

/*
generateLocalPort will generate a random port between `localMinPort` and `localMaxPort` making sure it is not already in use.

If an unused port is not found after 1000 attempts, an error is returned.
*/
func (s *service) generateLocalPort(logger l.Logger) (int32, l.Error) {
	for i := 0; i < localPortConflictRetryLimit; i++ {
		randomPort := rand.Intn(localMaxPort-localMinPort) + localMinPort
		var count int64
		err := s.db.Model(&pb.ServerORM{}).Where("port = ?", strconv.Itoa(randomPort)).Count(&count).Error
		if err != nil {
			return 0, logger.WrapError(l.NewError(err, "failed to query for port usage while generating random local port"))
		}
		if count == 0 {
			return int32(randomPort), nil
		}
	}
	return 0, logger.WrapError(errors.New("failed to generate local port after maximum conflict retries"))
}

/*
generateRemotePort will generate the remote port based on the protocol kind:

	http = 8080
	grpc = 8090

If an invalid protocol kind is given, an error is returned.
*/
func (s *service) generateRemotePort(logger l.Logger, kind pb.ServerKind) (int32, l.Error) {
	switch kind {
	case pb.ServerKind_SERVER_KIND_HTTP:
		return 8080, nil
	case pb.ServerKind_SERVER_KIND_GRPC:
		return 8090, nil
	default:
		return 0, logger.WrapError(errors.New("unsupported protocol kind"))
	}
}
