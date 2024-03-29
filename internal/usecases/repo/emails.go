package repo

import (
	"git.amocrm.ru/dmiroshnikov/unisender_integration/internal/entity"
	"git.amocrm.ru/dmiroshnikov/unisender_integration/pkg/mysql"
)

type EmailRepo struct {
	DB mysql.Database
}

func NewEmailRepo(db mysql.Database) *EmailRepo {
	return &EmailRepo{
		DB: db,
	}
}

func (repo *EmailRepo) Create(emails []*entity.Email) error {
	return repo.DB.DB.Create(emails).Error
}
func (repo *EmailRepo) Delete(emails []*entity.Email) error {
	for _, email := range emails {
		email.ContactID = 0
	}

	err := repo.DB.DB.Omit("contact_id").Updates(emails).Error

	return err
}
func (repo *EmailRepo) Update(email *entity.Email) error {
	panic("implement me")
}
func (repo *EmailRepo) Get(email string) (*entity.Email, error) {
	var emailEntity entity.Email
	err := repo.DB.DB.First(&emailEntity, email).Error

	return &emailEntity, err
}
func (repo *EmailRepo) GetEmailsByContactID(contactID uint64) ([]*entity.Email, error) {
	var emails []*entity.Email

	result := repo.DB.DB.Where("contact_id = ?", contactID).Find(emails)

	return emails, result.Error
}
