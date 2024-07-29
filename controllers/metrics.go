package controllers

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/soulter/tickstats/models"
	"github.com/soulter/tickstats/services"
)

type MetricsController interface {
	Add(c *gin.Context)
	Get(c *gin.Context)
}

type metricsController struct {
	metricsService services.MetricsService
}

func NewMetricsController(metricsService services.MetricsService) MetricsController {
	return &metricsController{
		metricsService: metricsService,
	}
}

func (controller *metricsController) Add(c *gin.Context) {
	var metricInput models.BasicMetricInput
	var metricData models.BasicMetricData

	appId := c.Param("appid")

	if err := c.ShouldBindJSON(&metricInput); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	current_time := time.Now()
	client_ip := c.ClientIP()

	metricData.AppId = appId
	metricData.IP = client_ip
	metricData.Time = current_time
	metricData.Value = metricInput.MetricsData

	if err := controller.metricsService.Add(metricData); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
}

func (controller *metricsController) Get(c *gin.Context) {
	appId := c.Param("appid")

	keyName := c.Query("key_name")
	chartType := c.Query("chart_type")

	metrics, err := controller.metricsService.GetByAppID(appId, chartType, keyName)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, metrics)
}
