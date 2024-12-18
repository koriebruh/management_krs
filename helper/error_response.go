package helper

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"koriebruh/try/dto"
	"net/http"
)

var (
	ErrValidationFailed   = errors.New("validation failed")
	ErrPasswordEncryption = errors.New("password encryption failed")
	ErrUserRegistration   = errors.New("user registration failed")
	ErrLoginFailed        = errors.New("login failed ")
	ErrNotFound           = errors.New("data not found")
	ErrBadRequest         = errors.New("bad request")
)

func ErrResponse(ctx *fiber.Ctx, err error) error {
	switch {
	case errors.Is(err, ErrValidationFailed):
		return ctx.Status(http.StatusBadRequest).JSON(dto.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "Validation Error",
			Data:   err.Error(),
		})
	case errors.Is(err, ErrLoginFailed):
		return ctx.Status(http.StatusBadRequest).JSON(dto.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "Failed Login",
			Data:   err.Error(),
		})
	case errors.Is(err, ErrNotFound):
		return ctx.Status(http.StatusBadRequest).JSON(dto.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "Data not found",
			Data:   err.Error(),
		})
	case errors.Is(err, ErrPasswordEncryption):
		return ctx.Status(http.StatusInternalServerError).JSON(dto.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "Password Encryption Error",
			Data:   err.Error(),
		})
	case errors.Is(err, ErrUserRegistration):
		return ctx.Status(http.StatusInternalServerError).JSON(dto.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "Registration Failed",
			Data:   err.Error(),
		})
	case errors.Is(err, ErrBadRequest):
		return ctx.Status(http.StatusBadRequest).JSON(dto.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Data:   err.Error(),
		})
	case err != nil:
		return ctx.Status(http.StatusInternalServerError).JSON(dto.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "Internal Server Error",
			Data:   err.Error(),
		})
	}
	return nil
}
