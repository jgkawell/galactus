package nats

import (
	"context"
	"fmt"
	"sort"
	"strings"

	events "github.com/cloudevents/sdk-go/binding/format/protobuf/v2/pb"
	"github.com/jgkawell/galactus/pkg/chassis/env"
	"github.com/jgkawell/galactus/pkg/chassis/messagebus"

	"github.com/nats-io/nats.go"
	"google.golang.org/protobuf/proto"
)

type (
	wrapper struct {
		connection    *nats.Conn
		configKey     string
		subscriptions map[string]*nats.Subscription
	}
)

// New instantiates a new client wrapper. A call to Initialize is required before use.
// The configKey parameter dictates which key in the configuration file will be read during
// initialization. If configKey is empty, the default value of "nats.url" will be used.
// The configuration can be in various formats, but the following is an example of a yaml file:
//
//	nats:
//	  url: 'nats://localhost:4222'
func New(configKey string) messagebus.Client {
	if configKey == "" {
		configKey = "nats.url"
	}
	return &wrapper{
		configKey:     configKey,
		subscriptions: make(map[string]*nats.Subscription),
	}
}

func (w *wrapper) Initialize(ctx context.Context, config env.Reader) error {
	url := config.GetString(w.configKey)
	nc, err := nats.Connect(url)
	if err != nil {
		return err
	}
	w.connection = nc
	return nil
}

func (w *wrapper) Publish(ctx context.Context, params messagebus.PublishParams) error {
	// serialize the event
	data, err := proto.Marshal(params.Event)
	if err != nil {
		return err
	}
	// publish the event
	err = w.connection.Publish(subject(params.Event, false, params.Tags...), data)
	if err != nil {
		return err
	}
	return nil
}

func (w *wrapper) Subscribe(ctx context.Context, params messagebus.SubscribeParams) error {
	if params.Group != "" {
		return w.queueSubscribe(ctx, params)
	}
	return w.subscribe(ctx, params)
}

// subscribe creates a subscription to the given event and tags
func (w *wrapper) subscribe(ctx context.Context, params messagebus.SubscribeParams) error {
	// subscribe to the subject
	subscription, err := w.connection.Subscribe(subject(params.Event, params.IgnoreType, params.Tags...), func(msg *nats.Msg) {
		// deserialize the event
		event := &events.CloudEvent{}
		err := proto.Unmarshal(msg.Data, event)
		if err != nil {
			msg.Nak()
		}
		// consume the event
		err = params.Consumer.Consume(ctx, event)
		if err != nil {
			msg.Nak()
		}
		msg.Ack()
	})
	if err != nil {
		return err
	}
	w.subscriptions[subject(params.Event, params.IgnoreType, params.Tags...)] = subscription
	return nil
}

// queueSubscribe creates a subscription to the given event and tags using a queue group
func (w *wrapper) queueSubscribe(ctx context.Context, params messagebus.SubscribeParams) error {
	// subscribe to the subject
	subscription, err := w.connection.QueueSubscribe(subject(params.Event, params.IgnoreType, params.Tags...), params.Group, func(msg *nats.Msg) {
		// deserialize the event
		event := &events.CloudEvent{}
		err := proto.Unmarshal(msg.Data, event)
		if err != nil {
			msg.Nak()
		}
		// consume the event
		err = params.Consumer.Consume(ctx, event)
		if err != nil {
			msg.Nak()
		}
		msg.Ack()
	})
	if err != nil {
		return err
	}
	w.subscriptions[subject(params.Event, params.IgnoreType, params.Tags...)] = subscription
	return nil
}

func (w *wrapper) Unsubscribe(ctx context.Context, params messagebus.UnsubscribeParams) error {
	subscription, ok := w.subscriptions[subject(params.Event, params.IgnoreType, params.Tags...)]
	if !ok {
		return fmt.Errorf("no subscription found for event")
	}
	err := subscription.Unsubscribe()
	if err != nil {
		return err
	}
	delete(w.subscriptions, subject(params.Event, params.IgnoreType, params.Tags...))
	return nil
}

func (w *wrapper) Shutdown(force bool) error {
	if force {
		w.connection.Close()
		return nil
	}
	return w.connection.Drain()
}

// HELPERS

// subject returns the subject for the given event and tags. The format will always follow the pattern:
//
// If no tags are provided and ignoreType is true: {event.Source}.>
// If no tags are provided and ignoreType is false: {event.Source}.{event.Type}.>
// If tags are provided and ignoreType is true: {event.Source}.*.{tags}
// If tags are provided and ignoreType is false: {event.Source}.{event.Type}.{tags}
//
// And the tags will be sorted alphabetically.
func subject(event *events.CloudEvent, ignoreType bool, tags ...string) string {
	if len(tags) == 0 {
		if ignoreType {
			return fmt.Sprintf("%s.>", event.Source)
		}
		return fmt.Sprintf("%s.%s.>", event.Source, event.Type)
	}
	sort.Strings(tags)
	if ignoreType {
		return fmt.Sprintf("%s.*.%s", event.Source, strings.Join(tags, "."))
	}
	return fmt.Sprintf("%s.%s.%s", event.Source, event.Type, strings.Join(tags, "."))
}
