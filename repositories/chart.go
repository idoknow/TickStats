package repositories

import (
	"github.com/soulter/tickstats/models"
	"gorm.io/gorm"
)

type ChartRepository interface {
	Create(chart *models.Chart) error
	FindByAppID(appId string) ([]models.Chart, error)
}

type chartRepository struct {
	db *gorm.DB
}

func NewChartRepository(db *gorm.DB) ChartRepository {
	return &chartRepository{db}
}

func (r *chartRepository) Create(chart *models.Chart) error {
	return r.db.Create(chart).Error
}

func (r *chartRepository) FindByAppID(appId string) ([]models.Chart, error) {
	var charts []models.Chart
	if err := r.db.Where("app_id = ?", appId).Find(&charts).Error; err != nil {
		return nil, err
	}
	return charts, nil
}
