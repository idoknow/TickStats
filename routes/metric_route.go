package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/soulter/tickstats/controllers"
)

func RegisterMetricRoutes(router *gin.Engine, metricsController controllers.MetricsController) {
	accountRoutes := router.Group("/api/metric")
	{
		accountRoutes.POST(":appid", metricsController.Add)
		accountRoutes.GET(":appid", metricsController.Get)
	}
}
