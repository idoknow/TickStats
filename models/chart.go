package models

import "github.com/soulter/tickstats/utils"

type Chart struct {
	ChartId     string        `json:"chart_id" gorm:"primaryKey"`
	AppId       string        `json:"appid" binding:"required,min=8"`
	ChartName   string        `json:"chart_name" binding:"required,min=1,max=128"`
	ChartType   string        `json:"chart_type" binding:"required,min=1,max=128"`
	KeyName     string        `json:"key_name" binding:"required,min=1,max=256"`
	ExtraConfig utils.JSONMap `json:"extra_config" gorm:"type:text"`
	Description string        `json:"description"`
	Public      bool          `json:"public"`
	CreatedTime string        `json:"created_time"`
	UpdatedTime string        `json:"updated_time"`
}
