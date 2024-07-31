package models

type Chart struct {
	ChartId     string `json:"-"`
	AppId       string `json:"appid" binding:"required,min=8"`
	AppName     string `json:"app_name" gorm:"-"`
	AccountName string `json:"account_name" gorm:"-"`
	ChartName   string `json:"chart_name" binding:"required,min=1,max=128"`
	ChartType   string `json:"chart_type" binding:"required,min=1,max=128"`
	KeyName     string `json:"key_name" binding:"required,min=1,max=256"`
	Description string `json:"description" binding:"required,min=1,max=512"`
	Public      bool   `json:"public" binding:"required"`
	CreatedTime string `json:"-"`
	UpdatedTime string `json:"-"`
}
