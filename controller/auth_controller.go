package controller

import (
	"github.com/gofiber/fiber/v2"
	"koriebruh/try/dto"
	"koriebruh/try/helper"
	"koriebruh/try/service"
	"net/http"
)

type AuthController interface {
	Register(ctx *fiber.Ctx) error
	Login(ctx *fiber.Ctx) error
	CurrentAcc(ctx *fiber.Ctx) error
}

type AuthControllerImpl struct {
	service.AuthService
}

func NewAuthController(authService service.AuthService) *AuthControllerImpl {
	return &AuthControllerImpl{AuthService: authService}
}

func (controller AuthControllerImpl) Register(ctx *fiber.Ctx) error {
	var req dto.RegisterReq
	if err := ctx.BodyParser(&req); err != nil {
		return helper.ErrResponse(ctx, err)
	}

	err := controller.AuthService.Register(ctx.Context(), req)
	if err != nil {
		return helper.ErrResponse(ctx, err)
	}

	return ctx.Status(http.StatusCreated).JSON(dto.WebResponse{
		Code:   http.StatusCreated,
		Status: "User Registered Successfully",
		Data:   nil,
	})
}

func (controller AuthControllerImpl) Login(ctx *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (controller AuthControllerImpl) CurrentAcc(ctx *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}
