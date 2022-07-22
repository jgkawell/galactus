// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: core/eventstore/v1/eventstore.proto

package v1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// define the regex for a UUID once up-front
var _eventstore_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on CreateEventRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateEventRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateEventRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateEventRequestMultiError, or nil if none found.
func (m *CreateEventRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateEventRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetEvent()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CreateEventRequestValidationError{
					field:  "Event",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreateEventRequestValidationError{
					field:  "Event",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetEvent()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateEventRequestValidationError{
				field:  "Event",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CreateEventRequestMultiError(errors)
	}

	return nil
}

// CreateEventRequestMultiError is an error wrapping multiple validation errors
// returned by CreateEventRequest.ValidateAll() if the designated constraints
// aren't met.
type CreateEventRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateEventRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateEventRequestMultiError) AllErrors() []error { return m }

// CreateEventRequestValidationError is the validation error returned by
// CreateEventRequest.Validate if the designated constraints aren't met.
type CreateEventRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateEventRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateEventRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateEventRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateEventRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateEventRequestValidationError) ErrorName() string {
	return "CreateEventRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateEventRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateEventRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateEventRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateEventRequestValidationError{}

// Validate checks the field values on CreateEventResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateEventResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateEventResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateEventResponseMultiError, or nil if none found.
func (m *CreateEventResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateEventResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for IsPublished

	if len(errors) > 0 {
		return CreateEventResponseMultiError(errors)
	}

	return nil
}

// CreateEventResponseMultiError is an error wrapping multiple validation
// errors returned by CreateEventResponse.ValidateAll() if the designated
// constraints aren't met.
type CreateEventResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateEventResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateEventResponseMultiError) AllErrors() []error { return m }

// CreateEventResponseValidationError is the validation error returned by
// CreateEventResponse.Validate if the designated constraints aren't met.
type CreateEventResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateEventResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateEventResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateEventResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateEventResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateEventResponseValidationError) ErrorName() string {
	return "CreateEventResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreateEventResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateEventResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateEventResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateEventResponseValidationError{}

// Validate checks the field values on Event with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Event) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Event with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in EventMultiError, or nil if none found.
func (m *Event) ValidateAll() error {
	return m.validate(true)
}

func (m *Event) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetEventId() != "" {

		if err := m._validateUuid(m.GetEventId()); err != nil {
			err = EventValidationError{
				field:  "EventId",
				reason: "value must be a valid UUID",
				cause:  err,
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if all {
		switch v := interface{}(m.GetReceivedDate()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, EventValidationError{
					field:  "ReceivedDate",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, EventValidationError{
					field:  "ReceivedDate",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetReceivedDate()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return EventValidationError{
				field:  "ReceivedDate",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetPublishedDate()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, EventValidationError{
					field:  "PublishedDate",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, EventValidationError{
					field:  "PublishedDate",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetPublishedDate()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return EventValidationError{
				field:  "PublishedDate",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if err := m._validateUuid(m.GetTransactionId()); err != nil {
		err = EventValidationError{
			field:  "TransactionId",
			reason: "value must be a valid UUID",
			cause:  err,
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Publish

	// no validation rules for AggregateType

	// no validation rules for EventType

	if err := m._validateUuid(m.GetAggregateId()); err != nil {
		err = EventValidationError{
			field:  "AggregateId",
			reason: "value must be a valid UUID",
			cause:  err,
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for EventData

	if len(errors) > 0 {
		return EventMultiError(errors)
	}

	return nil
}

func (m *Event) _validateUuid(uuid string) error {
	if matched := _eventstore_uuidPattern.MatchString(uuid); !matched {
		return errors.New("invalid uuid format")
	}

	return nil
}

// EventMultiError is an error wrapping multiple validation errors returned by
// Event.ValidateAll() if the designated constraints aren't met.
type EventMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m EventMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m EventMultiError) AllErrors() []error { return m }

// EventValidationError is the validation error returned by Event.Validate if
// the designated constraints aren't met.
type EventValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EventValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EventValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EventValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EventValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EventValidationError) ErrorName() string { return "EventValidationError" }

// Error satisfies the builtin error interface
func (e EventValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEvent.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EventValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EventValidationError{}
