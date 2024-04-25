package repositories

import (
	"orderfaz-test-go/models"

	"gorm.io/gorm"
)

type LogisticRepository interface {
	FindLogistics() ([]models.Logistic, error)
	GetLogistic(ID int) (models.Logistic, error)
	GetLogisticBody(originName string, destinationName string) ([]models.Logistic, error)
	CreateLogistic(logistic models.Logistic) (models.Logistic, error)
	UpdateLogistic(logistic models.Logistic) (models.Logistic, error)
	DeleteLogistic(logistic models.Logistic) (models.Logistic, error)
}

func RepositoryLogistic(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindLogistics() ([]models.Logistic, error) {
	var logistics []models.Logistic
	err := r.db.Find(&logistics).Error

	return logistics, err
}

func (r *repository) GetLogistic(ID int) (models.Logistic, error) {
	var logistic models.Logistic
	err := r.db.First(&logistic, ID).Error

	return logistic, err
}

func (r *repository) GetLogisticBody(originName string, destinationName string) ([]models.Logistic, error) {
	var logistic []models.Logistic
	err := r.db.Where("origin_name = ?", originName).Find(&logistic, "destination_name = ?", destinationName).Error

	return logistic, err
}

func (r *repository) CreateLogistic(logistic models.Logistic) (models.Logistic, error) {
	err := r.db.Create(&logistic).Error

	return logistic, err
}

func (r *repository) UpdateLogistic(logistic models.Logistic) (models.Logistic, error) {
	err := r.db.Save(&logistic).Error

	return logistic, err
}

func (r *repository) DeleteLogistic(logistic models.Logistic) (models.Logistic, error) {
	err := r.db.Delete(&logistic).Error

	return logistic, err
}
