package controller

import (
	"github.com/gofiber/fiber/v2"
	"koriebruh/try/dto"
	"koriebruh/try/helper"
	"koriebruh/try/service"
	"net/http"
)

type StudentStatusController interface {
	KrsOffers(ctx *fiber.Ctx) error
	KrsSchedule(ctx *fiber.Ctx) error
	InformationStudent(ctx *fiber.Ctx) error
	SetClassTime(ctx *fiber.Ctx) error
	GetAllKRSPick(ctx *fiber.Ctx) error
	InsertKRSPermit(ctx *fiber.Ctx) error
	StatusKRS(ctx *fiber.Ctx) error
}
type StudentStatusControllerImpl struct {
	service.StudentStatusService
}

func NewStudentStatusController(studentStatusService service.StudentStatusService) *StudentStatusControllerImpl {
	return &StudentStatusControllerImpl{StudentStatusService: studentStatusService}
}

func (c StudentStatusControllerImpl) KrsOffers(ctx *fiber.Ctx) error {
	kodeTA := ctx.Query("kode-ta")
	offers, err := c.StudentStatusService.KrsOffers(ctx.Context(), kodeTA)
	if err != nil {
		return helper.ErrResponse(ctx, err)
	}

	return ctx.Status(http.StatusOK).JSON(dto.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   offers,
	})
}

func (c StudentStatusControllerImpl) KrsSchedule(ctx *fiber.Ctx) error {
	NimDinus := ctx.Locals("nim_dinus").(string)

	schedule, err := c.StudentStatusService.KrsSchedule(ctx.Context(), NimDinus)
	if err != nil {
		return helper.ErrResponse(ctx, err)
	}

	return ctx.Status(http.StatusOK).JSON(dto.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   schedule,
	})
}

func (c StudentStatusControllerImpl) InformationStudent(ctx *fiber.Ctx) error {
	NimDinus := ctx.Locals("nim_dinus").(string)

	studentInfo, err := c.StudentStatusService.InformationStudent(ctx.Context(), NimDinus)
	if err != nil {
		return helper.ErrResponse(ctx, err)
	}

	return ctx.Status(http.StatusOK).JSON(dto.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   studentInfo,
	})
}

func (c StudentStatusControllerImpl) SetClassTime(ctx *fiber.Ctx) error {
	var req dto.ChangeClassReq
	if err := ctx.BodyParser(&req); err != nil {
		return helper.ErrResponse(ctx, err)
	}
	NimDinus := ctx.Locals("nim_dinus").(string)

	if err := c.StudentStatusService.SetClassTime(ctx.Context(), NimDinus, req); err != nil {
		return helper.ErrResponse(ctx, err)
	}

	return ctx.Status(http.StatusOK).JSON(dto.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data: map[string]interface{}{
			"message": "success update class",
		},
	})

}

func (c StudentStatusControllerImpl) GetAllKRSPick(ctx *fiber.Ctx) error {

	NimDinus := ctx.Locals("nim_dinus").(string)

	pick, err := c.StudentStatusService.GetAllKRSPick(ctx.Context(), NimDinus)
	if err != nil {
		return helper.ErrResponse(ctx, err)
	}

	return ctx.Status(http.StatusOK).JSON(dto.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   pick,
	})

}

func (c StudentStatusControllerImpl) InsertKRSPermit(ctx *fiber.Ctx) error {
	NimDinus := ctx.Locals("nim_dinus").(string)

	permit, err := c.StudentStatusService.InsertKRSPermit(ctx.Context(), NimDinus)
	if err != nil {
		return helper.ErrResponse(ctx, err)
	}

	return ctx.Status(http.StatusOK).JSON(dto.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data: map[string]interface{}{
			"message": permit,
		},
	})

}

func (c StudentStatusControllerImpl) StatusKRS(ctx *fiber.Ctx) error {
	NimDinus := ctx.Locals("nim_dinus").(string)

	krsStatus, err := c.StudentStatusService.StatusKRS(ctx.Context(), NimDinus)
	if err != nil {
		return helper.ErrResponse(ctx, err)
	}

	return ctx.Status(http.StatusOK).JSON(dto.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   krsStatus,
	})

}
