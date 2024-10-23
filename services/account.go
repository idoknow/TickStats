package services

import (
	"github.com/soulter/tickstats/models"
	"github.com/soulter/tickstats/repositories"
	"github.com/soulter/tickstats/utils"
)

type AccountService interface {
	Authenticate(email, password string) (*models.Account, error)
	Register(name, email, password string) error
	CreateApplication(application *models.Application) error
	DeleteApplication(accountId int, appId string) error
	GetApplications(accountId int) ([]models.Application, error)
	CreateChart(models.Chart) error
	DeleteChart(account_id int, chartId string) error
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

func (service *accountService) Authenticate(cred, password string) (*models.Account, error) {
	account, err := service.accountRepository.FindByCredientials(cred)
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
	if _, err := service.accountRepository.FindByCredientials(email); err == nil {
		return utils.ErrEmailTaken
	}
	// Check if the name is already taken
	if _, err := service.accountRepository.FindByCredientials(name); err == nil {
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

func (service *accountService) CreateApplication(application *models.Application) error {
	application.AppId = utils.GenerateUUID("")
	application.CreatedTime = utils.CurrentTime()
	application.UpdatedTime = utils.CurrentTime()

	// Save the application to the database
	if err := service.applicationRepository.Create(application); err != nil {
		return err
	}

	return nil
}

func (service *accountService) DeleteApplication(accountId int, appId string) error {
	// Find the application by application ID
	application, err := service.applicationRepository.FindByAppID(appId)
	if err != nil {
		return err
	}
	if application == nil {
		return utils.ErrAppNotFound
	}

	// Check if the application belongs to the account
	if application.AccountId != accountId {
		return utils.ErrUnauthorized
	}

	// Delete the application from the database
	if err := service.applicationRepository.Delete(appId); err != nil {
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

func (service *accountService) CreateChart(chart models.Chart) error {

	// Generate a new chart ID
	chart.ChartId = utils.GenerateUUID("chart")
	chart.CreatedTime = utils.CurrentTime()
	chart.UpdatedTime = utils.CurrentTime()

	// Save the chart to the database
	if err := service.chartRepository.Create(&chart); err != nil {
		return err
	}

	return nil
}

func (service *accountService) DeleteChart(accountId int, chartId string) error {

	// Find the chart by chart ID
	chart, err := service.chartRepository.FindByChartID(chartId)
	if err != nil {
		return err
	}

	// Find the application by application ID
	application, err := service.applicationRepository.FindByAppID(chart.AppId)

	if err != nil {
		return err
	}

	if application.AccountId != accountId {
		return utils.ErrUnauthorized
	}

	// Delete the chart from the database
	if err := service.chartRepository.Delete(chartId); err != nil {
		return err
	}

	return nil
}

func (service *accountService) GetAccount(accountId int) (*models.Account, error) {
	// Find the account by account ID
	account, err := service.accountRepository.FindByAccountID(accountId)
	if err != nil {
		return nil, err
	}

	return account, nil
}
