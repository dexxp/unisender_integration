package usecases

import (
	"errors"
	"fmt"
	"git.amocrm.ru/dmiroshnikov/unisender_integration/internal/entity"
)

type IntegrationUseCase struct {
	repo IntegrationRepo
}

func NewIntegrationUseCase(repo IntegrationRepo) Integration {
	return &IntegrationUseCase{
		repo: repo,
	}
}

func (i *IntegrationUseCase) Get(clientID string, accountID uint64) (*entity.Integration, error) {
	integration, err := i.repo.Get(clientID, accountID)

	if err != nil {
		fmt.Println("Get Integration: ", err)
		return nil, errors.New("Не удалось получить интеграцию")
	}

	return integration, nil
}

func (i *IntegrationUseCase) GetAllIntegrationsAccount(accountID uint64) ([]*entity.Integration, error) {
	integrations, err := i.repo.GetIntegrationByAccountID(accountID)

	if err != nil {
		fmt.Println("Get All integrations account: ", err)
		return nil, errors.New("Не удалось получить все интеграции аккаунта")
	}

	return integrations, nil
}
