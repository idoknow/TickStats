package utils

import (
	"fmt"
)

var (
	ErrInvalidPassword  = fmt.Errorf("invalid password")
	ErrEmailTaken       = fmt.Errorf("the email is already taken")
	ErrNameTaken        = fmt.Errorf("the name is already taken")
	ErrInvalidKeyName   = fmt.Errorf("invalid key name")
	ErrInvalidChartType = fmt.Errorf("invalid chart type")
	ErrAppNotFound      = fmt.Errorf("application not found")
	ErrAccountNotFound  = fmt.Errorf("account not found")
	ErrUnauthorized     = fmt.Errorf("unauthorized")
	ErrInvalidTimeRange = fmt.Errorf("invalid time range")
	ErrTimeRangeTooLong = fmt.Errorf("time range too long")
	ErrChartNotFound    = fmt.Errorf("chart not found")
)
