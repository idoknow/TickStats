package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/soulter/tickstats/services"
	"github.com/soulter/tickstats/types"
)

type StatsController interface {
	GetPublicApps(c *gin.Context)
	GetPublicAppCharts(c *gin.Context)
}

type statsController struct {
	statsService services.StatsService
}

func NewStatsController(statsService services.StatsService) StatsController {
	return &statsController{
		statsService: statsService,
	}
}

func (controller *statsController) GetPublicApps(c *gin.Context) {
	var pager types.Pager
	if err := c.ShouldBindQuery(&pager); err != nil {
		// c.JSON(400, gin.H{"error": err.Error()})
		c.JSON(400, types.Result{
			Code:    400,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	if pager.Size == 0 {
		pager.Size = 10
	}
	if pager.Page == 0 {
		pager.Page = 1
	}

	apps, err := controller.statsService.GetPublicApps(pager.Page, pager.Size)
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
		Message: "success",
		Data:    apps,
	})
}

func (controller *statsController) GetPublicAppCharts(c *gin.Context) {
	appID := c.Param("appid")

	charts, err := controller.statsService.GetAppCharts(appID, true)
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
		Message: "success",
		Data:    charts,
	})
}
