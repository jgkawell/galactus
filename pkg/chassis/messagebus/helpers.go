package messagebus

import (
	"math/rand"
	"net/url"
	"strings"
	"time"
)

// ExchangeKind converts an iota to a string which is used in the AMQP protocol.
type ExchangeKind int32

const (
	ExchangeKindDirect ExchangeKind = iota
	ExchangeKindFanout
	ExchangeKindTopic
	ExchangeKindHeaders
	// NOTE: This requires the `rabbitmq_delayed_message_exchange` plugin enabled on the server
	// ref: https://blog.rabbitmq.com/posts/2015/04/scheduling-messages-with-rabbitmq
	ExchangeKindDelayed
)

func (k ExchangeKind) String() string {
	switch k {
	case ExchangeKindFanout:
		return "fanout"
	case ExchangeKindDirect:
		return "direct"
	case ExchangeKindTopic:
		return "topic"
	case ExchangeKindHeaders:
		return "headers"
	case ExchangeKindDelayed:
		return "x-delayed-message"
	default:
		return "unknown"
	}
}

// DoneChannelResponse enables returning messages from a channel when completed
type DoneChannelResponse struct {
	Done    bool
	Error   error
	Message string
}

// RandomizeConnectionStrings shuffles connection strings.
// This can be used before calling Connect(ctx, ...connections) to randomly
// distribute connections to nodes in the cluster.
func RandomizeConnectionStrings(connections []string) []string {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(connections), func(i, j int) {
		connections[i], connections[j] = connections[j], connections[i]
	})
	return connections
}

// ConnectionString returns a connection URL string given a user, password, host, and scheme.
// It's a convenience function for net/url.Url.String().
func ConnectionString(user, password, host, scheme string) string {
	u := url.URL{
		Scheme: scheme,
		User:   url.UserPassword(user, password),
		Host:   host,
	}
	return u.String()
}

// GenerateConnectionStrings returns a list of randomized connection strings given
// a username, password, and list of hostname/IP:port strings.
func GenerateConnectionStrings(user string, password string, connections string) []string {
	cxns := make([]string, 0)
	for _, cxn := range strings.Split(connections, ",") {
		cxns = append(cxns, ConnectionString(user, password, cxn, "amqp"))
	}
	return RandomizeConnectionStrings(cxns)
}
