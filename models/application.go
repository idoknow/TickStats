package models

type Application struct {
	AppId       string `json:"app_id"`
	Name        string `json:"name"`
	AccountId   int    `json:"-"`
	CreatedTime string `json:"created_time"`
	UpdatedTime string `json:"updated_time"`
}
