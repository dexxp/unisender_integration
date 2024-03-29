package repo

import (
	"git.amocrm.ru/dmiroshnikov/unisender_integration/internal/entity"
	"git.amocrm.ru/dmiroshnikov/unisender_integration/pkg/mysql"
)

type IntegrationRepo struct {
	DB mysql.Database
}

func NewIntegrationRepo(db mysql.Database) *IntegrationRepo {
	return &IntegrationRepo{
		DB: db,
	}
}

func (repo *IntegrationRepo) Create(integration *entity.Integration) error {
	return repo.DB.DB.Create(integration).Error
}
func (repo *IntegrationRepo) Remove(clientID string, accountID uint64) error {
	var integration entity.Integration

	err := repo.DB.DB.Where("client_id = ? AND account_id = ?", clientID, accountID).First(&integration).Error

	if err != nil {
		return err
	}

	err = repo.DB.DB.Delete(&integration).Error

	return err
}
func (repo *IntegrationRepo) GetIntegrationByAccountID(accountID uint64) ([]*entity.Integration, error) {
	var integrations []*entity.Integration

	result := repo.DB.DB.
		Where("account_id = ?", accountID).
		Find(&integrations)

	return integrations, result.Error
}

func (repo *IntegrationRepo) Get(clientID string, accountID uint64) (*entity.Integration, error) {
	var integration entity.Integration

	err := repo.DB.DB.
		Where("client_id = ? AND account_id = ?", clientID, accountID).
		First(&integration).
		Error

	return &integration, err
}
