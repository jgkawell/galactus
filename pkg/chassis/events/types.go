package events

import (
	"context"
	"fmt"
	"strings"

	eventspb "github.com/cloudevents/sdk-go/binding/format/protobuf/v2/pb"
	"github.com/jgkawell/galactus/pkg/chassis/messagebus"

	"github.com/google/uuid"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/anypb"
)

func Emit(ctx context.Context, buses []messagebus.Client, msg protoreflect.ProtoMessage, preview bool) error {
	event, err := New(msg)
	if err != nil {
		return err
	}
	params := messagebus.PublishParams{
		Event: event,
	}
	if preview {
		params.Tags = []string{"preview"}
	}
	for _, bus := range buses {
		err := bus.Publish(ctx, params)
		if err != nil {
			return err
		}
	}
	return nil
}

// New creates a new CloudEvent from a proto message
func New(msg protoreflect.ProtoMessage) (*eventspb.CloudEvent, error) {
	// convert the message to an any
	any, err := anypb.New(msg)
	if err != nil {
		return nil, err
	}

	// anypb.New() always prefixes with the url "type.googleapis.com/" so we need to remove it
	typePath := strings.TrimPrefix(any.TypeUrl, "type.googleapis.com/")

	// the type is just the final part of the url (path) and the source is the rest
	parts := strings.Split(typePath, ".")
	eventSource := strings.Join(parts[:len(parts)-1], ".")
	eventType := parts[len(parts)-1]

	// replace the base url with our own
	// TODO: this should be configurable
	any.TypeUrl = fmt.Sprintf("github.com/jgkawell/galactus/api/gen/go/%s", typePath)

	// build and return the event
	return &eventspb.CloudEvent{
		Id:          uuid.New().String(),
		Source:      eventSource,
		SpecVersion: "1.0",
		Type:        eventType,
		Data: &eventspb.CloudEvent_ProtoData{
			ProtoData: any,
		},
	}, nil
}

// Extract extracts the data from a CloudEvent into the underlying proto message
func Extract(event *eventspb.CloudEvent, data protoreflect.ProtoMessage) error {
	d := event.GetProtoData()
	if d == nil {
		return fmt.Errorf("unable to extract event data")
	}
	err := d.UnmarshalTo(data)
	if err != nil {
		return err
	}
	return nil
}

func Validate(msg protoreflect.ProtoMessage) error {
	// convert the message to an any
	_, err := anypb.New(msg)
	if err != nil {
		return err
	}
	return nil
}
