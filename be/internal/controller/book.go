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

func (c *Book) Read(ctx *fiber.Ctx) error {
	var query dto.BookRequest
	if err := ctx.QueryParser(&query); err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(response.StdResponse{
			Message: "UNABLE_FETCH_REGISTER_REQUEST",
		})
	}
	if query.Page < 1 {
		query.Page = 1
	}
	if query.Size < 10 {
		query.Size = 10
	}
	res, total, err := c.bookService.Read(query)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.StdResponse{
			Data:    err.Error(),
			Message: "BOOK_FETCH_FAILED",
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(response.StdResponse{
		Message: "BOOK_FETCHED",
		Data:    res,
		Count:   total,
	})
}

func (c *Book) Update(ctx *fiber.Ctx) error {
	bookId := ctx.Params("id", "")
	if bookId == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.StdResponse{
			Message: "BOOK_ID_REQUIRED",
		})
	}
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
	err := c.bookService.Update(bookId, payload)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.StdResponse{
			Data:    err.Error(),
			Message: "BOOK_UPDATE_FAILED",
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(response.StdResponse{
		Message: "BOOK_UPDATED",
	})
}

func (c *Book) Detail(ctx *fiber.Ctx) error {
	bookId := ctx.Params("id", "")
	if bookId == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.StdResponse{
			Message: "BOOK_ID_REQUIRED",
		})
	}
	res, err := c.bookService.Detail(bookId)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.StdResponse{
			Data:    err.Error(),
			Message: "BOOK_DETAIL_FETCH_FAILED",
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(response.StdResponse{
		Message: "BOOK_DETAIL_FETCHED",
		Data:    res,
	})
}

func (c *Book) Delete(ctx *fiber.Ctx) error {
	bookId := ctx.Params("id", "")
	if bookId == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.StdResponse{
			Message: "BOOK_ID_REQUIRED",
		})
	}
	err := c.bookService.Delete(bookId)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.StdResponse{
			Data:    err.Error(),
			Message: "BOOK_DELETION_FAILED",
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(response.StdResponse{
		Message: "BOOK_DELETED",
	})
}
