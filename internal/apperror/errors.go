package apperror

import "errors"

var ErrNotFound = errors.New("not found")
var ErrForbidden = errors.New("forbidden")
var ErrInvalidInput = errors.New("invalid input")
var ErrInternal = errors.New("internal error")
