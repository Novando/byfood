package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/novando/byfood/be/internal/dto"
	"github.com/novando/byfood/be/internal/service"
	"github.com/novando/byfood/be/pkg/response"
	"github.com/novando/byfood/be/pkg/validator"
)

func ProcessUrl(ctx *fiber.Ctx) error {
	var payload dto.CleanupRequest
	if err := ctx.BodyParser(&payload); err != nil {
		return err
	}
	if err := validator.Validate(payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.StdResponse{
			Message: "VALIDATION_ERROR",
			Data:    err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(dto.CleanupResponse{
		ProcessedUrl: service.ProcessUrl(payload),
	})
}
