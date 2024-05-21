package controller

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/novando/byfood/be/internal/library/dto"
	"github.com/novando/byfood/be/internal/library/service"
	"github.com/novando/byfood/be/pkg/response"
	"github.com/novando/byfood/be/pkg/validator"
	"strconv"
)

type Book struct {
	bookService *service.Book
}

func NewBookController(sb *service.Book) *Book {
	return &Book{sb}
}

// Create godoc
// @Summary      Modify provided URL with certain operation
// @Tags         Books
// @Accept       json
// @Produce      json
// @Success      200  {object}  dto.CleanupResponse
// @Failure      400  {object}  response.StdResponse
// @Failure      404  {string}  "Not Found"
// @Router       /books/ [post]
func (c *Book) Create(ctx *fiber.Ctx) error {
	var payload dto.BookCreateRequest
	if err := ctx.BodyParser(&payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.StdResponse{
			Message: "REQUEST_ERROR",
			Data:    err.Error(),
		})
	}
	if err := validator.Validate(payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.StdResponse{
			Message: "VALIDATION_ERROR",
			Data:    err.Error(),
		})
	}
	if payload.Isbn != nil {
		err := fmt.Errorf("ISBN Should contain numeric with 10 or 13 characters long")
		if len(*payload.Isbn) == 10 || len(*payload.Isbn) == 13 {
			err = nil
		}
		if _, err = strconv.Atoi(*payload.Isbn); err != nil {
			err = fmt.Errorf("ISBN should be numeric")
		}
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(response.StdResponse{
				Message: "VALIDATION_ERROR",
				Data:    err.Error(),
			})
		}
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

// Read godoc
// @Summary     Get summary data of a book with pagination
// @Tags        Books
// @Produce     json
// @Param		title	query	string	false	"title of a book"
// @Success     200  {object}  response.StdResponse
// @Failure     400  {object}  response.StdResponse
// @Failure     404  {string}  "Not Found"
// @Router      /books [post]
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

// Update godoc
// @Summary     Update whole data of a book
// @Tags        Books
// @Accept      json
// @Produce     json
// @Param		id	path	uuid	true	"Book ID"
// @Success     200  {object}  response.StdResponse
// @Failure     400  {object}  response.StdResponse
// @Failure     404  {string}  "Not Found"
// @Router      /books/{id} [put]
func (c *Book) Update(ctx *fiber.Ctx) error {
	bookId := ctx.Params("id", "")
	if bookId == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.StdResponse{
			Message: "BOOK_ID_REQUIRED",
		})
	}
	var payload dto.BookCreateRequest
	if err := ctx.BodyParser(&payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.StdResponse{
			Message: "REQUEST_ERROR",
			Data:    err.Error(),
		})
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

// Detail godoc
// @Summary     Get detailed data about a book
// @Tags        Books
// @Produce     json
// @Param		id	path	uuid	true	"Book ID"
// @Success     200  {object}  response.StdResponse
// @Failure     400  {object}  response.StdResponse
// @Failure     404  {string}  "Not Found"
// @Router      /books/{id} [get]
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

// Delete godoc
// @Summary     Soft delete a book by an ID
// @Tags        Books
// @Produce     json
// @Param		id	path	uuid	true	"Book ID"
// @Success     200  {object}  response.StdResponse
// @Failure     400  {object}  response.StdResponse
// @Failure     404  {string}  "Not Found"
// @Router      /books/{id} [delete]
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
