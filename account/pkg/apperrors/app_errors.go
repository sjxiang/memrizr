package apperrors

import (
	"errors"
	"fmt"
	"net/http"
)

type Type string


const (
	Authorization   Type = "AUTHORIZATION"
	BadRequest      Type = "BADREQUEST"
	Conflict        Type = "CONFLICT"
	Internal        Type = "INTERNAL"
	NoFound         Type = "NOTFOUND"
	PayloadTooLarge Type = "PAYLOADTOOLARGE"
)

type Error struct {
	Type    Type    `json:"type"`
	Message string  `json:"message"`
}


func (e *Error) Status() int {
	switch e.Type {
	case Authorization:
		return http.StatusUnauthorized
	case BadRequest:
		return http.StatusBadRequest
	case Conflict:
		return http.StatusConflict
	case Internal:
		return http.StatusInternalServerError
	case NoFound:
		return http.StatusNotFound
	case PayloadTooLarge:
		return http.StatusRequestEntityTooLarge
	default:
		return http.StatusInternalServerError
	}
}

func Status(err error) int {
	var e *Error
	if errors.As(err, &e) {
		return e.Status()
	}
	return http.StatusInternalServerError
}



/*
 * Error 工厂
 */


//  NewAuthorization 401
func NewAuthorization(reason string) *Error {
	return &Error{
		Type: Authorization,
		Message: reason,
	}
}

// NewBadRequest 400 (e.g. 校验)
func NewBadRequest(reason string) *Error {
	return &Error{
		Type: BadRequest,
		Message: fmt.Sprintf("Bad request. Reason: %v", reason),
	}
}

// NewConflict 409
func NewConflict(name, value string) *Error {
	return &Error{
		Type: Conflict,
		Message: fmt.Sprintf("resource: %v with value: %v already exists", name, value),
	}
}

// NewInternal 500
func NewInternal() *Error {
	return &Error{
		Type: Internal,
		Message: fmt.Sprintf("Internal server error"),
	}
}

// NewNotFound 404
func NewNotFound(name, value string) *Error {
	return &Error{
		Type: NoFound,
		Message: fmt.Sprintf("resource: %v with value: %v already exists", name, value),
	}
}

// NewPayloadTooLarge 413
func NewPayloadTooLarge(maxBodySize, contentLength int64) *Error {
	return &Error{
		Type: PayloadTooLarge,
		Message: fmt.Sprintf("Max payload size of %v exceeded. Actual payload size: %v", maxBodySize, contentLength),
	}
}