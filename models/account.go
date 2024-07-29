package models

type Account struct {
	AccountId   int    `json:"-" gorm:"primaryKey"`
	Name        string `json:"name" binding:"required"`
	Email       string `json:"email" binding:"required"`
	Password    string `json:"password" binding:"required"`
	CreatedTime string `json:"-"`
	UpdatedTime string `json:"-"`
}

type Credential struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
