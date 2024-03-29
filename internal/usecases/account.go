package usecases

import (
	"errors"
	"fmt"
	"git.amocrm.ru/dmiroshnikov/unisender_integration/internal/entity"
)

type AccountUseCase struct {
	repo AccountRepo
}

func NewAccountUseCase(accountRepo AccountRepo) *AccountUseCase {
	return &AccountUseCase{
		repo: accountRepo,
	}
}

func (a *AccountUseCase) Get(accountID uint64) (*entity.Account, error) {
	account, err := a.repo.Get(accountID)

	if err != nil {
		fmt.Println("Account get: ", err)
		return nil, errors.New("Не удалось получить аккаунт")
	}

	return account, nil
}

func (a *AccountUseCase) Update(account *entity.Account) error {
	err := a.repo.Update(account)

	if err != nil {
		fmt.Println("Account Update: ", err)
		return errors.New("Не удалось обновить аккаунт")
	}

	return nil
}
