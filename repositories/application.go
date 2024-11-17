package repositories

import (
	"github.com/soulter/tickstats/models"
	"gorm.io/gorm"
)

type ApplicationRepository interface {
	Create(application *models.Application) error
	Delete(appId string) error
	FindByAccountID(accountId int) ([]models.Application, error)
	FindByAppID(appId string) (*models.Application, error)
	FindPublicByPage(page, size int) ([]models.Application, error)
	Update(application *models.Application) error
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

func (r *applicationRepository) Delete(appId string) error {
	return r.db.Where("app_id = ?", appId).Delete(&models.Application{}).Error
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

func (r *applicationRepository) FindPublicByPage(page, size int) ([]models.Application, error) {
	var applications []models.Application
	if err := r.db.Offset((page-1)*size).Limit(size).Where("public = ?", true).Find(&applications).Error; err != nil {
		return nil, err
	}
	return applications, nil
}

func (r *applicationRepository) Update(application *models.Application) error {
	return r.db.Where("app_id = ?", application.AppId).Updates(application).Error
}