package services

import (
	"github.com/soulter/tickstats/models"
	"github.com/soulter/tickstats/repositories"
	"github.com/soulter/tickstats/utils"
)

type MetricsService interface {
	Add(metric models.BasicMetricData) error
	GetByAppID(appId string, chartType string, keyName string) ([]models.BasicMetricOutput, error)
}

type metricsService struct {
	metricsRepository repositories.NumberMetricsRepository
	chartRepository   repositories.ChartRepository
}

func NewMetricsService(metricsRepository repositories.NumberMetricsRepository, chartRepository repositories.ChartRepository) MetricsService {
	return &metricsService{
		metricsRepository: metricsRepository,
		chartRepository:   chartRepository,
	}
}

func (s *metricsService) Add(metric models.BasicMetricData) error {
	return s.metricsRepository.Add(&metric)
}

func (s *metricsService) GetByAppID(appId string, chartType string, keyName string) ([]models.BasicMetricOutput, error) {
	var metrics []models.BasicMetricOutput
	var err error

	if keyName == "" {
		return nil, utils.ErrInvalidKeyName
	}

	charts, err := s.chartRepository.FindByAppID(appId, false)

	if err != nil {
		return nil, err
	}

	for _, chart := range charts {
		if chart.KeyName == keyName && chart.ChartType == chartType {
			if chartType == "simple_line" || chartType == "simple_bar" {
				metrics, err = s.metricsRepository.GetPlainNumberVal(appId, keyName, chart.ExtraConfig)
			} else if chartType == "simple_pie" {
				metrics, err = s.metricsRepository.GetPlainStringVal(appId, keyName, chart.ExtraConfig)
			} else {
				return nil, utils.ErrInvalidChartType
			}
		}
	}

	if err != nil {
		return nil, err
	}

	return metrics, nil
}
