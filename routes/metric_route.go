package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/soulter/tickstats/controllers"
	"github.com/soulter/tickstats/middlewares"
)

func RegisterMetricRoutes(router *gin.Engine, metricsController controllers.MetricsController) {
	accountRoutes := router.Group("/api/metric")
	{
		accountRoutes.POST(":appid", metricsController.Add)
	}
	accountRoutes.Use(middlewares.JWTAuthMiddleware(false))
	{
		accountRoutes.GET(":appid/:chartid", metricsController.Get)
	}
}
