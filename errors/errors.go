package errors

import "errors"

var (
	ErrBadRequest error = errors.New("bad request")
	ErrNotFound   error = errors.New("not found")
	ErrInternal   error = errors.New("internal error")
	ErrNotCreated error = errors.New("not created")
)
