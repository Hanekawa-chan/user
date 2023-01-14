package errors

import "errors"

var (
	ErrInternal   = errors.New("internal error")
	ErrValidation = errors.New("validation error")
)
