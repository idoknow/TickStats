package services

import (
	"github.com/gin-gonic/gin"
	"github.com/soulter/tickstats/models"
	"github.com/soulter/tickstats/repositories"
	"github.com/soulter/tickstats/utils"
)

type MetricsService interface {
	Add(metric models.BasicMetricData) error
	GetByAppID(c *gin.Context, string, chartId string, from int64, to int64) ([]models.BasicMetricOutput, error)
}

type metricsService struct {
	metricsRepository     repositories.NumberMetricsRepository
	chartRepository       repositories.ChartRepository
	applicationRepository repositories.ApplicationRepository
}

func NewMetricsService(metricsRepository repositories.NumberMetricsRepository,
	chartRepository repositories.ChartRepository,
	appRepository repositories.ApplicationRepository) MetricsService {
	return &metricsService{
		metricsRepository:     metricsRepository,
		chartRepository:       chartRepository,
		applicationRepository: appRepository,
	}
}

func (s *metricsService) Add(metric models.BasicMetricData) error {
	return s.metricsRepository.Add(&metric)
}

func (s *metricsService) GetByAppID(c *gin.Context, appId string, chartId string, from int64, to int64) ([]models.BasicMetricOutput, error) {
	var metrics []models.BasicMetricOutput
	var err error

	chart, err := s.chartRepository.FindByChartID(chartId)

	if err != nil {
		return nil, err
	}

	if !chart.Public {
		// auth
		if isAuth, _ := c.Get("isAuthorized"); !isAuth.(bool) {
			return nil, utils.ErrUnauthorized
		}
		// check if the user owns the app
		userId, _ := c.Get("userID")
		accountID := int(userId.(float64))
		app, err := s.applicationRepository.FindByAppID(appId)
		if err != nil {
			return nil, err
		}
		if app.AccountId != accountID {
			return nil, utils.ErrUnauthorized
		}
	}

	// from and to must be both set or both unset
	if (from == 0 && to != 0) || (from != 0 && to == 0) {
		return nil, utils.ErrInvalidTimeRange
	}
	if from != 0 && to != 0 && from > to {
		return nil, utils.ErrInvalidTimeRange
	}
	if from == 0 && to == 0 {
		// 1 days interval
		to = utils.GetCurrentTimestamp()
		from = to - 60*60*24
	} else {
		// cannot query more than 7 days
		if from < utils.GetCurrentTimestamp()-60*60*24*7 {
			return nil, utils.ErrInvalidTimeRange
		}
		if to-from > 60*60*24*7 {
			return nil, utils.ErrTimeRangeTooLong
		}
	}

	switch chart.ChartType {
	case models.SimpleLine:
		metrics, err = s.metricsRepository.GetSimpleLine(appId, chart.KeyName, chart.ExtraConfig, from, to)
	case models.SimplePie:
		metrics, err = s.metricsRepository.GetSimplePie(appId, chart.KeyName, chart.ExtraConfig, from, to)
	case models.Table:
		metrics, err = s.metricsRepository.GetTable(appId, chart.KeyName, chart.ExtraConfig, from, to)
		
	default:
		return nil, utils.ErrInvalidChartType
	}

	if err != nil {
		return nil, err
	}

	return metrics, nil
}
