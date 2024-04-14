package utils

import (
	"context"
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Error struct {
	Status ErrorStatus
	Msg    string
}

func (e Error) Error() string {
	return e.Msg
}

const (
	ErrInternal ErrorStatus = iota
	ErrBadRequest
	ErrNotFound
)

type ErrorStatus int

func NewError(st ErrorStatus, msg string) error {
	return &Error{
		Msg:    msg,
		Status: st,
	}
}

func FromError(e error) (ErrorStatus, string) {
	var err *Error
	if !errors.As(e, &err) {
		if errors.Is(e, context.DeadlineExceeded) {
			return ErrInternal, "request took for too long"
		}
		return ErrInternal, "unknown error type"
	}

	switch err.Status {
	case ErrBadRequest:
		return http.StatusBadRequest, err.Msg
	case ErrNotFound:
		return http.StatusNotFound, err.Msg
	default:
		return ErrInternal, e.Error()
	}
}

func HandleError(fn echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if err := fn(c); err != nil {
			st, msg := FromError(err)
			c.JSON(int(st), echo.Map{"error": msg})
		}
		return nil
	}
}
