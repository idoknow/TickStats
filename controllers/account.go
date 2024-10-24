package controllers

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/soulter/tickstats/config"
	"github.com/soulter/tickstats/models"
	"github.com/soulter/tickstats/services"
	"github.com/soulter/tickstats/types"
)

type AccountController interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
	CreateApplication(c *gin.Context)
	DeleteApplication(c *gin.Context)
	GetApplications(c *gin.Context)
	CreateChart(c *gin.Context)
	DeleteChart(c *gin.Context)
	UpdateChart(c *gin.Context)
	GetCharts(c *gin.Context)
	GetAuth(c *gin.Context)
	Logout(c *gin.Context)
}

type accountController struct {
	accountService services.AccountService
	statsService   services.StatsService
}

func NewAccountController(
	accountService services.AccountService,
	statsService services.StatsService) AccountController {
	return &accountController{
		accountService: accountService,
		statsService:   statsService,
	}
}

func (controller *accountController) Register(c *gin.Context) {
	// json
	var account models.Account
	if err := c.ShouldBindJSON(&account); err != nil {
		c.JSON(400, types.Result{
			Code:    400,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	if err := controller.accountService.Register(account.Name, account.Email, account.Password); err != nil {
		c.JSON(400, types.Result{
			Code:    400,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	c.JSON(200, types.Result{
		Code:    200,
		Message: "Register success",
		Data:    nil,
	})
}

func (controller *accountController) Login(c *gin.Context) {

	var credentials models.Credential
	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(400, types.Result{
			Code:    400,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	account, err := controller.accountService.Authenticate(credentials.Username, credentials.Password)
	if err != nil {
		c.JSON(400, types.Result{
			Code:    400,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	expirationTime := time.Now().Add(config.JWTExpirationDuration())
	claims := &jwt.MapClaims{
		"userID": account.AccountId,
		"exp":    expirationTime.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(config.JWTSecret)
	if err != nil {
		c.JSON(400, types.Result{
			Code:    400,
			Message: "Failed to generate token",
			Data:    nil,
		})
		return
	}

	domain := c.Request.Header.Get("Origin")

	if gin.Mode() == gin.DebugMode {
		http.SetCookie(c.Writer, &http.Cookie{
			Name:     "token",
			Value:    tokenString,
			Path:     "/",
			Domain:   domain,
			Secure:   true,
			SameSite: http.SameSiteNoneMode,
			HttpOnly: false,
			MaxAge:   int(expirationTime.Unix()),
		})
	} else {
		http.SetCookie(c.Writer, &http.Cookie{
			Name:     "token",
			Value:    tokenString,
			Path:     "/",
			Domain:   domain,
			Secure:   false,
			SameSite: http.SameSiteStrictMode,
			HttpOnly: false,
			MaxAge:   int(expirationTime.Unix()),
		})
	}

	c.JSON(200, types.Result{
		Code:    200,
		Message: "Login success",
		Data:    nil,
	})
}

func (controller *accountController) GetAuth(c *gin.Context) {
	userId, _ := c.Get("userID")
	accountId := int(userId.(float64))

	var ret struct {
		AccountId    int                  `json:"account_id"`
		Email        string               `json:"email"`
		Name         string               `json:"name"`
		Applications []models.Application `json:"apps"`
		Token        string               `json:"token"`
	}

	account, err := controller.accountService.GetAccount(accountId)
	if err != nil {
		c.JSON(400, types.Result{
			Code:    400,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	ret.AccountId = account.AccountId
	ret.Email = account.Email
	ret.Name = account.Name
	ret.Applications, _ = controller.accountService.GetApplications(accountId)

	c.JSON(200, types.Result{
		Code:    200,
		Message: "Get account success",
		Data:    ret,
	})
}

func (controller *accountController) CreateApplication(c *gin.Context) {
	var application models.Application
	if err := c.ShouldBindJSON(&application); err != nil {
		c.JSON(400, types.Result{
			Code:    400,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	userId, _ := c.Get("userID")
	accountId := int(userId.(float64))
	application.AccountId = accountId

	if err := controller.accountService.CreateApplication(&application); err != nil {
		c.JSON(400, types.Result{
			Code:    400,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	c.JSON(200, types.Result{
		Code:    200,
		Message: "Create application success",
		Data:    nil,
	})
}

func (controller *accountController) DeleteApplication(c *gin.Context) {
	appId := c.Param("appid")

	userId, _ := c.Get("userID")
	accountId := int(userId.(float64))

	if err := controller.accountService.DeleteApplication(accountId, appId); err != nil {
		c.JSON(400, types.Result{
			Code:    400,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	c.JSON(200, types.Result{
		Code:    200,
		Message: "Delete application success",
		Data:    nil,
	})
}

func (controller *accountController) GetApplications(c *gin.Context) {
	userId, _ := c.Get("userID")
	accountId := int(userId.(float64))

	applications, err := controller.accountService.GetApplications(accountId)
	if err != nil {
		c.JSON(400, types.Result{
			Code:    400,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	c.JSON(200, types.Result{
		Code:    200,
		Message: "Get applications success",
		Data:    applications,
	})
}

func (controller *accountController) CreateChart(c *gin.Context) {
	appId := c.Param("appid")
	var chart models.Chart

	if err := c.ShouldBindJSON(&chart); err != nil {
		c.JSON(400, types.Result{
			Code:    400,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	// get application by account id
	userId, _ := c.Get("userID")
	accountId := int(userId.(float64))

	applications, err := controller.accountService.GetApplications(accountId)
	if err != nil {
		c.JSON(400, types.Result{
			Code:    400,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}
	// check if the application exists
	var exists bool
	for _, app := range applications {
		if app.AppId == appId {
			exists = true
			break
		}
	}
	if !exists {
		c.JSON(400, types.Result{
			Code:    400,
			Message: "Application does not exist",
			Data:    nil,
		})
		return
	}

	if err := controller.accountService.CreateChart(chart); err != nil {
		c.JSON(400, types.Result{
			Code:    400,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	c.JSON(200, types.Result{
		Code:    200,
		Message: "Create chart success",
		Data:    nil,
	})
}

func (controller *accountController) DeleteChart(c *gin.Context) {
	chartId := c.Param("chartid")

	userId, _ := c.Get("userID")
	accountId := int(userId.(float64))

	if err := controller.accountService.DeleteChart(accountId, chartId); err != nil {
		c.JSON(400, types.Result{
			Code:    400,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	c.JSON(200, types.Result{
		Code:    200,
		Message: "Delete chart success",
		Data:    nil,
	})
}

func (controller *accountController) UpdateChart(c *gin.Context) {
	chartId := c.Param("chartid")
	var chart models.Chart

	if err := c.ShouldBindJSON(&chart); err != nil {
		c.JSON(400, types.Result{
			Code:    400,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	chart.ChartId = chartId

	userId, _ := c.Get("userID")
	accountId := int(userId.(float64))

	if err := controller.accountService.UpdateChart(accountId, &chart); err != nil {
		c.JSON(400, types.Result{
			Code:    400,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	c.JSON(200, types.Result{
		Code:    200,
		Message: "Update chart success",
		Data:    nil,
	})
}

func (controller *accountController) GetCharts(c *gin.Context) {
	appId := c.Param("appid")

	charts, err := controller.statsService.GetAppCharts(appId, false)
	if err != nil {
		c.JSON(400, types.Result{
			Code:    400,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	c.JSON(200, types.Result{
		Code:    200,
		Message: "Get charts success",
		Data:    charts,
	})
}

func (controller *accountController) Logout(c *gin.Context) {
	domain := c.Request.Header.Get("Origin")

	if gin.Mode() == gin.DebugMode {
		http.SetCookie(c.Writer, &http.Cookie{
			Name:     "token",
			Value:    ":(",
			Path:     "/",
			Domain:   domain,
			Secure:   true,
			SameSite: http.SameSiteNoneMode,
			HttpOnly: false,
			MaxAge:   0,
		})
	} else {
		http.SetCookie(c.Writer, &http.Cookie{
			Name:     "token",
			Value:    ":(",
			Path:     "/",
			Domain:   domain,
			Secure:   false,
			SameSite: http.SameSiteStrictMode,
			HttpOnly: false,
			MaxAge:   0,
		})
	}

	c.JSON(200, types.Result{
		Code:    200,
		Message: "Logout success",
		Data:    nil,
	})
}
