package dto

import "time"

type BookCreateRequest struct {
	Title  string  `json:"title" validate:"required"`
	Yop    int16   `json:"yop" validate:"min=0,max=2024"`
	Author string  `json:"author" validate:"required"`
	Isbn   *string `json:"isbn"`
	Page   *int    `json:"page"`
}

type BookRequest struct {
	Title  string `json:"title" query:"title"`
	Yop    int16  `json:"yop" query:"yop"`
	Page   int    `json:"page" query:"page"`
	Size   int    `json:"size" query:"size"`
	SortBy string `json:"sortBy" query:"sortBy"`
	Asc    bool   `json:"asc" query:"asc"`
}

type BookResponse struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Yop    int16  `json:"yop"`
	Author string `json:"author"`
}

type BookDetailResponse struct {
	Title     string    `json:"title"`
	Yop       int16     `json:"yop"`
	Author    string    `json:"author"`
	Isbn      string    `json:"isbn"`
	Page      int       `json:"page"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
