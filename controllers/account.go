package controllers

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/soulter/tickstats/config"
	"github.com/soulter/tickstats/models"
	"github.com/soulter/tickstats/services"
)

type AccountController interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
	CreateApplication(c *gin.Context)
	GetApplications(c *gin.Context)
	CreateChart(c *gin.Context)
	GetCharts(c *gin.Context)
	GetAuth(c *gin.Context)
}

type accountController struct {
	accountService services.AccountService
}

func NewAccountController(accountService services.AccountService) AccountController {
	return &accountController{
		accountService: accountService,
	}
}

func (controller *accountController) Register(c *gin.Context) {
	// json
	var account models.Account
	if err := c.ShouldBindJSON(&account); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := controller.accountService.Register(account.Name, account.Email, account.Password); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
}

func (controller *accountController) Login(c *gin.Context) {

	var credentials models.Credential
	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	account, err := controller.accountService.Authenticate(credentials.Username, credentials.Password)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
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

	// set cookie
	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

func (controller *accountController) GetAuth(c *gin.Context) {
	userId, _ := c.Get("userID")
	accountId := int(userId.(float64))

	account, err := controller.accountService.GetAccount(accountId)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, account)
}

func (controller *accountController) CreateApplication(c *gin.Context) {
	var application models.Application
	if err := c.ShouldBindJSON(&application); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	userId, _ := c.Get("userID")
	accountId := int(userId.(float64))

	if err := controller.accountService.CreateApplication(accountId, application.Name); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
}

func (controller *accountController) GetApplications(c *gin.Context) {
	userId, _ := c.Get("userID")
	accountId := int(userId.(float64))

	applications, err := controller.accountService.GetApplications(accountId)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, applications)
}

func (controller *accountController) CreateChart(c *gin.Context) {
	appId := c.Param("appid")
	var chart models.Chart

	if err := c.ShouldBindJSON(&chart); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// get application by account id
	userId, _ := c.Get("userID")
	accountId := int(userId.(float64))

	applications, err := controller.accountService.GetApplications(accountId)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
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
		c.JSON(400, gin.H{"error": "Application does not exist"})
		return
	}

	if err := controller.accountService.CreateChart(appId, chart.ChartName, chart.ChartType, chart.KeyName); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
}

func (controller *accountController) GetCharts(c *gin.Context) {
	appId := c.Param("appid")

	charts, err := controller.accountService.GetCharts(appId)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, charts)
}
