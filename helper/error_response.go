package helper

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"koriebruh/try/dto"
	"net/http"
	"strings"
)

var (
	ErrInternalServer = errors.New("INTERNAL SERVER ERROR")
	ErrNotFound       = errors.New("NOT FOUND")
	ErrBadRequest     = errors.New("BAD REQUEST")
)

func extractErrorMessage(err error) string {
	parts := strings.SplitN(err.Error(), ": ", 2)
	if len(parts) > 1 {
		return parts[1]
	}
	return err.Error()
}

func ErrResponse(ctx *fiber.Ctx, err error) error {
	switch {
	case errors.Is(err, ErrInternalServer):
		return ctx.Status(http.StatusInternalServerError).JSON(dto.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL SERVER ERROR",
			Data: map[string]interface{}{
				"error": extractErrorMessage(err),
			},
		})
	case errors.Is(err, ErrNotFound):
		return ctx.Status(http.StatusNotFound).JSON(dto.WebResponse{
			Code:   http.StatusNotFound,
			Status: "NOT FOUND",
			Data: map[string]interface{}{
				"error": extractErrorMessage(err),
			},
		})
	case errors.Is(err, ErrBadRequest):
		return ctx.Status(http.StatusBadRequest).JSON(dto.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data: map[string]interface{}{
				"error": extractErrorMessage(err),
			},
		})
	case err != nil:
		return ctx.Status(http.StatusInternalServerError).JSON(dto.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL SERVER ERROR 2",
			Data: map[string]interface{}{
				"error": extractErrorMessage(err),
			},
		})
	}
	return nil
}
