package service

import (
	"context"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/novando/byfood/be/internal/dto"
	"github.com/novando/byfood/be/pkg/reposqlc"
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
