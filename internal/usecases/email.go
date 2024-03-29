package usecases

import (
	"errors"
	"fmt"
	"git.amocrm.ru/dmiroshnikov/unisender_integration/internal/entity"
)

type EmailUseCase struct {
	repo EmailRepo
}

func NewEmailUseCase(repo EmailRepo) *EmailUseCase {
	return &EmailUseCase{
		repo: repo,
	}
}

func (e *EmailUseCase) CreateEmails(emails []*entity.Email) error {
	err := e.repo.Create(emails)

	if err != nil {
		fmt.Println("Create emails", err)
		return errors.New("Не получилось создать контакты")
	}

	return nil
}

func (e *EmailUseCase) GetEmailsByContactID(contactID uint64) ([]*entity.Email, error) {
	emails, err := e.repo.GetEmailsByContactID(contactID)

	if err != nil {
		fmt.Println("GetEmailsByContactID: ", err)
		return nil, errors.New("Не удалось получить emails")
	}

	return emails, nil
}
