package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/novando/byfood/be/internal/url-modifier/dto"
	"github.com/novando/byfood/be/internal/url-modifier/service"
	"github.com/novando/byfood/be/pkg/response"
	"github.com/novando/byfood/be/pkg/validator"
)

// ProcessUrl godoc
// @Summary      Modify provided URL with certain operation
// @Tags         URL Modify
// @Accept       json
// @Produce      json
// @Success      200  {object}  dto.CleanupResponse
// @Failure      400  {object}  response.StdResponse
// @Failure      404  {string}  "Not Found"
// @Router       /url-modifier [post]
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
