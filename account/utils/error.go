package utils

import (
	"errors"
	"fmt"
	"net/http"
)

//type holds a type string and integer code for the error
type Type string

const (
	Authorization Type = "AUTHORIZATION"
	BadRequest    Type = "BAD_REQUEST"
	Conflict	  Type = "CONFLICT"
	Internal	  Type = "INTERNAL"
	NotFound	  Type = "NOT_FOUND"
	PayloadTooLarge Type = "PAYLOAD_TOO_LARGE"
)

type Error struct {
	Type	Type   `json:"type"`
	Message string `json:"message"`
}


func (e *Error) Error() string {
	return e.Message;
}

func (e *Error) Status() int {
	switch e.Type{
		case Authorization:
			return http.StatusUnauthorized;
		case BadRequest:
			return http.StatusBadRequest;
		case Conflict:
			return http.StatusConflict;
		case Internal:
			return http.StatusInternalServerError;
		case NotFound:
			return http.StatusNotFound;
		case PayloadTooLarge:
			return http.StatusRequestEntityTooLarge;
		default:
			return http.StatusInternalServerError;
	}
}

func Status (err error) int {
	var e *Error;
	if errors.As(err, &e){ //if err is an instance of Error
		return e.Status(); //return the status code of the error
	}
	return http.StatusInternalServerError;
}

//Error Factories
func NewAuthorization(reason string) *Error {
	return &Error{
		Type: Authorization,
		Message: reason,
	}
}

func NewBadRequest(reason string) *Error {
	return &Error{
		Type: BadRequest,
		Message: "Bad request: " + reason,
	}
}

func NewConflict(name string, value string) *Error {
	return &Error{
		Type: Conflict,
		Message: fmt.Sprintf("resource: %v with value: %v already exists", name, value),
    }
}

func NewInternal() *Error {
	return &Error{
		Type: Internal,
		Message: "Internal server error",
	}
}

func NewNotFound(name string, value string) *Error {
	return &Error{
		Type: NotFound,
		Message: fmt.Sprintf("resource: %v with value: %v not found", name, value),
	}
}

func NewPayloadTooLarge(maxSize int64) *Error {
	return &Error{
		Type: PayloadTooLarge,
		Message: fmt.Sprintf("payload size limit exceeded. Max payload size is %v bytes", maxSize),
	}
}