package usecases

import (
	"errors"
	"fmt"
	"git.amocrm.ru/dmiroshnikov/unisender_integration/api"
	"git.amocrm.ru/dmiroshnikov/unisender_integration/config"
	"git.amocrm.ru/dmiroshnikov/unisender_integration/internal/entity"
	"strconv"
)

type AmoClient interface {
	SetBaseUrl(url string)
	GetToken(auth *api.AuthRequest) (*api.AuthResponse, error)
	GetAccountID(accessToken string) (*api.AccountResponse, error)
	GetContacts(accessToken string) (*api.ContactsResponse, error)
}

type UnisenderClient interface {
	ImportContacts(ic api.ImportContacts) (*api.ImportContactsResponse, error)
	SetApiKey(apiKey string)
	SetBaseUrl(url string)
}

type UnisenderUseCase struct {
	ClientID        string
	accountRepo     AccountRepo
	integrationRepo IntegrationRepo
	amoClient       AmoClient
	UniClient       UnisenderClient
	contacts        []*entity.Contact
	emails          []*entity.Email
	account         *entity.Account
}

func NewUnisenderUseCase(repo AccountRepo, amoClient AmoClient, UniClient UnisenderClient, cfg config.Auth) *UnisenderUseCase {
	clientID := cfg.ClientID
	return &UnisenderUseCase{
		accountRepo: repo,
		amoClient:   amoClient,
		UniClient:   UniClient,
		ClientID:    clientID,
	}
}

func (u *UnisenderUseCase) SaveApiKeyInAccount(apiKey string, accountID uint64) error {
	account, err := u.accountRepo.Get(accountID)

	if err != nil {
		fmt.Println("Unisender SaveApiKeyInAccount: ", err)
		return errors.New("Не удалось получить аккаунт")
	}

	account.ApiKey = apiKey
	err = u.accountRepo.Update(account)

	u.account = account

	if err != nil {
		fmt.Println("Save api key in account: ", err)
		return errors.New("Не удалось сохранить API KEY в аккаунт")
	}

	return nil
}

func (u *UnisenderUseCase) GetContacts(accountID uint64) (api.ImportContacts, error) {
	u.amoClient.SetBaseUrl("https://" + u.account.Referer)
	contacts, _ := u.amoClient.GetContacts(u.account.AccessToken)

	// UPDATE TOKEN
	if contacts.Status == 401 {
		integration, err := u.integrationRepo.Get(u.ClientID, u.account.AccountID)
		if err != nil {

		}
		authReq := &api.AuthRequest{
			ClientId:     integration.ClientID,
			RefreshToken: u.account.RefreshToken,
			GrantType:    "refresh_token",
			ClientSecret: integration.SecretKey,
			RedirectUri:  integration.RedirectURL,
		}

		authResp, err := u.amoClient.GetToken(authReq)
		if err != nil {

		}

		u.account.AccessToken = authResp.AccessToken
		u.account.RefreshToken = authResp.RefreshToken
		u.account.Expires = authResp.ExpiresIn

		err = u.accountRepo.Update(u.account)

		if err != nil {

		}

		contacts, err = u.amoClient.GetContacts(u.account.AccessToken)

		if err != nil {

		}
	}

	u.contacts, u.emails = u.ConvertContactsAndEmails(contacts, u.account.AccountID)

	fmt.Println(u.contacts, u.emails)

	result := u.ConvertUnisenderContacts(u.emails)

	return result, nil
}

func (u *UnisenderUseCase) ConvertContactsAndEmails(contacts *api.ContactsResponse, accountID uint64) ([]*entity.Contact, []*entity.Email) {
	var contactEntities []*entity.Contact
	var emailEntities []*entity.Email
	for _, contact := range contacts.Embedded.Contacts {
		id := contact.ID
		name := contact.Name
		emails := make([]string, 0)
		for _, field := range contact.CustomFields {
			if field.FieldName == "Email" {
				for _, val := range field.Values {
					emails = append(emails, val.Value.(string))
				}
			}
		}
		for _, email := range emails {
			emailEntities = append(emailEntities, &entity.Email{
				Email:     email,
				Name:      name,
				Sync:      true,
				ContactID: id,
			})
		}

		contactEntities = append(contactEntities, &entity.Contact{
			ContactID: id,
			Name:      name,
			Sync:      true,
			AccountID: accountID,
		})
	}

	return contactEntities, emailEntities
}

func (u *UnisenderUseCase) ConvertUnisenderContacts(emails []*entity.Email) api.ImportContacts {
	data := make([][]string, len(emails))

	for i, email := range emails {
		data[i] = []string{email.Email, email.Name}
	}

	ic := api.ImportContacts{
		FieldNames: []string{"email", "Name"},
		Data:       data,
	}

	return ic
}

func (u *UnisenderUseCase) ImportContacts(ic api.ImportContacts) (*api.ImportContactsResponse, error) {
	u.UniClient.SetBaseUrl("https://api.unisender.com/ru/api")
	u.UniClient.SetApiKey(u.account.ApiKey)
	icResp, err := u.UniClient.ImportContacts(ic)

	for _, log := range icResp.Result.Log {
		index, _ := strconv.Atoi(log.Index)
		invalidEmail := u.emails[index]
		contactID := invalidEmail.ContactID
		invalidEmail.Sync = false

		for _, contact := range u.contacts {
			if contact.ContactID == contactID {
				contact.Sync = false
			}
		}
	}

	return icResp, err
}

func (u *UnisenderUseCase) CheckContacts(icResp *api.ImportContactsResponse) ([]*entity.Contact, []*entity.Email) {
	if icResp.Result.Inserted == 0 {
		return nil, nil
	}

	emails := make([]*entity.Email, icResp.Result.Inserted)
	for i, j := 0, 0; i < icResp.Result.Total; i++ {
		email := u.emails[i]
		if email.Sync {
			emails[j] = email
			j++
		}
	}

	fmt.Println(emails)

	return u.contacts, emails
}
