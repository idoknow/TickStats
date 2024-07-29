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

// // Client input
// type NumberMetricInput struct {
// 	MetricsData map[string]float64 `json:"metrics_data"` // key-value pair
// }

// type StringMetricInput struct {
// 	MetricsData map[string]string `json:"metrics_data"` // key-value pair
// }

// // interact with the TimeScale database
// type NumberMetricData struct {
// 	Time    time.Time `json:"time"`
// 	AppId   string    `json:"appid"`
// 	KeyName string    `json:"key_name"`
// 	Value   float64   `json:"value"`
// 	IP      string    `json:"ip"`
// }

// type StringMetricData struct {
// 	Time    time.Time `json:"time"`
// 	AppId   string    `json:"appid"`
// 	KeyName string    `json:"key_name"`
// 	Value   string    `json:"value"`
// 	IP      string    `json:"ip"`
// }

// output
// type NumberMetricOutput struct {
// 	Bucket time.Time `json:"time"`
// 	Value  float64   `json:"value"`
// }

// type StringMetricOutput struct {
// 	Bucket time.Time `json:"time"`
// 	Value  string    `json:"value"`
// }
