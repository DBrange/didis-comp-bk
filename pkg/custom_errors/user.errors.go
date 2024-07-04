package customerrors

import "errors"

var (
	ErrUserNotFound        = errors.New("user not found")
	ErrUserInsertionFailed = errors.New("user insertion failed")
	ErrUserInvalidID       = errors.New("invalid ID format")
	ErrUserDuplicateKey    = errors.New("duplicate key error")
	ErrValidationFailed    = errors.New("validation failed")
	ErrConnectionFailed    = errors.New("connection failed")
	ErrSchemaViolation     = errors.New("schema violation")
)

const (
	ErrCodeNotFound         = "not_found"
	ErrCodeCouldNotBeAdded  = "could_not_be_added"
	ErrCodeInsertionFailed  = "could_not_be_inserted"
	ErrCodeDuplicateKey     = "duplicate_value"
	ErrCodeInvalidID        = "invalid_id"
	ErrCodeValidationFailed = "validation_failed"
	ErrCodeConnectionFailed = "connection_failed"
	ErrCodeSchemaViolation  = "schema_violation"
)
