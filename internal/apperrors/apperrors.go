package apperrors

import (
	"errors"
	"net/http"
)

type Type string

const (
	BadRequest Type = "BAD_REQUEST"
	Conflict   Type = "CONFLICT"
	Internal   Type = "INTERNAL"
	NotFound   Type = "NOT_FOUND"
)

type Error struct {
	Message string
	Type    Type
}

func NewError(msg string, code Type) *Error {
	return &Error{
		Message: msg,
		Type:    code,
	}
}

func (err *Error) Error() string {
	return err.Message
}

func (err *Error) Status() int {
	switch err.Type {
	case BadRequest:
		return http.StatusBadRequest
	case Conflict:
		return http.StatusConflict
	case Internal:
		return http.StatusInternalServerError
	case NotFound:
		return http.StatusNotFound
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

var (
	ErrBadRequest     = NewError("Bad request", BadRequest)
	ErrInternalServer = NewError("Internal server error", Internal)
	ErrDatabase       = NewError("Database raised an error", Internal)
)
