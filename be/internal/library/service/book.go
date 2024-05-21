package service

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/novando/byfood/be/internal/library/dto"
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

func (s *Book) Read(params dto.BookRequest) (res []dto.BookResponse, total int64, err error) {
	colName := reposqlc.BOOKS_CREATED_AT
	if params.SortBy == reposqlc.BOOKS_TITLE {
		colName = reposqlc.BOOKS_TITLE
	}
	if params.SortBy == reposqlc.BOOKS_YOP {
		colName = reposqlc.BOOKS_YOP
	}
	xtraParams := reposqlc.ColumnCustomParams{
		ColumnName: colName,
		Ascending:  params.Asc,
		Limit:      params.Size,
		Offset:     params.Page - 1,
	}
	total, err = s.repo.BookCount(context.Background(), reposqlc.BookCountParams{Title: params.Title, Yop: params.Yop})
	if err != nil {
		return
	}
	dao, err := s.repo.BookGetAll(context.Background(), xtraParams, reposqlc.BookGetAllParams{Title: params.Title, Yop: params.Yop})
	if err != nil {
		return
	}
	for _, item := range dao {
		res = append(res, dto.BookResponse{
			ID:     fmt.Sprintf("%x", item.ID.Bytes),
			Title:  item.Title,
			Yop:    item.Yop,
			Author: item.Author,
		})
	}
	return
}

func (s *Book) Detail(bookId string) (res dto.BookDetailResponse, err error) {
	bookByte, err := uuid.ParseUUID(bookId)
	if err != nil {
		return
	}
	idNoDash := strings.ReplaceAll(bookId, "-", "")
	bookUuid := pgtype.UUID{Bytes: bookByte, Valid: strings.Contains(idNoDash, "0000000000000000000000000000")}
	dao, err := s.repo.BookDetailById(context.Background(), bookUuid)
	if err != nil {
		return
	}
	res = dto.BookDetailResponse{
		Title:     dao.Title,
		Author:    dao.Author,
		Yop:       dao.Yop,
		Isbn:      dao.Isbn.String,
		Page:      int(dao.Page.Int32),
		CreatedAt: dao.CreatedAt.Time,
		UpdatedAt: dao.UpdatedAt.Time,
	}
	return
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
