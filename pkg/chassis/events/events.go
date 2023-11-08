package events

import (
	cf "github.com/jgkawell/galactus/pkg/chassis/clientfactory"
	c "github.com/jgkawell/galactus/pkg/chassis/context"

	evpb "github.com/jgkawell/galactus/api/gen/go/core/eventer/v1"
	rgpb "github.com/jgkawell/galactus/api/gen/go/core/registry/v1"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type EventManager interface {
	CreateAndSendEvent(ctx c.Context, event protoreflect.ProtoMessage) error
	Close()

	// TODO: ThrowSystemError - Create a system level error that is related to core components of the system, and not application level
}

type manager struct {
	client evpb.EventerClient
	conn   *grpc.ClientConn
}

func NewEventManager(ctx c.Context, registryAddress string) (EventManager, error) {
	// create registry client
	clientFactory := cf.NewClientFactory()
	registryClient, registryConn, err := clientFactory.CreateRegistryClient(ctx.Logger(), registryAddress)
	defer registryConn.Close()
	if err != nil {
		return nil, err
	}

	// get eventer address from registry
	resp, err := registryClient.Connection(ctx, &rgpb.ConnectionRequest{
		Route: evpb.Eventer_ServiceDesc.ServiceName,
	})
	if err != nil {
		return nil, err
	}

	// create eventer client
	eventerClient, conn, err := clientFactory.CreatEventerClient(ctx.Logger(), resp.GetAddress())
	if err != nil {
		return nil, err
	}

	return &manager{
		client: eventerClient,
		conn:   conn,
	}, nil
}

// CreateAndSendEvent creates an event on the eventer service
func (m *manager) CreateAndSendEvent(ctx c.Context, event protoreflect.ProtoMessage) error {
	ctx, span := ctx.Span()
	defer span.End()

	e, err := New(event)
	if err != nil {
		return err
	}
	req := evpb.EmitRequest{
		Event: e,
	}
	_, err = m.client.Emit(ctx, &req)
	if err != nil {
		return err
	}
	return nil
}

func (m *manager) Close() {
	m.conn.Close()
}
