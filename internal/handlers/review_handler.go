package handlers

import (
	"go-rest-api/internal/dto"
	"go-rest-api/internal/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetBookReviews godoc
// @Summary Get reviews for a book
// @Description Get all reviews for a specific book
// @Tags reviews
// @Accept json
// @Produce json
// @Param id path int true "Book ID" minimum(1)
// @Success 200 {array} dto.ReviewResponse
// @Failure 400 {object} map[string]string "Invalid ID format"
// @Failure 404 {object} map[string]string "Book not found"
// @Failure 500 {object} map[string]string "Server error"
// @Router /api/v1/books/{id}/reviews [get]
func GetBookReviews(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	// Verify that the book exists
	_, err = repository.GetBookByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	reviews, err := repository.GetReviewsByBookID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Convert models to DTOs
	reviewResponses := make([]dto.ReviewResponse, len(reviews))
	for i, review := range reviews {
		reviewResponses[i] = dto.ToReviewResponse(review)
	}

	c.JSON(http.StatusOK, reviewResponses)
}

// AddReview godoc
// @Summary Add review to book
// @Description Add a new review to a book
// @Tags reviews
// @Accept json
// @Produce json
// @Param id path int true "Book ID" minimum(1)
// @Param review body dto.CreateReviewRequest true "Review object that needs to be added"
// @Success 201 {object} dto.ReviewResponse
// @Failure 400 {object} map[string]string "Invalid ID format or request body"
// @Failure 404 {object} map[string]string "Book not found"
// @Failure 500 {object} map[string]string "Server error"
// @Router /api/v1/books/{id}/reviews [post]
func AddReview(c *gin.Context) {
	bookID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	// Verify that the book exists
	_, err = repository.GetBookByID(uint(bookID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	var req dto.CreateReviewRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Convert DTO to model
	review := dto.CreateReviewRequestToModel(req, uint(bookID))

	if err := repository.CreateReview(&review); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, dto.ToReviewResponse(review))
}

// UpdateReview godoc
// @Summary Update review
// @Description Update an existing review
// @Tags reviews
// @Accept json
// @Produce json
// @Param id path int true "Review ID" minimum(1)
// @Param review body dto.UpdateReviewRequest true "Review object that needs to be updated"
// @Success 200 {object} dto.ReviewResponse
// @Failure 400 {object} map[string]string "Invalid ID format or request body"
// @Failure 404 {object} map[string]string "Review not found"
// @Failure 500 {object} map[string]string "Server error"
// @Router /api/v1/reviews/{id} [put]
func UpdateReview(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	review, err := repository.GetReviewByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Review not found"})
		return
	}

	var req dto.UpdateReviewRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update model from DTO
	dto.UpdateReviewModelFromRequest(review, req)

	if err := repository.UpdateReview(review); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.ToReviewResponse(*review))
}

// DeleteReview godoc
// @Summary Delete review
// @Description Delete a review
// @Tags reviews
// @Accept json
// @Produce json
// @Param id path int true "Review ID" minimum(1)
// @Success 204 "No Content"
// @Failure 400 {object} map[string]string "Invalid ID format"
// @Failure 500 {object} map[string]string "Server error"
// @Router /api/v1/reviews/{id} [delete]
func DeleteReview(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	if err := repository.DeleteReview(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
