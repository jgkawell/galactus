package messagebus

import "context"

// ClientHandler must be implemented by the client to receive messages.
type ClientHandler interface {
	// Handle is responsible for processing messages from a queue/topic.
	// NOTE: Consumers SHOULD ALWAYS call Complete/Retry/Reject on the QueuedMessage
	//       prior to returning. However, if an error is returned from this function without
	//       having called Complete/Retry/Reject on the msg, then the msg will be automatically
	//       Retried so that it is redelivered.  If this is not the desired behavior, then
	//       explictly call Reject so that the message is discarded.
	Handle(ctx context.Context, msg Message) error
}
