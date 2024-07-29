package repositories

import (
	"github.com/soulter/tickstats/models"
	"gorm.io/gorm"
)

type ApplicationRepository interface {
	Create(application *models.Application) error
	FindByAccountID(accountId int) ([]models.Application, error)
	FindByAppID(appId string) (*models.Application, error)
}

type applicationRepository struct {
	db *gorm.DB
}

func NewApplicationRepository(db *gorm.DB) ApplicationRepository {
	return &applicationRepository{db}
}

func (r *applicationRepository) Create(application *models.Application) error {
	return r.db.Create(application).Error
}

func (r *applicationRepository) FindByAccountID(accountId int) ([]models.Application, error) {
	var applications []models.Application
	if err := r.db.Where("account_id = ?", accountId).Find(&applications).Error; err != nil {
		return nil, err
	}
	return applications, nil
}

func (r *applicationRepository) FindByAppID(appId string) (*models.Application, error) {
	var application models.Application
	if err := r.db.Where("app_id = ?", appId).First(&application).Error; err != nil {
		return nil, err
	}
	return &application, nil
}
