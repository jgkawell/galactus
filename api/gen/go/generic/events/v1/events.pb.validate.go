// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: generic/events/v1/events.proto

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

// Validate checks the field values on EventType with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *EventType) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on EventType with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in EventTypeMultiError, or nil
// if none found.
func (m *EventType) ValidateAll() error {
	return m.validate(true)
}

func (m *EventType) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	switch v := m.Code.(type) {
	case *EventType_SystemCode:
		if v == nil {
			err := EventTypeValidationError{
				field:  "Code",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		// no validation rules for SystemCode
	case *EventType_NotificationCode:
		if v == nil {
			err := EventTypeValidationError{
				field:  "Code",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		// no validation rules for NotificationCode
	case *EventType_TodoEventCode:
		if v == nil {
			err := EventTypeValidationError{
				field:  "Code",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		// no validation rules for TodoEventCode
	default:
		_ = v // ensures v is used
	}

	if len(errors) > 0 {
		return EventTypeMultiError(errors)
	}

	return nil
}

// EventTypeMultiError is an error wrapping multiple validation errors returned
// by EventType.ValidateAll() if the designated constraints aren't met.
type EventTypeMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m EventTypeMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m EventTypeMultiError) AllErrors() []error { return m }

// EventTypeValidationError is the validation error returned by
// EventType.Validate if the designated constraints aren't met.
type EventTypeValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EventTypeValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EventTypeValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EventTypeValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EventTypeValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EventTypeValidationError) ErrorName() string { return "EventTypeValidationError" }

// Error satisfies the builtin error interface
func (e EventTypeValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEventType.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EventTypeValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EventTypeValidationError{}
