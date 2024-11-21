package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"koriebruh/try/conf"
	"koriebruh/try/dto"
	"koriebruh/try/helper"
	"koriebruh/try/service"
	"net/http"
	"time"
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
	var req dto.LoginReq
	if err := ctx.BodyParser(&req); err != nil {
		return helper.ErrResponse(ctx, err)
	}

	nimUser, err := controller.AuthService.Login(ctx.Context(), req)
	if err != nil {
		return helper.ErrResponse(ctx, err)
	}

	//SETTING GENERATE JWT
	expTime := time.Now().Add(time.Minute * 3) // << KADALUARSA DALAM 3 minute
	claims := conf.JWTClaim{
		NIM: nimUser,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "koriebruh",
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}
	tokenAlgo := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenValue, err := tokenAlgo.SignedString([]byte(conf.JWT_KEY))

	return ctx.Status(http.StatusCreated).JSON(dto.WebResponse{
		Code:   http.StatusCreated,
		Status: "User Registered Successfully",
		Data: map[string]string{
			"message": "Login Success",
			"token":   tokenValue,
		},
	})
}

func (controller AuthControllerImpl) CurrentAcc(ctx *fiber.Ctx) error {
	nim := ctx.Locals("userNIM").(string)

	acc, err := controller.AuthService.CurrentAcc(ctx.Context(), nim)
	if err != nil {
		return helper.ErrResponse(ctx, err)
	}

	return ctx.Status(http.StatusCreated).JSON(dto.WebResponse{
		Code:   http.StatusCreated,
		Status: "User Registered Successfully",
		Data:   acc,
	})

}
