// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: v1/account.proto

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

// Validate checks the field values on SignupRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *SignupRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SignupRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in SignupRequestMultiError, or
// nil if none found.
func (m *SignupRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *SignupRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Email

	// no validation rules for Password

	if len(errors) > 0 {
		return SignupRequestMultiError(errors)
	}

	return nil
}

// SignupRequestMultiError is an error wrapping multiple validation errors
// returned by SignupRequest.ValidateAll() if the designated constraints
// aren't met.
type SignupRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SignupRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SignupRequestMultiError) AllErrors() []error { return m }

// SignupRequestValidationError is the validation error returned by
// SignupRequest.Validate if the designated constraints aren't met.
type SignupRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SignupRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SignupRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SignupRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SignupRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SignupRequestValidationError) ErrorName() string { return "SignupRequestValidationError" }

// Error satisfies the builtin error interface
func (e SignupRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSignupRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SignupRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SignupRequestValidationError{}

// Validate checks the field values on SignupResponse with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *SignupResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SignupResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in SignupResponseMultiError,
// or nil if none found.
func (m *SignupResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *SignupResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Message

	if len(errors) > 0 {
		return SignupResponseMultiError(errors)
	}

	return nil
}

// SignupResponseMultiError is an error wrapping multiple validation errors
// returned by SignupResponse.ValidateAll() if the designated constraints
// aren't met.
type SignupResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SignupResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SignupResponseMultiError) AllErrors() []error { return m }

// SignupResponseValidationError is the validation error returned by
// SignupResponse.Validate if the designated constraints aren't met.
type SignupResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SignupResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SignupResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SignupResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SignupResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SignupResponseValidationError) ErrorName() string { return "SignupResponseValidationError" }

// Error satisfies the builtin error interface
func (e SignupResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSignupResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SignupResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SignupResponseValidationError{}

// Validate checks the field values on VerifyRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *VerifyRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on VerifyRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in VerifyRequestMultiError, or
// nil if none found.
func (m *VerifyRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *VerifyRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Token

	if len(errors) > 0 {
		return VerifyRequestMultiError(errors)
	}

	return nil
}

// VerifyRequestMultiError is an error wrapping multiple validation errors
// returned by VerifyRequest.ValidateAll() if the designated constraints
// aren't met.
type VerifyRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m VerifyRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m VerifyRequestMultiError) AllErrors() []error { return m }

// VerifyRequestValidationError is the validation error returned by
// VerifyRequest.Validate if the designated constraints aren't met.
type VerifyRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e VerifyRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e VerifyRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e VerifyRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e VerifyRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e VerifyRequestValidationError) ErrorName() string { return "VerifyRequestValidationError" }

// Error satisfies the builtin error interface
func (e VerifyRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sVerifyRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = VerifyRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = VerifyRequestValidationError{}

// Validate checks the field values on VerifyResponse with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *VerifyResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on VerifyResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in VerifyResponseMultiError,
// or nil if none found.
func (m *VerifyResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *VerifyResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Message

	if len(errors) > 0 {
		return VerifyResponseMultiError(errors)
	}

	return nil
}

// VerifyResponseMultiError is an error wrapping multiple validation errors
// returned by VerifyResponse.ValidateAll() if the designated constraints
// aren't met.
type VerifyResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m VerifyResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m VerifyResponseMultiError) AllErrors() []error { return m }

// VerifyResponseValidationError is the validation error returned by
// VerifyResponse.Validate if the designated constraints aren't met.
type VerifyResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e VerifyResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e VerifyResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e VerifyResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e VerifyResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e VerifyResponseValidationError) ErrorName() string { return "VerifyResponseValidationError" }

// Error satisfies the builtin error interface
func (e VerifyResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sVerifyResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = VerifyResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = VerifyResponseValidationError{}

// Validate checks the field values on LoginRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *LoginRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on LoginRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in LoginRequestMultiError, or
// nil if none found.
func (m *LoginRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *LoginRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Email

	// no validation rules for Password

	if len(errors) > 0 {
		return LoginRequestMultiError(errors)
	}

	return nil
}

// LoginRequestMultiError is an error wrapping multiple validation errors
// returned by LoginRequest.ValidateAll() if the designated constraints aren't met.
type LoginRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m LoginRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m LoginRequestMultiError) AllErrors() []error { return m }

// LoginRequestValidationError is the validation error returned by
// LoginRequest.Validate if the designated constraints aren't met.
type LoginRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LoginRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LoginRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LoginRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LoginRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LoginRequestValidationError) ErrorName() string { return "LoginRequestValidationError" }

// Error satisfies the builtin error interface
func (e LoginRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLoginRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LoginRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LoginRequestValidationError{}

// Validate checks the field values on LoginResponse with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *LoginResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on LoginResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in LoginResponseMultiError, or
// nil if none found.
func (m *LoginResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *LoginResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Token

	if len(errors) > 0 {
		return LoginResponseMultiError(errors)
	}

	return nil
}

// LoginResponseMultiError is an error wrapping multiple validation errors
// returned by LoginResponse.ValidateAll() if the designated constraints
// aren't met.
type LoginResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m LoginResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m LoginResponseMultiError) AllErrors() []error { return m }

// LoginResponseValidationError is the validation error returned by
// LoginResponse.Validate if the designated constraints aren't met.
type LoginResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LoginResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LoginResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LoginResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LoginResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LoginResponseValidationError) ErrorName() string { return "LoginResponseValidationError" }

// Error satisfies the builtin error interface
func (e LoginResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLoginResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LoginResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LoginResponseValidationError{}
