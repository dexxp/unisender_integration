package usecases

import (
	"git.amocrm.ru/dmiroshnikov/unisender_integration/api"
	"git.amocrm.ru/dmiroshnikov/unisender_integration/internal/entity"
)

type (
	// Account -.
	Account interface {
		Get(accountID uint64) (*entity.Account, error)
		Update(account *entity.Account) error
	}
	// AccountRepo -.
	AccountRepo interface {
		Create(account *entity.Account) error
		Remove(account *entity.Account) error
		Get(accountID uint64) (*entity.Account, error)
		Update(account *entity.Account) error
	}

	// Integration -.
	Integration interface {
		Get(clientID string, accountID uint64) (*entity.Integration, error)
	}
	// IntegrationRepo -.
	IntegrationRepo interface {
		Create(integration *entity.Integration) error
		Remove(clientID string, accountID uint64) error
		GetIntegrationByAccountID(accountID uint64) ([]*entity.Integration, error)
		Get(clientID string, accountID uint64) (*entity.Integration, error)
	}

	// Contact -.
	Contact interface {
		CreateContacts(contacts []*entity.Contact) error
		GetContacts(accountID uint64, which string) ([]*entity.Contact, error)
	}
	// ContactRepo -.
	ContactRepo interface {
		Create(contacts []*entity.Contact) error
		Get(contactID uint64) (*entity.Contact, error)
		GetContactsByAccountID(accountID uint64) ([]*entity.Contact, error)
		Updates(contacts []*entity.Contact) error
		Remove(contacts []*entity.Contact) error
		RemoveContactsByAccountID(accountID uint64) error
	}

	Email interface {
		CreateEmails(emails []*entity.Email) error
		GetEmailsByContactID(contactID uint64) ([]*entity.Email, error)
	}

	// EmailRepo -.
	EmailRepo interface {
		Create(email []*entity.Email) error
		Delete(email []*entity.Email) error
		Update(email *entity.Email) error
		Get(email string) (*entity.Email, error)
		GetEmailsByContactID(contactID uint64) ([]*entity.Email, error)
	}

	// Auth -.
	Auth interface {
		CreateIntegration(integration *entity.Integration) error
		CreateAccount(account *entity.Account) error
		GetToken(code string, referer string) error
	}

	// Unisender -.
	Unisender interface {
		SaveApiKeyInAccount(apiKey string, accountID uint64) error
		GetContacts(accountID uint64) (api.ImportContacts, error)
		ImportContacts(ic api.ImportContacts) (*api.ImportContactsResponse, error)
		CheckContacts(icResp *api.ImportContactsResponse) ([]*entity.Contact, []*entity.Email)
	}

	Disable interface {
		Disable(accountID uint64) error
	}
)
