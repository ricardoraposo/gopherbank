package handlers

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type Error struct {
	Code int    `json:"code"`
	Err  string `json:"error"`
}

func ErrorHandler(c *fiber.Ctx, err error) error {
	if e, ok := err.(Error); ok {
		return c.Status(e.Code).JSON(fiber.Map{"error": e.Err})
	}
	e := NewError(http.StatusInternalServerError, err.Error())
	return c.Status(e.Code).JSON(fiber.Map{"error": e.Err})
}

func (e Error) Error() string {
	return e.Err
}

func NewError(code int, err string) Error {
	return Error{
		Code: code,
		Err:  err,
	}
}

func ErrNotAuthorized() Error {
	return NewError(http.StatusUnauthorized, "Operation not authorized")
}

func ErrTokenNotFound() Error {
	return NewError(http.StatusUnauthorized, "Token not found")
}

func ErrInvalidCredentials() Error {
	return NewError(http.StatusUnauthorized, "Invalid credentials")
}

func ErrNotFound(entity string) Error {
	return NewError(http.StatusNotFound, fmt.Sprintf("%s not found", entity))
}

func ErrBadRequest() Error {
	return NewError(http.StatusBadRequest, "Bad request")
}
