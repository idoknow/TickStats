package models

import "time"

type BasicMetricInput struct {
	MetricsData map[string]interface{} `json:"metrics_data"` // key-value pair
}

type BasicMetricData struct {
	Time  time.Time              `json:"time"`
	AppId string                 `json:"appid"`
	Value map[string]interface{} `json:"value" gorm:"type:jsonb"`
	IP    string                 `json:"ip"`
}

type BasicMetricOutput struct {
	Key   interface{} `json:"k" gorm:"column:k"`
	Value interface{} `json:"v" gorm:"column:v"`
}