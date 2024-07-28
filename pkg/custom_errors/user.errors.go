package customerrors

import "errors"

// Errors
var (
	ErrNotFound            = errors.New("not found")
	ErrInsertionFailed     = errors.New("insertion failed")
	ErrInvalidID           = errors.New("invalid ID format")
	ErrDuplicateKey        = errors.New("duplicate key error")
	ErrSchemaViolation     = errors.New("scheme violation")
	ErrUpdated             = errors.New("not updated")
	ErrDeleted             = errors.New("not deleted")
	ErrValidationFailed    = errors.New("validation failed")
	ErrGetJSON             = errors.New("the json was not obtained")
	ErrTransaction         = errors.New("transacion failed")
	ErrStartSessionFailed  = errors.New("start of session failed")
	ErrAuthorizationHeader = errors.New("authorization header is required")
	ErrHashedFailed        = errors.New("hashed failed")
	ErrTokenSigned         = errors.New("signed failed")
	ErrSingedMethod        = errors.New("unexpected signing method")
	ErrComparedHash        = errors.New("camparison failed")
	ErrAuthorization       = errors.New("authorization failed")
	// ErrConnectionFailed   = errors.New("connection failed")
)

// Error codes
const (
	ErrCodeNotFound            = "not_found"
	ErrCodeInsertionFailed     = "could_not_be_inserted"
	ErrCodeInvalidID           = "invalid_id"
	ErrCodeDuplicateKey        = "duplicate_value"
	ErrCodeValidationFailed    = "validation_failed"
	ErrCodeSchemaViolation     = "schema_violation"
	ErrCodeUpdated             = "could_not_be_updated"
	ErrCodeDeleted             = "could_not_be_deleted"
	ErrCodeGetJSON             = "could_not_get_the_json"
	ErrCodeTransaction         = "transacion_failed"
	ErrCodeStartSessionFailed  = "start_of_session_failed"
	ErrCodeAuthorizationHeader = "authorization_header_not_found"
	ErrCodeHashedFailed        = "hashed_failed"
	ErrCodeTokenSigned         = "signed_failed"
	ErrCodeSingedMethod        = "unexpected_signing_method"
	ErrCodeComparedHash        = "camparison_failed"
	ErrCodeAuthorization       = "authorized_failed"
	// ErrCodeConnectionFailed   = "connection_failed"
)
