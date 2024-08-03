package repositories

import (
	"fmt"

	"github.com/soulter/tickstats/models"
	"gorm.io/gorm"
)

type NumberMetricsRepository interface {
	Add(metric *models.BasicMetricData) error
	GetPlainNumberVal(appId string, keyName string) ([]models.BasicMetricOutput, error)
	GetPlainStringVal(appId string, keyName string) ([]models.BasicMetricOutput, error)
}

type metricsRepository struct {
	db *gorm.DB
}

func NewMetricsRepository(db *gorm.DB) NumberMetricsRepository {
	return &metricsRepository{db}
}

func (r *metricsRepository) Add(metric *models.BasicMetricData) error {
	return r.db.Create(metric).Error
}

func (r *metricsRepository) GetPlainNumberVal(appId string, keyName string) ([]models.BasicMetricOutput, error) {
	var err error

	query := `
		SELECT time_bucket('30 minutes', time) as k,
		SUM((value->>?)::numeric) as v
		FROM basic_metric_data
		WHERE app_id = ?
		AND jsonb_typeof(value->?) = 'number'
		GROUP BY k
		ORDER BY k
		LIMIT 1440;
	`
	var metrics []models.BasicMetricOutput = []models.BasicMetricOutput{}
	err = r.db.Raw(query, keyName, appId, keyName).Scan(&metrics).Error
	fmt.Println(metrics)
	return metrics, err
}

func (r *metricsRepository) GetPlainStringVal(appId string, keyName string) ([]models.BasicMetricOutput, error) {
	var err error

	query := `
		SELECT value->>? as k, 
		COUNT(*) as v
		FROM basic_metric_data
		WHERE app_id = ?
		AND jsonb_typeof(value->?) = 'string'
		AND time > NOW() - INTERVAL '1 hour'
		GROUP BY k
	`
	var metrics []models.BasicMetricOutput = []models.BasicMetricOutput{}
	err = r.db.Raw(query, keyName, appId, keyName).Scan(&metrics).Error
	return metrics, err
}
