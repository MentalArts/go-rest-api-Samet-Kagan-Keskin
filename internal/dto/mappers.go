package dto

import (
	"go-rest-api/internal/models"
	"time"
)

// Convert models to DTOs

// ToAuthorResponse converts an Author model to AuthorResponse DTO
func ToAuthorResponse(author models.Author) AuthorResponse {
	return AuthorResponse{
		ID:        author.ID,
		Name:      author.Name,
		Biography: author.Biography,
		BirthDate: author.BirthDate,
	}
}

// ToAuthorDetailResponse converts an Author model to AuthorDetailResponse DTO
func ToAuthorDetailResponse(author models.Author) AuthorDetailResponse {
	bookResponses := make([]BookResponse, len(author.Books))
	for i, book := range author.Books {
		bookResponses[i] = ToBookResponse(book)
	}

	return AuthorDetailResponse{
		ID:        author.ID,
		Name:      author.Name,
		Biography: author.Biography,
		BirthDate: author.BirthDate,
		Books:     bookResponses,
	}
}

// ToBookResponse converts a Book model to BookResponse DTO
func ToBookResponse(book models.Book) BookResponse {
	return BookResponse{
		ID:              book.ID,
		Title:           book.Title,
		AuthorID:        book.AuthorID,
		ISBN:            book.ISBN,
		PublicationYear: book.PublicationYear,
		Description:     book.Description,
	}
}

// ToBookDetailResponse converts a Book model to BookDetailResponse DTO
func ToBookDetailResponse(book models.Book) BookDetailResponse {
	reviewResponses := make([]ReviewResponse, len(book.Reviews))
	for i, review := range book.Reviews {
		reviewResponses[i] = ToReviewResponse(review)
	}

	return BookDetailResponse{
		ID:              book.ID,
		Title:           book.Title,
		AuthorID:        book.AuthorID,
		Author:          ToAuthorResponse(book.Author),
		ISBN:            book.ISBN,
		PublicationYear: book.PublicationYear,
		Description:     book.Description,
		Reviews:         reviewResponses,
	}
}

// ToReviewResponse converts a Review model to ReviewResponse DTO
func ToReviewResponse(review models.Review) ReviewResponse {
	return ReviewResponse{
		ID:         review.ID,
		Rating:     review.Rating,
		Comment:    review.Comment,
		DatePosted: review.DatePosted,
		BookID:     review.BookID,
	}
}

// Convert DTOs to models

// CreateAuthorRequestToModel converts CreateAuthorRequest DTO to Author model
func CreateAuthorRequestToModel(req CreateAuthorRequest) models.Author {
	return models.Author{
		Name:      req.Name,
		Biography: req.Biography,
		BirthDate: req.BirthDate,
	}
}

// UpdateAuthorModelFromRequest updates Author model from UpdateAuthorRequest DTO
func UpdateAuthorModelFromRequest(author *models.Author, req UpdateAuthorRequest) {
	if req.Name != "" {
		author.Name = req.Name
	}
	if req.Biography != "" {
		author.Biography = req.Biography
	}
	if req.BirthDate != "" {
		author.BirthDate = req.BirthDate
	}
}

// CreateBookRequestToModel converts CreateBookRequest DTO to Book model
func CreateBookRequestToModel(req CreateBookRequest) models.Book {
	return models.Book{
		Title:           req.Title,
		AuthorID:        req.AuthorID,
		ISBN:            req.ISBN,
		PublicationYear: req.PublicationYear,
		Description:     req.Description,
	}
}

// UpdateBookModelFromRequest updates Book model from UpdateBookRequest DTO
func UpdateBookModelFromRequest(book *models.Book, req UpdateBookRequest) {
	if req.Title != "" {
		book.Title = req.Title
	}
	if req.AuthorID != 0 {
		book.AuthorID = req.AuthorID
	}
	if req.ISBN != "" {
		book.ISBN = req.ISBN
	}
	if req.PublicationYear != 0 {
		book.PublicationYear = req.PublicationYear
	}
	if req.Description != "" {
		book.Description = req.Description
	}
}

// CreateReviewRequestToModel converts CreateReviewRequest DTO to Review model
func CreateReviewRequestToModel(req CreateReviewRequest, bookID uint) models.Review {
	return models.Review{
		Rating:     req.Rating,
		Comment:    req.Comment,
		DatePosted: time.Now().Format("2006-01-02"),
		BookID:     bookID,
	}
}

// UpdateReviewModelFromRequest updates Review model from UpdateReviewRequest DTO
func UpdateReviewModelFromRequest(review *models.Review, req UpdateReviewRequest) {
	if req.Rating != 0 {
		review.Rating = req.Rating
	}
	if req.Comment != "" {
		review.Comment = req.Comment
	}
}
