package models

import (
	"github.com/soulter/tickstats/utils"
)

type ChartType string

const (
	SimpleLine ChartType = "simple_line"
	SimplePie  ChartType = "simple_pie"
	Table      ChartType = "table"
)

type Chart struct {
	ChartId     string        `json:"chart_id" gorm:"primaryKey"`
	AppId       string        `json:"appid" validate:"nonzero,min=8"`
	ChartName   string        `json:"chart_name" validate:"nonzero,min=1,max=128"`
	ChartType   ChartType     `json:"chart_type" validate:"nonzero,min=1,max=128"`
	KeyName     string        `json:"key_name" validate:"nonzero,min=1,max=256,regexp=^[a-zA-Z0-9_]+(\\,[a-zA-Z0-9_]+)*$"`
	ExtraConfig utils.JSONMap `json:"extra_config" gorm:"type:text"`
	Description string        `json:"description"`
	Public      bool          `json:"public"`
	RowId       int64         `json:"row_id"`
	CreatedTime string        `json:"created_time"`
	UpdatedTime string        `json:"updated_time"`
}

var ChartUpdatableFields = []string{"chart_name", "chart_type", "key_name", "extra_config", "description", "public", "row_id", "updated_time"}
