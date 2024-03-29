package repo

import (
	"fmt"
	"git.amocrm.ru/dmiroshnikov/unisender_integration/internal/entity"
	"git.amocrm.ru/dmiroshnikov/unisender_integration/pkg/mysql"
)

type AccountRepo struct {
	DB mysql.Database
}

func NewAccountRepo(db mysql.Database) *AccountRepo {
	return &AccountRepo{
		DB: db,
	}
}

func (repo *AccountRepo) Create(account *entity.Account) error {
	return repo.DB.DB.Create(account).Error
}
func (repo *AccountRepo) Remove(account *entity.Account) error {
	account.Disabled = true

	fmt.Println(account)

	return repo.DB.DB.Save(account).Error
}
func (repo *AccountRepo) Update(account *entity.Account) error {
	err := repo.DB.DB.Save(account).Error

	return err
}
func (repo *AccountRepo) Get(accountID uint64) (*entity.Account, error) {
	var account entity.Account
	err := repo.DB.DB.First(&account, accountID).Error

	return &account, err
}
