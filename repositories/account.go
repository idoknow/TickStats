package repositories

import (
	"github.com/soulter/tickstats/models"
	"gorm.io/gorm"
)

type AccountRepository interface {
	Create(account *models.Account) error
	FindByEmail(email string) (*models.Account, error)
	FindByName(name string) (*models.Account, error)
	FindByAccountID(accountId int) (*models.Account, error)
}

type accountRepository struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) AccountRepository {
	return &accountRepository{db}
}

func (r *accountRepository) Create(account *models.Account) error {
	return r.db.Create(account).Error
}

func (r *accountRepository) FindByEmail(email string) (*models.Account, error) {
	var account models.Account
	if err := r.db.Where("email = ?", email).First(&account).Error; err != nil {
		return nil, err
	}
	return &account, nil
}

func (r *accountRepository) FindByName(name string) (*models.Account, error) {
	var account models.Account
	if err := r.db.Where("name = ?", name).First(&account).Error; err != nil {
		return nil, err
	}
	return &account, nil
}

func (r *accountRepository) FindByAccountID(accountId int) (*models.Account, error) {
	var account models.Account
	if err := r.db.Where("account_id = ?", accountId).Select("account_id", "name", "email").First(&account).Error; err != nil {
		return nil, err
	}
	return &account, nil

}
