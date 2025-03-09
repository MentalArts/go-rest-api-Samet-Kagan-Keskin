package repository

import (
	"go-rest-api/internal/database"
	"go-rest-api/internal/models"
)

func CreateAuthor(author *models.Author) error {
	return database.DB.Create(author).Error
}

func GetAuthorByID(id uint) (*models.Author, error) {
	var author models.Author
	result := database.DB.Preload("Books").First(&author, id)
	return &author, result.Error
}

func GetAllAuthors() ([]models.Author, error) {
	var authors []models.Author
	result := database.DB.Preload("Books").Find(&authors)
	return authors, result.Error
}

func UpdateAuthor(author *models.Author) error {
	return database.DB.Save(author).Error
}

func DeleteAuthor(id uint) error {
	return database.DB.Delete(&models.Author{}, id).Error
}
