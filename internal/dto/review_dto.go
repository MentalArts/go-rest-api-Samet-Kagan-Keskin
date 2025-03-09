package dto

// CreateReviewRequest represents the request body for creating a review
type CreateReviewRequest struct {
	Rating  int    `json:"rating" binding:"required,min=1,max=5" example:"5"`
	Comment string `json:"comment" binding:"required" example:"A masterpiece of fantasy literature!"`
	BookID  uint   `json:"book_id" binding:"required" example:"1"`
}

// UpdateReviewRequest represents the request body for updating a review
type UpdateReviewRequest struct {
	Rating  int    `json:"rating" binding:"min=1,max=5" example:"4"`
	Comment string `json:"comment" example:"Very good book, but a bit too long."`
}

// ReviewResponse represents the response body for review information
type ReviewResponse struct {
	ID         uint   `json:"id" example:"1"`
	Rating     int    `json:"rating" example:"5"`
	Comment    string `json:"comment" example:"A masterpiece of fantasy literature!"`
	DatePosted string `json:"date_posted" example:"2025-03-09"`
	BookID     uint   `json:"book_id" example:"1"`
}
