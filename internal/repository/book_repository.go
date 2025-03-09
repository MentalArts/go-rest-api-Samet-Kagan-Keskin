package repository

import (
	"go-rest-api/internal/database"
	"go-rest-api/internal/models"
	"gorm.io/gorm"
)

func CreateBook(book *models.Book) error {
	return database.DB.Create(book).Error
}

func GetBookByID(id uint) (*models.Book, error) {
	var book models.Book
	result := database.DB.Preload("Author").Preload("Reviews").First(&book, id)
	return &book, result.Error
}

func GetAllBooks(page, pageSize int) ([]models.Book, int64, error) {
	var books []models.Book
	var count int64

	// Get total count
	if err := database.DB.Model(&models.Book{}).Count(&count).Error; err != nil {
		return nil, 0, err
	}

	// Get paginated books
	offset := (page - 1) * pageSize
	result := database.DB.Preload("Author").Offset(offset).Limit(pageSize).Find(&books)
	return books, count, result.Error
}

func UpdateBook(book *models.Book) error {
	return database.DB.Save(book).Error
}

func DeleteBook(id uint) error {
	// Start a transaction to delete the book and its reviews
	return database.DB.Transaction(func(tx *gorm.DB) error {
		// Delete all reviews for this book
		if err := tx.Where("book_id = ?", id).Delete(&models.Review{}).Error; err != nil {
			return err
		}

		// Delete the book
		return tx.Delete(&models.Book{}, id).Error
	})
}
