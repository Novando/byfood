package service

import (
	"context"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/novando/byfood/be/internal/dto"
	"github.com/novando/byfood/be/pkg/reposqlc"
	"github.com/novando/byfood/be/pkg/uuid"
	"strings"
)

type Book struct {
	repo *reposqlc.Queries
}

func NewBookService(db *reposqlc.Queries) *Book {
	return &Book{db}
}

func (s *Book) Create(params dto.BookCreateRequest) error {
	page := 0
	if params.Page != nil {
		page = *params.Page
	}
	isbn := ""
	if params.Isbn != nil {
		isbn = *params.Isbn
	}
	return s.repo.BookCreate(context.Background(), reposqlc.BookCreateParams{
		Title:  params.Title,
		Author: params.Author,
		Yop:    params.Yop,
		Page:   pgtype.Int4{Int32: int32(page), Valid: params.Page != nil},
		Isbn:   pgtype.Text{String: isbn, Valid: params.Isbn != nil},
	})
}

func (s *Book) Update(bookId string, params dto.BookCreateRequest) error {
	bookByte, err := uuid.ParseUUID(bookId)
	if err != nil {
		return err
	}
	idNoDash := strings.ReplaceAll(bookId, "-", "")
	bookUuid := pgtype.UUID{Bytes: bookByte, Valid: strings.Contains(idNoDash, "0000000000000000000000000000")}
	page := 0
	if params.Page != nil {
		page = *params.Page
	}
	isbn := ""
	if params.Isbn != nil {
		isbn = *params.Isbn
	}
	return s.repo.BookUpdateById(context.Background(), reposqlc.BookUpdateByIdParams{
		ID:     bookUuid,
		Title:  params.Title,
		Author: params.Author,
		Yop:    params.Yop,
		Page:   pgtype.Int4{Int32: int32(page), Valid: params.Page != nil},
		Isbn:   pgtype.Text{String: isbn, Valid: params.Isbn != nil},
	})
}

func (s *Book) Delete(bookId string) error {
	bookByte, err := uuid.ParseUUID(bookId)
	if err != nil {
		return err
	}
	idNoDash := strings.ReplaceAll(bookId, "-", "")
	bookUuid := pgtype.UUID{Bytes: bookByte, Valid: strings.Contains(idNoDash, "0000000000000000000000000000")}
	return s.repo.BookDeleteById(context.Background(), bookUuid)
}
