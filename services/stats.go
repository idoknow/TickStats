package services

import (
	"github.com/soulter/tickstats/models"
	"github.com/soulter/tickstats/repositories"
	"github.com/soulter/tickstats/types"
	"github.com/soulter/tickstats/utils"
)

type StatsService interface {
	GetPublicApps(page, size int) ([]models.Application, error)
	GetAppCharts(appId string, onlyPublic bool) (*types.ChartResponse, error)
}

type statsService struct {
	applicationRepository repositories.ApplicationRepository
	chartRepository       repositories.ChartRepository
	accountRepository     repositories.AccountRepository
}

func NewStatsService(
	applicationRepository repositories.ApplicationRepository,
	chartRepository repositories.ChartRepository,
	accountRepository repositories.AccountRepository,
) StatsService {
	return &statsService{
		applicationRepository: applicationRepository,
		chartRepository:       chartRepository,
		accountRepository:     accountRepository,
	}
}

func (service *statsService) GetPublicApps(page, size int) ([]models.Application, error) {
	return service.applicationRepository.FindPublicByPage(page, size)
}

func (service *statsService) GetAppCharts(appId string, onlyPublic bool) (*types.ChartResponse, error) {
	var chartResponse types.ChartResponse

	app, err := service.applicationRepository.FindByAppID(appId)
	if err != nil {
		return nil, err
	}
	if app == nil {
		return nil, utils.ErrAppNotFound
	}
	account, err := service.accountRepository.FindByAccountID(app.AccountId)
	if err != nil {
		return nil, err
	}
	if account == nil {
		return nil, utils.ErrAccountNotFound
	}

	lineCharts, err := service.chartRepository.FindByAppID(appId, onlyPublic)
	if err != nil {
		return nil, err
	}
	chartResponse.Chart = lineCharts
	chartResponse.AppName = app.Name
	chartResponse.AccountName = account.Name

	return &chartResponse, nil
}
