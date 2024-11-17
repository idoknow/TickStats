package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/soulter/tickstats/controllers"
	"github.com/soulter/tickstats/middlewares"
)

func RegisterAccountRoutes(router *gin.Engine, accountController controllers.AccountController) {
	accountRoutes := router.Group("/api/account")
	{
		accountRoutes.POST("register", accountController.Register) // Register a new account
		accountRoutes.POST("login", accountController.Login)       // Login
	}

	accountRoutes.Use(middlewares.JWTAuthMiddleware(true))
	{
		accountRoutes.GET("app", accountController.GetApplications)                      // Get all applications
		accountRoutes.POST("app/new", accountController.CreateApplication)               // Create a new application
		accountRoutes.DELETE("app/:appid", accountController.DeleteApplication)          // Delete an application
		accountRoutes.PUT("app/:appid", accountController.UpdateApplication)             // Update an application
		accountRoutes.POST("app/:appid/chart/new", accountController.CreateChart)        // Create a new line chart
		accountRoutes.DELETE("app/:appid/chart/:chartid", accountController.DeleteChart) // Delete a chart
		accountRoutes.PUT("app/:appid/chart/:chartid", accountController.UpdateChart)    // Update a chart
		accountRoutes.GET("app/:appid/chart", accountController.GetCharts)               // Get all charts
		accountRoutes.GET("auth", accountController.GetAuth)                             // Get auth info
		accountRoutes.GET("logout", accountController.Logout)                            // Logout
	}
}
