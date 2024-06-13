package errors

import "errors"

var (
	InvalidRequests     = errors.New("invalid requests")
	InternalServerError = errors.New("internal server error")

	FailedToCreateUser = errors.New("failed to create user")

	ErrEmailExists        = errors.New("email already exists")
	ErrInvalidCredentials = errors.New("invalid email or password")
)
