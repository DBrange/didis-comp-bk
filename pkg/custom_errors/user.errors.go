package customerrors

import "errors"

var (
	ErrUserNotFound        = errors.New("user not found")
	ErrUserInsertionFailed = errors.New("user insertion failed")
	ErrUserInvalidID       = errors.New("invalid ID format")
	ErrUserDuplicateKey    = errors.New("duplicate key error")
	ErrUserUpdated          = errors.New("user not updated")
	ErrUserDeleted          = errors.New("user not deleted")
)

var (
	ErrValidationFailed = errors.New("validation failed")
	ErrConnectionFailed = errors.New("connection failed")
	ErrSchemaViolation  = errors.New("schema violation")
)

var (
	ErrLocationNotFound        = errors.New("location not found")
	ErrLocationInsertionFailed = errors.New("location insertion failed")
	ErrLocationInvalidID       = errors.New("invalid ID format")
	ErrLocationDuplicateKey    = errors.New("duplicate key error")
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
	ErrCodeUpdated          = "could_not_be_updated"
	ErrCodeDeleted          = "could_not_be_deleted"
)
