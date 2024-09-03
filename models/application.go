package models

type Application struct {
	AppId       string `json:"app_id"`
	Name        string `json:"name" binding:"required,min=1,max=128"`
	Emoji       string `json:"emoji" binding:"required,min=1,max=16"`
	Public      *bool  `json:"public" binding:"required"`
	AccountId   int    `json:"-"`
	CreatedTime string `json:"created_time"`
	UpdatedTime string `json:"updated_time"`
}
