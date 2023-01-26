package app

import "errors"

var (
	ErrNotFound   = errors.New("rows not found")
	ErrInternal   = errors.New("internal error")
	ErrValidation = errors.New("validation error")
)
