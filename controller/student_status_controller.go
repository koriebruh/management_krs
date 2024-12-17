package controller

import (
	"github.com/gofiber/fiber/v2"
	"koriebruh/try/dto"
	"koriebruh/try/helper"
	"koriebruh/try/service"
	"net/http"
)

type StudentStatusController interface {
	InformationStudent(ctx *fiber.Ctx) error
	//SetClassTime()
	//GetAllKRSPick()
	//ExceptionInsertKRS()
	//StatusKRS()

}
type StudentStatusControllerImpl struct {
	service.StudentStatusService
}

func NewStudentStatusController(studentStatusService service.StudentStatusService) *StudentStatusControllerImpl {
	return &StudentStatusControllerImpl{StudentStatusService: studentStatusService}
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
