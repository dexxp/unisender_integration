package usecases

import (
	"errors"
	"fmt"
	"git.amocrm.ru/dmiroshnikov/unisender_integration/internal/entity"
)

type ContactUseCase struct {
	contactRepo ContactRepo
}

func NewContactUseCase(contactRepo ContactRepo) *ContactUseCase {
	return &ContactUseCase{
		contactRepo: contactRepo,
	}
}

func (c *ContactUseCase) CreateContacts(contacts []*entity.Contact) error {
	err := c.contactRepo.Create(contacts)

	if err != nil {
		fmt.Println("Create Contacts: ", err)
		return errors.New("Не получилось создать контакт")
	}
	return nil
}

func (c *ContactUseCase) GetContacts(accountID uint64, which string) ([]*entity.Contact, error) {
	contacts, err := c.contactRepo.GetContactsByAccountID(accountID)
	result := make([]*entity.Contact, 0)

	if err != nil {
		fmt.Println("Get Contacts: ", err)
		return nil, errors.New("Не удалось получить контакты")
	}

	switch which {
	case "sync":
		for _, contact := range contacts {
			if contact.Sync {
				result = append(result, contact)
			}
		}
	case "nosync":
		for _, contact := range contacts {
			if !contact.Sync {
				result = append(result, contact)
			}
		}
	default:
		result = contacts
	}

	return result, nil
}
