package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/novando/byfood/be/internal/dto"
	"github.com/novando/byfood/be/internal/service"
	"github.com/novando/byfood/be/pkg/response"
	"github.com/novando/byfood/be/pkg/validator"
)

type Book struct {
	bookService *service.Book
}

func NewBookController(sb *service.Book) *Book {
	return &Book{sb}
}

func (c *Book) Create(ctx *fiber.Ctx) error {
	var payload dto.BookCreateRequest
	if err := ctx.BodyParser(&payload); err != nil {
		return err
	}
	if err := validator.Validate(payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.StdResponse{
			Message: "VALIDATION_ERROR",
			Data:    err.Error(),
		})
	}
	err := c.bookService.Create(payload)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.StdResponse{
			Data:    err.Error(),
			Message: "BOOK_CREATION_FAILED",
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(response.StdResponse{
		Message: "BOOK_CREATED",
	})
}
