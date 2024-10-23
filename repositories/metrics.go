package repositories

import (
	"time"

	"github.com/soulter/tickstats/models"
	"gorm.io/gorm"
)

type NumberMetricsRepository interface {
	Add(metric *models.BasicMetricData) error
	GetPlainNumberVal(appId string,
		keyName string,
		extraConfig map[string]interface{}) ([]models.BasicMetricOutput, error)
	GetPlainStringVal(appId string,
		keyName string,
		extraConfig map[string]interface{}) ([]models.BasicMetricOutput, error)
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

func (r *metricsRepository) GetPlainNumberVal(
	appId string,
	keyName string,
	extraConfig map[string]interface{}) ([]models.BasicMetricOutput, error) {
	var err error
	type TimeMetrics struct {
		Time  time.Time `json:"k" gorm:"column:k"`
		Value float64   `json:"v" gorm:"column:v"`
	}
	var metrics []models.BasicMetricOutput = []models.BasicMetricOutput{}
	var metrics_ []TimeMetrics = []TimeMetrics{}

	// extraConfig:
	// method: sum, count, accumulate
	// distinct_ip: true, false

	query := `
		SELECT time_bucket('30 minutes', time) as k,
	`

	switch extraConfig["method"] {
	case "count":
		query += `COUNT(*) as v FROM basic_metric_data WHERE app_id = ? AND jsonb_typeof(value->?) = 'number' GROUP BY k ORDER BY k LIMIT 1440;`
		err = r.db.Raw(query, appId, keyName).Scan(&metrics_).Error
	case "accumulate":
		query += `SUM((value->>?)::numeric) OVER (ORDER BY time) as v FROM basic_metric_data WHERE app_id = ? AND jsonb_typeof(value->?) = 'number' GROUP BY k,value->>?,time ORDER BY k LIMIT 1440;`
		err = r.db.Raw(query, keyName, appId, keyName, keyName).Scan(&metrics_).Error
	default:
		// sum
		query += `SUM((value->>?)::numeric) as v FROM basic_metric_data WHERE app_id = ? AND jsonb_typeof(value->?) = 'number' GROUP BY k ORDER BY k LIMIT 1440;`
		err = r.db.Raw(query, keyName, appId, keyName).Scan(&metrics_).Error
	}

	for _, metric := range metrics_ {
		metrics = append(metrics, models.BasicMetricOutput{
			Key:   metric.Time.UnixMilli(),
			Value: metric.Value,
		})
	}

	return metrics, err
}

func (r *metricsRepository) GetPlainStringVal(
	appId string,
	keyName string,
	extraConfig map[string]interface{}) ([]models.BasicMetricOutput, error) {
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
