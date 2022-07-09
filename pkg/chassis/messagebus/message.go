package messagebus

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	jsonpb "google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// A Message is received by the ClientMessageHandler's Handle method.
type Message interface {
	// MessageID returns the unique ID for this message.
	MessageID() string
	// MessageType returns the (short) name of the message type that was sent in the message.
	MessageType() string
	// Timestamp returns the time the original message was sent.
	Timestamp() time.Time
	// GetMessage populates the provided object with the message received from the queue.
	GetMessage(target interface{}) error
	// Complete indicates that the message has been received and successfully processed (or is being ignored), and can be deleted from the queue/topic.
	Complete() error
	// Retry sends the message back to the queue/topic for redelivery to this service.
	Retry() error
	// RetryAfter Sends the message back to the queue/topic for redelivery to this service after a specified time.
	RetryAfter(seconds int) error
	// Reject notifies the queue that the service had a problem processing the message, and doesn't want it redelivered.
	Reject() error
	// GetCustomHeader returns the value of a custom header set with MessageOptionCustom(name, value) and a bool indicating whether the
	// header exists on the message or not.
	GetCustomHeader(name string) (string, bool)
}

type message struct {
	acked bool
	msg   *amqp.Delivery
}

func serializeProtoMessage(p proto.Message) ([]byte, error) {
	data, err := jsonpb.Marshal(p)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal the proto message to json: %#+v: %+v", p, err)
	}
	return data, nil
}

func serializeMessage(val interface{}) ([]byte, error) {
	if val == nil {
		return nil, errors.New("provided interface is nil")
	}
	if v, ok := val.(proto.Message); ok {
		return serializeProtoMessage(v)
	} else {
		rv := reflect.ValueOf(val)
		if rv.Kind() != reflect.Ptr {
			// rv isn't addressable, so can't get its address directly
			n := reflect.New(rv.Type()) // create a new *Val
			n.Elem().Set(rv)            // Set *n = val
			val = n.Interface()
		}
		if v, ok := val.(proto.Message); ok {
			return serializeProtoMessage(v)
		}
		return json.Marshal(val)
	}
}

func deserializeMessage(data []byte, target interface{}) error {
	if target == nil {
		return errors.New("provided interface is nil")
	}

	if reflect.TypeOf(target).Kind() != reflect.Ptr {
		return errors.New("provided interface must be a pointer")
	}

	switch t := target.(type) {
	case proto.Message:
		return jsonpb.Unmarshal(data, t)
	default:
		if err := json.Unmarshal(data, target); err != nil {
			return err
		}
		return nil
	}
}

func getMessageType(val interface{}) string {
	if val == nil {
		return "nil"
	}
	if reflect.TypeOf(val).Kind() == reflect.Ptr {
		v := reflect.ValueOf(val)
		e := v.Elem()
		t := e.Type()
		return t.Name()
	}
	return reflect.TypeOf(val).Name()
}

// Timestamp indicates when the message was sent.
func (qm *message) Timestamp() time.Time {
	return qm.msg.Timestamp
}

// MessageID is the uniqueID for the sent message.
func (qm *message) MessageID() string {
	return qm.msg.MessageId
}

// MessageType is the name of the sent message type.
func (qm *message) MessageType() string {
	return qm.msg.Type
}

// GetMessage deserializes the message into a variable of the expected type.
func (qm *message) GetMessage(target interface{}) error {
	type1 := getMessageType(target)
	type2 := qm.msg.Type
	if type1 != type2 {
		return errors.New("provided message type does not match the serialized message")
	}
	err := deserializeMessage(qm.msg.Body, target)
	if err != nil {
		return err
	}

	return nil
}

// Complete accepts the message, and removes it from the queue.
func (qm *message) Complete() error {
	if qm.acked {
		return nil
	}
	qm.acked = true
	return qm.msg.Ack(false)
}

// Complete accepts the message, and removes it from the queue.
func (qm *message) Retry() error {
	return qm.reject(0, true)
}

// Complete accepts the message, and removes it from the queue.
func (qm *message) RetryAfter(seconds int) error {
	return qm.reject(seconds, true)
}

// Reject  accepts the message, and removes it from the queue.
func (qm *message) Reject() error {
	return qm.reject(0, false)
}

// Reject send the message back to the queue for retransmission.
func (qm *message) reject(retryDelay int, retransmit bool) error {
	if qm.acked {
		return nil
	}
	qm.acked = true

	if !retransmit {
		// dead-letter
		return qm.msg.Nack(false, false)
	}

	// check for retries, dead-letter if they've been exceeded
	val, ok := qm.msg.Headers[messageOptionRetryCount]
	if ok {
		retryVal := fmt.Sprintf("%s", val)
		remainingRetries, err := strconv.Atoi(retryVal)
		if err != nil {
			return err
		}
		if remainingRetries <= 0 {
			return qm.msg.Nack(false, false)
		}

		remainingRetries--
		qm.msg.Headers[messageOptionRetryCount] = strconv.Itoa(remainingRetries)
	}

	delay, ok := qm.msg.Headers[messageOptionRetryDelaySeconds]
	if ok {
		sdelay := fmt.Sprintf("%s", delay)
		idelay, err := strconv.Atoi(sdelay)
		if err != nil {
			return err
		}
		qm.msg.Headers[messageOptionDeliverAfterMilliseconds] = strconv.Itoa(idelay)
	}
	if retryDelay > 0 {
		qm.msg.Headers[messageOptionDeliverAfterMilliseconds] = strconv.Itoa(retryDelay)
	}
	// normal retrans, no header changes
	return qm.msg.Nack(false, true)
}

func (qm *message) GetCustomHeader(name string) (string, bool) {
	name = fmt.Sprintf("%s%s", customHeaderPrefix, name)
	value, ok := qm.msg.Headers[name]
	if !ok {
		return "", false
	}

	strval, ok := value.(string)
	if !ok {
		return "", false
	}
	return strval, ok
}
