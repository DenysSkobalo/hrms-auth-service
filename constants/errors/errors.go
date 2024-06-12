package errors

import "errors"

var (
	InvalidRequests     = errors.New("invalid requests")
	InternalServerError = errors.New("internal server error")
	ErrEmailExists      = errors.New("email already exists")
	FailedToCreateUser  = errors.New("failed to create user")
)
