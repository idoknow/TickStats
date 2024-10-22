package types

import (
	"github.com/soulter/tickstats/models"
)

type Result struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ChartResponse struct {
	Chart       []models.Chart `json:"chart"`
	AppName     string         `json:"app_name"`
	AccountName string         `json:"account_name"`
}

var NotAuthorizedResult = Result{
	Code:    401,
	Message: "Not authorized",
	Data:    nil,
}
