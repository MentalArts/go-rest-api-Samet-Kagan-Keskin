package handlers

import (
	"go-rest-api/internal/dto"
	"go-rest-api/internal/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetAuthors godoc
// @Summary Get all authors
// @Description Get a list of all authors with their books
// @Tags authors
// @Accept json
// @Produce json
// @Success 200 {array} dto.AuthorDetailResponse
// @Failure 500 {object} map[string]string "Server error"
// @Router /api/v1/authors [get]
func GetAuthors(c *gin.Context) {
	authors, err := repository.GetAllAuthors()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Convert to DTOs
	response := make([]dto.AuthorDetailResponse, len(authors))
	for i, author := range authors {
		response[i] = dto.ToAuthorDetailResponse(author)
	}

	c.JSON(http.StatusOK, response)
}

// GetAuthor godoc
// @Summary Get author by ID
// @Description Get an author's details by ID
// @Tags authors
// @Accept json
// @Produce json
// @Param id path int true "Author ID" minimum(1)
// @Success 200 {object} dto.AuthorDetailResponse
// @Failure 400 {object} map[string]string "Invalid ID format"
// @Failure 404 {object} map[string]string "Author not found"
// @Router /api/v1/authors/{id} [get]
func GetAuthor(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	author, err := repository.GetAuthorByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Author not found"})
		return
	}

	c.JSON(http.StatusOK, dto.ToAuthorDetailResponse(*author))
}

// CreateAuthor godoc
// @Summary Create new author
// @Description Create a new author
// @Tags authors
// @Accept json
// @Produce json
// @Param author body dto.CreateAuthorRequest true "Author object that needs to be added"
// @Success 201 {object} dto.AuthorResponse
// @Failure 400 {object} map[string]string "Invalid input"
// @Failure 500 {object} map[string]string "Server error"
// @Router /api/v1/authors [post]
func CreateAuthor(c *gin.Context) {
	var req dto.CreateAuthorRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Convert DTO to model
	author := dto.CreateAuthorRequestToModel(req)

	if err := repository.CreateAuthor(&author); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, dto.ToAuthorResponse(author))
}

// UpdateAuthor godoc
// @Summary Update author
// @Description Update an existing author
// @Tags authors
// @Accept json
// @Produce json
// @Param id path int true "Author ID" minimum(1)
// @Param author body dto.UpdateAuthorRequest true "Author object that needs to be updated"
// @Success 200 {object} dto.AuthorResponse
// @Failure 400 {object} map[string]string "Invalid ID format or request body"
// @Failure 404 {object} map[string]string "Author not found"
// @Failure 500 {object} map[string]string "Server error"
// @Router /api/v1/authors/{id} [put]
func UpdateAuthor(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	author, err := repository.GetAuthorByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Author not found"})
		return
	}

	var req dto.UpdateAuthorRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update model from DTO
	dto.UpdateAuthorModelFromRequest(author, req)

	if err := repository.UpdateAuthor(author); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.ToAuthorResponse(*author))
}

// DeleteAuthor godoc
// @Summary Delete author
// @Description Delete an author
// @Tags authors
// @Accept json
// @Produce json
// @Param id path int true "Author ID" minimum(1)
// @Success 204 "No Content"
// @Failure 400 {object} map[string]string "Invalid ID format"
// @Failure 500 {object} map[string]string "Server error"
// @Router /api/v1/authors/{id} [delete]
func DeleteAuthor(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	if err := repository.DeleteAuthor(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
