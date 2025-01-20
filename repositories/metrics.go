package repositories

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/soulter/tickstats/models"
	"gorm.io/gorm"
)

type NumberMetricsRepository interface {
	Add(metric *models.BasicMetricData) error
	GetSimpleLine(appId string,
		keyName string,
		extraConfig map[string]interface{},
		from int64,
		to int64) ([]models.BasicMetricOutput, error)
	GetSimplePie(appId string,
		keyName string,
		extraConfig map[string]interface{},
		from int64,
		to int64) ([]models.BasicMetricOutput, error)
	GetTable(appId string,
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

func (r *metricsRepository) GetSimpleLine(
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

	// 检查 bucket_mins 是否存在
	_, ok := extraConfig["bucket_mins"]
	var bucket_mins int
	if !ok {
		bucket_mins = 30
	} else {
		bucket_mins, err = strconv.Atoi(extraConfig["bucket_mins"].(string))
		if err != nil || bucket_mins <= 30 {
			bucket_mins = 30
		}
	}

	data_points := (to - from) / 60 / int64(bucket_mins)

	fmt.Println("data_points: ", data_points)

	switch extraConfig["method"] {
	case "count":
		var _v_select string
		if extraConfig["distinct_ip"] == true {
			_v_select = "COUNT(DISTINCT ip) as v"
		} else {
			_v_select = "COUNT(ip) as v"
		}
		query := fmt.Sprintf(`
			SELECT time_bucket('%d minutes', time) as k, %s FROM basic_metric_data 
			WHERE app_id = ? AND jsonb_typeof(value->?) = 'number'
			GROUP BY k 
			ORDER BY k DESC
			LIMIT %d;
		`, bucket_mins, _v_select, data_points)
		err = r.db.Raw(query, appId, keyName).Scan(&metrics_).Error
	case "accumulate":
		query := fmt.Sprintf(`
			WITH RECURSIVE time_series AS (
				SELECT generate_series(
					date_trunc('hour', NOW()) - INTERVAL '30 days',
					date_trunc('hour', NOW() + INTERVAL '%d minutes'),
					'%d minutes'
				) AS time
			),
			metric_data AS (
				SELECT ts.time as k, COALESCE(SUM((value->>$1)::int), 0) as v
				FROM time_series ts
				LEFT JOIN basic_metric_data bmd
				ON bmd.time >= ts.time
				AND bmd.time < ts.time + INTERVAL '%d minutes'
				AND bmd.app_id = $2
				GROUP BY ts.time
				ORDER BY ts.time
			)
			SELECT k, SUM(v) OVER (ORDER BY k) as v
			FROM metric_data;
			ORDER BY k DESC
			LIMIT %d;
		`, bucket_mins, bucket_mins, bucket_mins, data_points)
		err = r.db.Raw(query, keyName, appId).Scan(&metrics_).Error
	default:
		// sum
		query := fmt.Sprintf(`
		SELECT time_bucket('%d minutes', time, NOW()) as k,
		SUM((value->>?)::numeric) as v FROM basic_metric_data 
		WHERE app_id = ? AND jsonb_typeof(value->?) = 'number' 
		GROUP BY k 
		ORDER BY k DESC
		LIMIT %d;`, bucket_mins, data_points)
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

func (r *metricsRepository) GetSimplePie(
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

func (r *metricsRepository) GetTable(
	appId string,
	keyName string,
	extraConfig map[string]interface{},
	from int64,
	to int64) ([]models.BasicMetricOutput, error) {
	var err error
	query := `
		SELECT EXTRACT(EPOCH FROM time)::bigint as k, ` + buildSelectColumns(keyName) + `
		FROM basic_metric_data
		WHERE app_id = ?
		AND time >= to_timestamp(?)
		AND time <= to_timestamp(?)
	`
	var metrics []models.BasicMetricOutput = []models.BasicMetricOutput{}

	rows, err := r.db.Raw(query, appId, from, to).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	for rows.Next() {

		values := make([]interface{}, len(columns))
		valuePtrs := make([]interface{}, len(columns))
		for i := range values {
			valuePtrs[i] = &values[i]
		}

		err = rows.Scan(valuePtrs...)
		if err != nil {
			return nil, err
		}
		// parse values
		var metric models.BasicMetricOutput
		for i, col := range columns {
			if col == "k" {
				metric.Key = values[i].(int64)
			} else {
				if metric.Value == nil {
					metric.Value = []interface{}{}
				}
				metric.Value = append(metric.Value.([]interface{}), values[i])
			}
		}
		metrics = append(metrics, metric)
	}

	return metrics, err
}

func buildSelectColumns(keyName string) string {
	// keyName 是逗号分隔的字符串
	keyNames := strings.Split(keyName, ",")

	columns := ""
	for i, key := range keyNames {
		if i > 0 {
			columns += ", "
		}
		columns += "value->>'" + key + "' as v" + strconv.Itoa(i)
	}
	return columns
}
