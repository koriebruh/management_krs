package controller

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"koriebruh/try/dto"
	"koriebruh/try/helper"
	"koriebruh/try/service"
	"net/http"
	"strconv"
)

type StudentStatusController interface {
	KrsOffers(ctx *fiber.Ctx) error
	KrsSchedule(ctx *fiber.Ctx) error
	InformationStudent(ctx *fiber.Ctx) error
	SetClassTime(ctx *fiber.Ctx) error
	GetAllKRSPick(ctx *fiber.Ctx) error
	InsertKRSPermit(ctx *fiber.Ctx) error
	StatusKRSMhs(ctx *fiber.Ctx) error
	KrsOffersProdi(ctx *fiber.Ctx) error
	GetAllScores(ctx *fiber.Ctx) error
	ScheduleConflicts(ctx *fiber.Ctx) error
	InsertSchedule(ctx *fiber.Ctx) error
	GetKrsLog(ctx *fiber.Ctx) error
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

func (c StudentStatusControllerImpl) StatusKRSMhs(ctx *fiber.Ctx) error {
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

func (c StudentStatusControllerImpl) KrsOffersProdi(ctx *fiber.Ctx) error {
	NimDinus := ctx.Locals("nim_dinus").(string)
	kodeTA := ctx.Query("kode-ta")

	prodiSchedule, err := c.StudentStatusService.KrsOffersProdi(ctx.Context(), NimDinus, kodeTA)
	if err != nil {
		return helper.ErrResponse(ctx, err)
	}

	return ctx.Status(http.StatusOK).JSON(dto.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   prodiSchedule,
	})

}

func (c StudentStatusControllerImpl) GetAllScores(ctx *fiber.Ctx) error {
	NimDinus := ctx.Locals("nim_dinus").(string)

	scores, err := c.StudentStatusService.GetAllScores(ctx.Context(), NimDinus)
	if err != nil {
		return helper.ErrResponse(ctx, err)
	}

	return ctx.Status(http.StatusOK).JSON(dto.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   scores,
	})

}

func (c StudentStatusControllerImpl) ScheduleConflicts(ctx *fiber.Ctx) error {
	kodeTA := ctx.Query("kode-ta")
	NimDinus := ctx.Locals("nim_dinus").(string)

	conflicts, err := c.StudentStatusService.ScheduleConflicts(ctx.Context(), NimDinus, kodeTA)
	if err != nil {
		return helper.ErrResponse(ctx, err)
	}

	return ctx.Status(http.StatusOK).JSON(dto.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   conflicts,
	})
}

func (c StudentStatusControllerImpl) InsertSchedule(ctx *fiber.Ctx) error {
	kodeTA := ctx.Query("kode-ta")
	NimDinus := ctx.Locals("nim_dinus").(string)
	params := ctx.Params("id")

	idSchedule, err := strconv.Atoi(params)
	if err != nil {
		err = fmt.Errorf("%w: %v", helper.ErrNotFound, fmt.Errorf("err param"))
		return helper.ErrResponse(ctx, err)
	}

	msg, err := c.StudentStatusService.InsertSchedule(ctx.Context(), NimDinus, kodeTA, idSchedule)
	if err != nil {
		return helper.ErrResponse(ctx, err)
	}

	return ctx.Status(http.StatusOK).JSON(dto.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data: map[string]interface{}{
			"message": msg,
		},
	})
}

func (c StudentStatusControllerImpl) GetKrsLog(ctx *fiber.Ctx) error {
	kodeTA := ctx.Query("kode-ta")
	NimDinus := ctx.Locals("nim_dinus").(string)

	log, err := c.StudentStatusService.GetKrsLog(ctx.Context(), NimDinus, kodeTA)
	if err != nil {
		return helper.ErrResponse(ctx, err)
	}

	return ctx.Status(http.StatusOK).JSON(dto.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   log,
	})
}
