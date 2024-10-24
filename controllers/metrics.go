package controllers

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/soulter/tickstats/models"
	"github.com/soulter/tickstats/services"
	"github.com/soulter/tickstats/types"
	"github.com/soulter/tickstats/utils"
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
		c.JSON(400, types.Result{
			Code:    400,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	current_time := time.Now()

	clientHost := utils.GetRemoteHost(c.Request)

	metricData.AppId = appId
	metricData.IP = clientHost
	metricData.Time = current_time
	metricData.Value = metricInput.MetricsData

	if err := controller.metricsService.Add(metricData); err != nil {
		c.JSON(400, types.Result{
			Code:    400,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	c.JSON(200, types.Result{
		Code:    200,
		Message: "success",
		Data:    nil,
	})
}

func (controller *metricsController) Get(c *gin.Context) {
	appId := c.Param("appid")
	chartId := c.Param("chartid")

	// keyName := c.Query("key_name")
	// chartType := c.Query("chart_type")

	if appId == "" || chartId == "" {
		c.JSON(400, types.Result{
			Code:    400,
			Message: "Invalid request",
			Data:    nil,
		})
		return
	}

	metrics, err := controller.metricsService.GetByAppID(c, appId, chartId)
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
		Message: "Get metrics success",
		Data:    metrics,
	})
}
