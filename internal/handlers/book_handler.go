package handlers

import (
	"go-rest-api/internal/dto"
	"go-rest-api/internal/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetBooks godoc
// @Summary Get all books
// @Description Get a list of all books with pagination
// @Tags books
// @Accept json
// @Produce json
// @Param page query int false "Page number" minimum(1) default(1)
// @Param page_size query int false "Number of items per page" minimum(1) maximum(100) default(10)
// @Success 200 {object} dto.PaginatedBooksResponse "Returns paginated books data"
// @Failure 500 {object} map[string]string "Server error"
// @Router /api/v1/books [get]
func GetBooks(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	pageSizeStr := c.DefaultQuery("page_size", "10")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil || pageSize < 1 {
		pageSize = 10
	}
	if pageSize > 100 {
		pageSize = 100
	}

	books, totalCount, err := repository.GetAllBooks(page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Convert models to DTOs
	bookResponses := make([]dto.BookResponse, len(books))
	for i, book := range books {
		bookResponses[i] = dto.ToBookResponse(book)
	}

	response := dto.PaginatedBooksResponse{
		Data:       bookResponses,
		Total:      totalCount,
		Page:       page,
		PageSize:   pageSize,
		TotalPages: int((totalCount + int64(pageSize) - 1) / int64(pageSize)),
	}

	c.JSON(http.StatusOK, response)
}

// GetBook godoc
// @Summary Get book by ID
// @Description Get a book's details by ID with author and reviews
// @Tags books
// @Accept json
// @Produce json
// @Param id path int true "Book ID" minimum(1)
// @Success 200 {object} dto.BookDetailResponse
// @Failure 400 {object} map[string]string "Invalid ID format"
// @Failure 404 {object} map[string]string "Book not found"
// @Router /api/v1/books/{id} [get]
func GetBook(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	book, err := repository.GetBookByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	c.JSON(http.StatusOK, dto.ToBookDetailResponse(*book))
}

// CreateBook godoc
// @Summary Create new book
// @Description Add a new book to the library
// @Tags books
// @Accept json
// @Produce json
// @Param book body dto.CreateBookRequest true "Book object that needs to be added"
// @Success 201 {object} dto.BookResponse
// @Failure 400 {object} map[string]string "Invalid input or author does not exist"
// @Failure 500 {object} map[string]string "Server error"
// @Router /api/v1/books [post]
func CreateBook(c *gin.Context) {
	var req dto.CreateBookRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if author exists
	_, err := repository.GetAuthorByID(req.AuthorID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Author does not exist"})
		return
	}

	// Convert DTO to model
	book := dto.CreateBookRequestToModel(req)

	if err := repository.CreateBook(&book); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, dto.ToBookResponse(book))
}

// UpdateBook godoc
// @Summary Update book
// @Description Update an existing book
// @Tags books
// @Accept json
// @Produce json
// @Param id path int true "Book ID" minimum(1)
// @Param book body dto.UpdateBookRequest true "Book object that needs to be updated"
// @Success 200 {object} dto.BookResponse
// @Failure 400 {object} map[string]string "Invalid ID format or request body"
// @Failure 404 {object} map[string]string "Book not found"
// @Failure 500 {object} map[string]string "Server error"
// @Router /api/v1/books/{id} [put]
func UpdateBook(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	book, err := repository.GetBookByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	var req dto.UpdateBookRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update model from DTO
	dto.UpdateBookModelFromRequest(book, req)

	// Check if author exists if author_id has changed
	if req.AuthorID != 0 {
		_, err := repository.GetAuthorByID(req.AuthorID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Author does not exist"})
			return
		}
	}

	if err := repository.UpdateBook(book); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.ToBookResponse(*book))
}

// DeleteBook godoc
// @Summary Delete book
// @Description Delete a book
// @Tags books
// @Accept json
// @Produce json
// @Param id path int true "Book ID" minimum(1)
// @Success 204 "No Content"
// @Failure 400 {object} map[string]string "Invalid ID format"
// @Failure 500 {object} map[string]string "Server error"
// @Router /api/v1/books/{id} [delete]
func DeleteBook(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	if err := repository.DeleteBook(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
