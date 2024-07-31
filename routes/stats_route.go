package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/soulter/tickstats/controllers"
)

func RegisterStatsRoutes(router *gin.Engine, statsController controllers.StatsController) {
	accountRoutes := router.Group("/api/stats")
	{
		accountRoutes.GET("apps", statsController.GetPublicApps)               // Get all public apps
		accountRoutes.GET(":appid/charts", statsController.GetPublicAppCharts) // Get all public charts in this app
	}
}
