package repository

import (
	"go-rest-api/internal/database"
	"go-rest-api/internal/models"
)

func CreateReview(review *models.Review) error {
	return database.DB.Create(review).Error
}

func GetReviewsByBookID(bookID uint) ([]models.Review, error) {
	var reviews []models.Review
	result := database.DB.Where("book_id = ?", bookID).Find(&reviews)
	return reviews, result.Error
}

func GetReviewByID(id uint) (*models.Review, error) {
	var review models.Review
	result := database.DB.First(&review, id)
	return &review, result.Error
}

func UpdateReview(review *models.Review) error {
	return database.DB.Save(review).Error
}

func DeleteReview(id uint) error {
	return database.DB.Delete(&models.Review{}, id).Error
}
