package storage

import "errors"

// Common errors
var (
	ErrNotFound         = errors.New("not found")
	ErrAlreadyExists    = errors.New("already exists")
	ErrInvalidInput     = errors.New("invalid input")
	ErrPermissionDenied = errors.New("permission denied")
	ErrUnauthorized     = errors.New("unauthorized")
	ErrInternalServer   = errors.New("internal server error")
	ErrConflict         = errors.New("conflict")
)

// Database constraint errors
var (
	ErrDuplicateKey        = errors.New("duplicate key violation")
	ErrForeignKeyViolation = errors.New("foreign key violation")
	ErrCheckConstraint     = errors.New("check constraint violation")
	ErrNotNullConstraint   = errors.New("not null constraint violation")
	ErrUniqueConstraint    = errors.New("unique constraint violation")
)
