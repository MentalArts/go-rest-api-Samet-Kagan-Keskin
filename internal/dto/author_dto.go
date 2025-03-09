package dto

// CreateAuthorRequest represents the request body for creating an author
type CreateAuthorRequest struct {
	Name      string `json:"name" binding:"required" example:"Oguz Atay"`
	Biography string `json:"biography" binding:"required" example:"The mixing of dream and reality in his works, metafiction being the main principle of fiction"`
	BirthDate string `json:"birth_date" binding:"required" example:"1961-10-12"`
}

// UpdateAuthorRequest represents the request body for updating an author
type UpdateAuthorRequest struct {
	Name      string `json:"name" example:"Oguz Atay"`
	Biography string `json:"biography" example:"Famous Turkish writer Oguz Atay"`
	BirthDate string `json:"birth_date" example:"1961-10-12"`
}

// AuthorResponse represents the response body for author information
type AuthorResponse struct {
	ID        uint   `json:"id" example:"1"`
	Name      string `json:"name" example:"Oguz Atay"`
	Biography string `json:"biography" example:"The mixing of dream and reality in his works, metafiction being the main principle of fiction"`
	BirthDate string `json:"birth_date" example:"1961-10-12"`
}

// AuthorDetailResponse includes books in the response
type AuthorDetailResponse struct {
	ID        uint           `json:"id" example:"1"`
	Name      string         `json:"name" example:"Oguz Atay"`
	Biography string         `json:"biography" example:"The mixing of dream and reality in his works, metafiction being the main principle of fiction"`
	BirthDate string         `json:"birth_date" example:"1961-10-12"`
	Books     []BookResponse `json:"books,omitempty"`
}
