package repo

import (
	"git.amocrm.ru/dmiroshnikov/unisender_integration/internal/entity"
	"git.amocrm.ru/dmiroshnikov/unisender_integration/pkg/mysql"
)

type ContactRepo struct {
	DB mysql.Database
}

func NewContactRepo(db mysql.Database) *ContactRepo {
	return &ContactRepo{
		DB: db,
	}
}

func (repo *ContactRepo) Create(contacts []*entity.Contact) error {
	return repo.DB.DB.Create(contacts).Error
}
func (repo *ContactRepo) Get(contactID uint64) (*entity.Contact, error) {
	var contact entity.Contact
	err := repo.DB.DB.First(&contact, contactID).Error

	return &contact, err
}
func (repo *ContactRepo) GetContactsByAccountID(accountID uint64) ([]*entity.Contact, error) {
	var contacts []*entity.Contact

	result := repo.DB.DB.Where("account_id = ?", accountID).Find(contacts)

	return contacts, result.Error
}
func (repo *ContactRepo) Updates(contacts []*entity.Contact) error {
	panic("implement me")
}
func (repo *ContactRepo) Remove(contacts []*entity.Contact) error {
	for _, contact := range contacts {
		contact.AccountID = 0
	}

	err := repo.DB.DB.Omit("account_id").Updates(contacts).Error

	return err
}
func (repo *ContactRepo) RemoveContactsByAccountID(accountID uint64) error {
	panic("implement me")
}
