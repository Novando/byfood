package dto

type BookCreateRequest struct {
	Title  string  `json:"title" validate:"required"`
	Yop    int16   `json:"yop" validate:"min=0,max=2024"`
	Author string  `json:"author" validate:"required"`
	Isbn   *string `json:"isbn" validate:"numeric"`
	Page   *int    `json:"page" validate:"min=0"`
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
