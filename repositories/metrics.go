package repositories

import (
	"fmt"
	"time"

	"github.com/soulter/tickstats/models"
	"gorm.io/gorm"
)

type NumberMetricsRepository interface {
	Add(metric *models.BasicMetricData) error
	GetPlainNumberVal(appId string,
		keyName string,
		extraConfig map[string]interface{},
		from int64,
		to int64) ([]models.BasicMetricOutput, error)
	GetPlainStringVal(appId string,
		keyName string,
		extraConfig map[string]interface{},
		from int64,
		to int64) ([]models.BasicMetricOutput, error)
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
	extraConfig map[string]interface{},
	from int64,
	to int64) ([]models.BasicMetricOutput, error) {
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

	switch extraConfig["method"] {
	case "count":
		var _v_select string
		if extraConfig["distinct_ip"] == true {
			_v_select = "COUNT(DISTINCT ip) as v"
		} else {
			_v_select = "COUNT(ip) as v"
		}
		query := fmt.Sprintf(`
			SELECT time_bucket('30 minutes', time) as k, %s FROM basic_metric_data 
			WHERE app_id = ? AND jsonb_typeof(value->?) = 'number'
			GROUP BY k 
			ORDER BY k 
			LIMIT 1440;
		`, _v_select)
		err = r.db.Raw(query, appId, keyName).Scan(&metrics_).Error
	case "accumulate":
		query := `
        WITH RECURSIVE time_series AS (
            SELECT generate_series(
                date_trunc('hour', NOW()) - INTERVAL '30 day',
                date_trunc('hour', NOW() + INTERVAL '30 minutes'),
                '30 minutes'
            ) AS time
        ),
        metric_data AS (
            SELECT ts.time as k, COALESCE(SUM((value->>?)::int), 0) as v
            FROM time_series ts
            LEFT JOIN basic_metric_data bmd
            ON bmd.time >= ts.time
            AND bmd.time < ts.time + INTERVAL '30 minutes'
            AND bmd.app_id = ?
            GROUP BY ts.time
            ORDER BY ts.time
        )
        SELECT k, SUM(v) OVER (ORDER BY k) as v
        FROM metric_data;
		`
		err = r.db.Raw(query, keyName, appId).Scan(&metrics_).Error
	default:
		// sum
		query := `
		SELECT time_bucket('30 minutes', time) as k,
		SUM((value->>?)::numeric) as v FROM basic_metric_data 
		WHERE app_id = ? AND jsonb_typeof(value->?) = 'number' 
		GROUP BY k 
		ORDER BY k LIMIT 1440;`
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
	extraConfig map[string]interface{},
	from int64,
	to int64) ([]models.BasicMetricOutput, error) {
	var err error

	query := `
		SELECT value->>? as k, 
		COUNT(*) as v
		FROM basic_metric_data
		WHERE app_id = ?
		AND jsonb_typeof(value->?) = 'string'
		AND time >= to_timestamp(?)
		AND time <= to_timestamp(?)
		GROUP BY k
	`
	var metrics []models.BasicMetricOutput = []models.BasicMetricOutput{}
	err = r.db.Raw(query, keyName, appId, keyName, from, to).Scan(&metrics).Error
	return metrics, err
}
