package models

type Chart struct {
	ChartId     string                 `json:"-"`
	AppId       string                 `json:"appid" binding:"required,min=8"`
	ChartName   string                 `json:"chart_name" binding:"required,min=1,max=128"`
	ChartType   string                 `json:"chart_type" binding:"required,min=1,max=128"`
	KeyName     string                 `json:"key_name" binding:"required,min=1,max=256"`
	ExtraConfig map[string]interface{} `json:"extra_config"`
	Description string                 `json:"description"`
	Public      bool                   `json:"public"`
	CreatedTime string                 `json:"-"`
	UpdatedTime string                 `json:"-"`
}
