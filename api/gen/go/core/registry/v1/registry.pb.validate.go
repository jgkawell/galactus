// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: core/registry/v1/registry.proto

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

	v1 "github.com/jgkawell/galactus/api/gen/go/core/aggregates/v1"
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

	_ = v1.ProtocolKind(0)
)

// Validate checks the field values on RegisterRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *RegisterRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RegisterRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// RegisterRequestMultiError, or nil if none found.
func (m *RegisterRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *RegisterRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	// no validation rules for Domain

	// no validation rules for Version

	// no validation rules for Description

	for idx, item := range m.GetProtocols() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, RegisterRequestValidationError{
						field:  fmt.Sprintf("Protocols[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, RegisterRequestValidationError{
						field:  fmt.Sprintf("Protocols[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return RegisterRequestValidationError{
					field:  fmt.Sprintf("Protocols[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	for idx, item := range m.GetConsumers() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, RegisterRequestValidationError{
						field:  fmt.Sprintf("Consumers[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, RegisterRequestValidationError{
						field:  fmt.Sprintf("Consumers[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return RegisterRequestValidationError{
					field:  fmt.Sprintf("Consumers[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return RegisterRequestMultiError(errors)
	}

	return nil
}

// RegisterRequestMultiError is an error wrapping multiple validation errors
// returned by RegisterRequest.ValidateAll() if the designated constraints
// aren't met.
type RegisterRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RegisterRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RegisterRequestMultiError) AllErrors() []error { return m }

// RegisterRequestValidationError is the validation error returned by
// RegisterRequest.Validate if the designated constraints aren't met.
type RegisterRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RegisterRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RegisterRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RegisterRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RegisterRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RegisterRequestValidationError) ErrorName() string { return "RegisterRequestValidationError" }

// Error satisfies the builtin error interface
func (e RegisterRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRegisterRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RegisterRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RegisterRequestValidationError{}

// Validate checks the field values on ProtocolRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ProtocolRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ProtocolRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ProtocolRequestMultiError, or nil if none found.
func (m *ProtocolRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ProtocolRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Kind

	// no validation rules for Route

	if len(errors) > 0 {
		return ProtocolRequestMultiError(errors)
	}

	return nil
}

// ProtocolRequestMultiError is an error wrapping multiple validation errors
// returned by ProtocolRequest.ValidateAll() if the designated constraints
// aren't met.
type ProtocolRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ProtocolRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ProtocolRequestMultiError) AllErrors() []error { return m }

// ProtocolRequestValidationError is the validation error returned by
// ProtocolRequest.Validate if the designated constraints aren't met.
type ProtocolRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ProtocolRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ProtocolRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ProtocolRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ProtocolRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ProtocolRequestValidationError) ErrorName() string { return "ProtocolRequestValidationError" }

// Error satisfies the builtin error interface
func (e ProtocolRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProtocolRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ProtocolRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ProtocolRequestValidationError{}

// Validate checks the field values on ConsumerRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ConsumerRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ConsumerRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ConsumerRequestMultiError, or nil if none found.
func (m *ConsumerRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ConsumerRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Kind

	// no validation rules for Order

	// no validation rules for AggregateType

	// no validation rules for EventType

	// no validation rules for EventCode

	if len(errors) > 0 {
		return ConsumerRequestMultiError(errors)
	}

	return nil
}

// ConsumerRequestMultiError is an error wrapping multiple validation errors
// returned by ConsumerRequest.ValidateAll() if the designated constraints
// aren't met.
type ConsumerRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ConsumerRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ConsumerRequestMultiError) AllErrors() []error { return m }

// ConsumerRequestValidationError is the validation error returned by
// ConsumerRequest.Validate if the designated constraints aren't met.
type ConsumerRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ConsumerRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ConsumerRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ConsumerRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ConsumerRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ConsumerRequestValidationError) ErrorName() string { return "ConsumerRequestValidationError" }

// Error satisfies the builtin error interface
func (e ConsumerRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sConsumerRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ConsumerRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ConsumerRequestValidationError{}

// Validate checks the field values on RegisterResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *RegisterResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RegisterResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// RegisterResponseMultiError, or nil if none found.
func (m *RegisterResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *RegisterResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetProtocols() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, RegisterResponseValidationError{
						field:  fmt.Sprintf("Protocols[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, RegisterResponseValidationError{
						field:  fmt.Sprintf("Protocols[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return RegisterResponseValidationError{
					field:  fmt.Sprintf("Protocols[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	for idx, item := range m.GetConsumers() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, RegisterResponseValidationError{
						field:  fmt.Sprintf("Consumers[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, RegisterResponseValidationError{
						field:  fmt.Sprintf("Consumers[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return RegisterResponseValidationError{
					field:  fmt.Sprintf("Consumers[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return RegisterResponseMultiError(errors)
	}

	return nil
}

// RegisterResponseMultiError is an error wrapping multiple validation errors
// returned by RegisterResponse.ValidateAll() if the designated constraints
// aren't met.
type RegisterResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RegisterResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RegisterResponseMultiError) AllErrors() []error { return m }

// RegisterResponseValidationError is the validation error returned by
// RegisterResponse.Validate if the designated constraints aren't met.
type RegisterResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RegisterResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RegisterResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RegisterResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RegisterResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RegisterResponseValidationError) ErrorName() string { return "RegisterResponseValidationError" }

// Error satisfies the builtin error interface
func (e RegisterResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRegisterResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RegisterResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RegisterResponseValidationError{}

// Validate checks the field values on ProtocolResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ProtocolResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ProtocolResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ProtocolResponseMultiError, or nil if none found.
func (m *ProtocolResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ProtocolResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Kind

	// no validation rules for Port

	if len(errors) > 0 {
		return ProtocolResponseMultiError(errors)
	}

	return nil
}

// ProtocolResponseMultiError is an error wrapping multiple validation errors
// returned by ProtocolResponse.ValidateAll() if the designated constraints
// aren't met.
type ProtocolResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ProtocolResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ProtocolResponseMultiError) AllErrors() []error { return m }

// ProtocolResponseValidationError is the validation error returned by
// ProtocolResponse.Validate if the designated constraints aren't met.
type ProtocolResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ProtocolResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ProtocolResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ProtocolResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ProtocolResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ProtocolResponseValidationError) ErrorName() string { return "ProtocolResponseValidationError" }

// Error satisfies the builtin error interface
func (e ProtocolResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProtocolResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ProtocolResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ProtocolResponseValidationError{}

// Validate checks the field values on ConsumerResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ConsumerResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ConsumerResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ConsumerResponseMultiError, or nil if none found.
func (m *ConsumerResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ConsumerResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Kind

	// no validation rules for Order

	// no validation rules for RoutingKey

	// no validation rules for Exchange

	// no validation rules for QueueName

	if len(errors) > 0 {
		return ConsumerResponseMultiError(errors)
	}

	return nil
}

// ConsumerResponseMultiError is an error wrapping multiple validation errors
// returned by ConsumerResponse.ValidateAll() if the designated constraints
// aren't met.
type ConsumerResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ConsumerResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ConsumerResponseMultiError) AllErrors() []error { return m }

// ConsumerResponseValidationError is the validation error returned by
// ConsumerResponse.Validate if the designated constraints aren't met.
type ConsumerResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ConsumerResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ConsumerResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ConsumerResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ConsumerResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ConsumerResponseValidationError) ErrorName() string { return "ConsumerResponseValidationError" }

// Error satisfies the builtin error interface
func (e ConsumerResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sConsumerResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ConsumerResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ConsumerResponseValidationError{}

// Validate checks the field values on ConnectionRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ConnectionRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ConnectionRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ConnectionRequestMultiError, or nil if none found.
func (m *ConnectionRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ConnectionRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	// no validation rules for Version

	// no validation rules for Type

	if len(errors) > 0 {
		return ConnectionRequestMultiError(errors)
	}

	return nil
}

// ConnectionRequestMultiError is an error wrapping multiple validation errors
// returned by ConnectionRequest.ValidateAll() if the designated constraints
// aren't met.
type ConnectionRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ConnectionRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ConnectionRequestMultiError) AllErrors() []error { return m }

// ConnectionRequestValidationError is the validation error returned by
// ConnectionRequest.Validate if the designated constraints aren't met.
type ConnectionRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ConnectionRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ConnectionRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ConnectionRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ConnectionRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ConnectionRequestValidationError) ErrorName() string {
	return "ConnectionRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ConnectionRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sConnectionRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ConnectionRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ConnectionRequestValidationError{}

// Validate checks the field values on ConnectionResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ConnectionResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ConnectionResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ConnectionResponseMultiError, or nil if none found.
func (m *ConnectionResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ConnectionResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Address

	// no validation rules for Port

	// no validation rules for Status

	if len(errors) > 0 {
		return ConnectionResponseMultiError(errors)
	}

	return nil
}

// ConnectionResponseMultiError is an error wrapping multiple validation errors
// returned by ConnectionResponse.ValidateAll() if the designated constraints
// aren't met.
type ConnectionResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ConnectionResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ConnectionResponseMultiError) AllErrors() []error { return m }

// ConnectionResponseValidationError is the validation error returned by
// ConnectionResponse.Validate if the designated constraints aren't met.
type ConnectionResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ConnectionResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ConnectionResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ConnectionResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ConnectionResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ConnectionResponseValidationError) ErrorName() string {
	return "ConnectionResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ConnectionResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sConnectionResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ConnectionResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ConnectionResponseValidationError{}
