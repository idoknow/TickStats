package repositories

import (
	"github.com/soulter/tickstats/models"
	"gorm.io/gorm"
)

type ChartRepository interface {
	Create(chart *models.Chart) error
	Delete(chartId string) error
	Update(chart *models.Chart) error
	FindByAppID(appId string, onlyPublic bool) ([]models.Chart, error)
	FindByChartID(chartId string) (*models.Chart, error)
	GetMaxRowID(appId string) (int64, error)
	GetCountOfRowID(chartId string, rowId int64) (int64, error)
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

func (r *chartRepository) Delete(chartId string) error {
	return r.db.Where("chart_id = ?", chartId).Delete(&models.Chart{}).Error
}

func (r *chartRepository) Update(chart *models.Chart) error {
	return r.db.Model(chart).Where("chart_id = ?", chart.ChartId).Select(models.ChartUpdatableFields).Updates(chart).Error
}

func (r *chartRepository) FindByAppID(appId string, onlyPublic bool) ([]models.Chart, error) {
	var charts []models.Chart
	var err error

	if onlyPublic {
		err = r.db.Where("app_id = ? AND public = ?", appId, true).Find(&charts).Error
	} else {
		err = r.db.Where("app_id = ?", appId).Find(&charts).Error
	}

	if err != nil {
		return nil, err
	}

	return charts, nil
}

func (r *chartRepository) FindByChartID(chartId string) (*models.Chart, error) {
	var chart models.Chart
	chart.ChartId = chartId
	err := r.db.Where("chart_id = ?", chartId).First(&chart).Error
	if err != nil {
		return nil, err
	}

	return &chart, nil
}

func (r *chartRepository) GetMaxRowID(appId string) (int64, error) {
	var chart models.Chart
	result := r.db.Where("app_id = ?", appId).Order("row_id desc").First(&chart)
	if result.RowsAffected == 0 {
		return 1, nil
	}
	if err := result.Error; err != nil {
		return 0, err
	}

	return chart.RowId, nil
}

func (r *chartRepository) GetCountOfRowID(chartId string, rowId int64) (int64, error) {
	var count int64
	err := r.db.Model(&models.Chart{}).Where("chart_id = ? AND row_id = ?", chartId, rowId).Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}
