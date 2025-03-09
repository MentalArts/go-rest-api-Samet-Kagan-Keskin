package dto

// CreateBookRequest represents the request body for creating a book
type CreateBookRequest struct {
	Title           string `json:"title" binding:"required" example:"Oguz Atay and The Unbearables"`
	AuthorID        uint   `json:"author_id" binding:"required" example:"1"`
	ISBN            string `json:"isbn" binding:"required" example:"9780747532699"`
	PublicationYear int    `json:"publication_year" binding:"required" example:"1997"`
	Description     string `json:"description" binding:"required" example:"Oguz Atay'ın first adventure"`
}

// UpdateBookRequest represents the request body for updating a book
type UpdateBookRequest struct {
	Title           string `json:"title" example:"Oguz Atay and The Unbearables"`
	AuthorID        uint   `json:"author_id" example:"1"`
	ISBN            string `json:"isbn" example:"9780747532699"`
	PublicationYear int    `json:"publication_year" example:"1997"`
	Description     string `json:"description" example:"Oguz Atay'ın first adventure"`
}

// BookResponse represents the response body for book information
type BookResponse struct {
	ID              uint   `json:"id" example:"1"`
	Title           string `json:"title" example:"Oguz Atay and The Unbearables"`
	AuthorID        uint   `json:"author_id" example:"1"`
	ISBN            string `json:"isbn" example:"9780747532699"`
	PublicationYear int    `json:"publication_year" example:"1997"`
	Description     string `json:"description" example:"Oguz Atay'ın first adventure"`
}

// BookDetailResponse includes author and reviews in the response
type BookDetailResponse struct {
	ID              uint             `json:"id" example:"1"`
	Title           string           `json:"title" example:"Oguz Atay and The Unbearables"`
	AuthorID        uint             `json:"author_id" example:"1"`
	Author          AuthorResponse   `json:"author,omitempty"`
	ISBN            string           `json:"isbn" example:"9780747532699"`
	PublicationYear int              `json:"publication_year" example:"1997"`
	Description     string           `json:"description" example:"Oguz Atay'ın first adventure"`
	Reviews         []ReviewResponse `json:"reviews,omitempty"`
}

// PaginatedBooksResponse represents paginated book list response
type PaginatedBooksResponse struct {
	Data       []BookResponse `json:"data"`
	Total      int64          `json:"total" example:"100"`
	Page       int            `json:"page" example:"1"`
	PageSize   int            `json:"page_size" example:"10"`
	TotalPages int            `json:"total_pages" example:"10"`
}
