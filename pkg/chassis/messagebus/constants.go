package messagebus

import (
	"time"
)

const qNameDivider = "."

const rabbitMqTTLKey = "x-message-ttl"
const rabbitMqDeadLetterExchangeKey = "x-dead-letter-exchange"
const rabbitMqDeadLetterRoutingKey = "x-dead-letter-routing-key"
const rabbitMqDelayedExchangeArg = "x-delayed-type"
const rabbitMqDelayedExchangeType = "x-delayed-message"
const deadLetterQueueName = "dead-letter"
const serviceQueueName = "queue"

// NOTE: if you change this, you'll need to delete the dead-letter queue before deploying a service
const serviceDeadLetter = 24 * time.Hour

const (
	customHeaderPrefix = "x-custom-"
	customCtxPrefix    = "x-custom-ctx-"
	// messageOptionPriority set the message priority from 0 to 9 (default: 0).
	messageOptionPriority = "x-custom-send-priority"
	// messageOptionDeliverAfterMilliseconds sets the delay before delivering a message.
	messageOptionDeliverAfterMilliseconds = "x-delay"
	// messageOptionRetryCount controls the number of times the message will be delivered if rejected by a recipient - default is "0" (indefinite).
	messageOptionRetryCount = "x-custom-topic-retrycount"
	// messageOptionRetryDelaySeconds controls the number of seconds the message will be delayed before retry if retry count is set - default is "0" (immediate).
	messageOptionRetryDelaySeconds = "x-custom-topic-retrydelay"
	// messageOptionTTLSeconds sets the expiration time for the message, which is checked by the receiving queue.
	messageOptionTTLSeconds = "x-custom-topic-ttl"
	// messageOptionId sets the message id for the message. The value here is just a sentinel value used by the
	// library to set the MessageId property on the message.  There is no x-custom-message-id header added to the message.
	messageOptionId = "x-custom-message-id"
	// serviceOptionDedupeSeconds activates/deactivates service-level deduplication with the specified duration.
	serviceOptionDedupeSeconds = "x-service-dedupe-seconds"
	// serviceOptionDelayIntervalSeconds sets how often to check if the message should be delivered yet, for messages with delayedDelivery.
	serviceOptionDelayIntervalSeconds = "x-service-delay-increment"
	// registerOptionHaModeAll is a sentinel value that creates a policy on a queue with an HA mode "all" set.
	registerOptionHaModeAll = "x-ha-all"
	// registerOptionHaModeAll is a sentinel value that creates a policy on a queue with an HA mode "exactly" set.
	registerOptionHaModeExactly = "x-ha-exactly"
	// registerOptionHaModeAll is a sentinel value that creates a policy on a queue with an HA mode "nodes" set.
	registerOptionHaModeNodes = "x-ha-nodes"
)
