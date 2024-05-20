package dto

type BookCreateRequest struct {
	Title  string  `json:"title" validate:"required"`
	Yop    int16   `json:"yop" validate:"min=0,max=2024"`
	Author string  `json:"author" validate:"required"`
	Isbn   *string `json:"isbn" validate:"numeric"`
	Page   *int    `json:"page" validate:"min=0"`
}
