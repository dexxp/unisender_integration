package usecases

import (
	"errors"
	"fmt"
	"git.amocrm.ru/dmiroshnikov/unisender_integration/api"
	"git.amocrm.ru/dmiroshnikov/unisender_integration/config"
	"git.amocrm.ru/dmiroshnikov/unisender_integration/internal/entity"
)

type AuthUseCase struct {
	accountRepo     AccountRepo
	integrationRepo IntegrationRepo
	client          AmoClient
	cfg             config.Auth
}

func NewAuthUseCase(accountRepo AccountRepo, integrationRepo IntegrationRepo, client AmoClient, cfg config.Auth) *AuthUseCase {
	return &AuthUseCase{
		accountRepo:     accountRepo,
		integrationRepo: integrationRepo,
		client:          client,
		cfg:             cfg,
	}
}

func (a *AuthUseCase) GetToken(code string, referer string) error {
	authReq := &api.AuthRequest{
		RedirectUri:  a.cfg.RedirectURI,
		GrantType:    "authorization_code",
		ClientSecret: a.cfg.ClientSecret,
		ClientId:     a.cfg.ClientID,
		Code:         code,
	}

	fmt.Println(authReq)

	baseUrl := fmt.Sprintf("https://%s", referer)

	a.client.SetBaseUrl(baseUrl)

	authResp, err := a.client.GetToken(authReq)

	fmt.Println("Токен получен! ", authResp)

	if err != nil {
		fmt.Println(err)
		return errors.New("Ошибка получения токена")
	}

	accountResp, err := a.client.GetAccountID(authResp.AccessToken)

	if err != nil {
		fmt.Println(err)
		return errors.New("Ошибка получения токена")
	}

	account := &entity.Account{
		AccountID:    accountResp.AccountID,
		Referer:      referer,
		AccessToken:  authResp.AccessToken,
		RefreshToken: authResp.RefreshToken,
		Expires:      authResp.ExpiresIn,
		Disabled:     false,
	}

	err = a.CreateAccount(account)

	if err != nil {
		fmt.Println(err)
		return errors.New("Ошибка создания аккаунта")
	}

	integration := &entity.Integration{
		ClientID:           authReq.ClientId,
		SecretKey:          authReq.ClientSecret,
		RedirectURL:        authReq.RedirectUri,
		AuthenticationCode: authReq.Code,
		AccountID:          account.AccountID,
	}

	err = a.CreateIntegration(integration)

	if err != nil {
		fmt.Println(err)
		return errors.New("Ошибка создания интеграции")
	}

	return nil
}

func (a *AuthUseCase) CreateAccount(account *entity.Account) error {
	err := a.accountRepo.Create(account)

	if err != nil {
		fmt.Println(err)
		return errors.New("Не удалось создать аккаунт")
	}

	fmt.Println("Account created! ", account)

	return nil
}

func (a *AuthUseCase) CreateIntegration(integration *entity.Integration) error {
	err := a.integrationRepo.Create(integration)

	if err != nil {
		fmt.Println("Create integration: ", err)
		return errors.New("Не получилось создать интеграцию")
	}

	fmt.Println("Новая интеграция: ", integration)

	return nil
}
