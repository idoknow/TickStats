package services

import (
	"github.com/soulter/tickstats/models"
	"github.com/soulter/tickstats/repositories"
	"github.com/soulter/tickstats/utils"
)

type AccountService interface {
	Authenticate(email, password string) (*models.Account, error)
	Register(name, email, password string) error
	CreateApplication(accountId int, name string) error
	GetApplications(accountId int) ([]models.Application, error)
	CreateChart(appId string, name string, chartType string, keyName string) error
	GetCharts(appId string) ([]models.Chart, error)
	GetAccount(accountId int) (*models.Account, error)
}

type accountService struct {
	accountRepository     repositories.AccountRepository
	applicationRepository repositories.ApplicationRepository
	chartRepository       repositories.ChartRepository
}

func NewAccountService(
	accountRepository repositories.AccountRepository,
	applicationRepository repositories.ApplicationRepository,
	chartRepository repositories.ChartRepository,
) AccountService {
	return &accountService{
		accountRepository:     accountRepository,
		applicationRepository: applicationRepository,
		chartRepository:       chartRepository,
	}
}

func (service *accountService) Authenticate(email, password string) (*models.Account, error) {
	// Find the account by email
	account, err := service.accountRepository.FindByEmail(email)
	if err != nil {
		return nil, err
	}
	if account.Password != password {
		return nil, utils.ErrInvalidPassword
	}

	return account, nil
}

func (service *accountService) Register(name, email, password string) error {
	// Check if the email is already taken
	if _, err := service.accountRepository.FindByEmail(email); err == nil {
		return utils.ErrEmailTaken
	}
	// Check if the name is already taken
	if _, err := service.accountRepository.FindByName(name); err == nil {
		return utils.ErrNameTaken
	}

	// Create a new account
	account := models.Account{
		Name:        name,
		Email:       email,
		Password:    password,
		CreatedTime: utils.CurrentTime(),
		UpdatedTime: utils.CurrentTime(),
	}

	// Save the account to the database
	if err := service.accountRepository.Create(&account); err != nil {
		return err
	}

	return nil
}

func (service *accountService) CreateApplication(accountId int, name string) error {
	// Create a new application
	application := models.Application{
		AppId:       utils.GenerateUUID(""),
		Name:        name,
		AccountId:   accountId,
		CreatedTime: utils.CurrentTime(),
		UpdatedTime: utils.CurrentTime(),
	}

	// Save the application to the database
	if err := service.applicationRepository.Create(&application); err != nil {
		return err
	}

	return nil
}

func (service *accountService) GetApplications(accountId int) ([]models.Application, error) {
	// Find the applications by account ID
	applications, err := service.applicationRepository.FindByAccountID(accountId)
	if err != nil {
		return nil, err
	}

	return applications, nil
}

func (service *accountService) CreateChart(appId string, name string, chartType string, keyName string) error {
	// Create a new line chart
	lineChart := models.Chart{
		ChartId:     utils.GenerateUUID("chart_"),
		ChartName:   name,
		ChartType:   chartType,
		AppId:       appId,
		KeyName:     keyName,
		CreatedTime: utils.CurrentTime(),
		UpdatedTime: utils.CurrentTime(),
	}

	// Save the line chart to the database
	if err := service.chartRepository.Create(&lineChart); err != nil {
		return err
	}

	return nil
}

func (service *accountService) GetCharts(appId string) ([]models.Chart, error) {
	// Find the line charts by application ID

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

	lineCharts, err := service.chartRepository.FindByAppID(appId)
	if err != nil {
		return nil, err
	}
	for i := range lineCharts {
		lineCharts[i].AppName = app.Name
		lineCharts[i].AccountName = account.Name
	}

	return lineCharts, nil
}

func (service *accountService) GetAccount(accountId int) (*models.Account, error) {
	// Find the account by account ID
	account, err := service.accountRepository.FindByAccountID(accountId)
	if err != nil {
		return nil, err
	}

	return account, nil
}
