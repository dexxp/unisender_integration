package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

const (
	oauth    = "/oauth2/access_token"
	contacts = "/api/v4/contacts"
	account  = "/api/v4/account"
)

type AmocrmRequest struct {
	*Request
}

type AuthRequest struct {
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	GrantType    string `json:"grant_type"`
	Code         string `json:"code"`
	RefreshToken string `json:"refresh_token"`
	RedirectUri  string `json:"redirect_uri"`
}

type AuthResponse struct {
	TokenType    string `json:"token_type"`
	ExpiresIn    int64  `json:"expires_in"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type AccountResponse struct {
	AccountID uint64 `json:"id"`
	Status    int    `json:"status"`
}

type ContactValue struct {
	Value interface{} `json:"value"`
}

type CustomField struct {
	FieldID   int            `json:"field_id"`
	FieldName string         `json:"field_name"`
	Values    []ContactValue `json:"values"`
}

type Contact struct {
	ID           uint64        `json:"id"`
	Name         string        `json:"name"`
	CustomFields []CustomField `json:"custom_fields_values"`
}

type ContactsResponse struct {
	Embedded struct {
		Contacts []Contact `json:"contacts"`
	} `json:"_embedded"`
	Status int `json:"status"`
}

func NewAmocrmRequest(request *Request) *AmocrmRequest {
	return &AmocrmRequest{
		Request: request,
	}
}

func (r *AmocrmRequest) GetToken(auth *AuthRequest) (*AuthResponse, error) {
	fmt.Println(auth)
	jsonBody, err := json.Marshal(auth)

	if err != nil {
		fmt.Println(err)
		return nil, errors.New("Не удалось преобразовать данные в JSON")
	}

	req, err := r.makeRequest(http.MethodPost, oauth, bytes.NewReader(jsonBody))

	if err != nil {
		fmt.Println(err)
		return nil, errors.New("Не удалось создать запрос")
	}

	req.Header.Add("User-Agent", "amoCRM-oAuth-client/1.0")
	req.Header.Add("Content-Type", "application/json")

	body, err := r.Exec(req)

	if err != nil {
		fmt.Println(err)
		return nil, errors.New("Не удалось выполнить запрос")
	}

	var authResp AuthResponse

	fmt.Println("Auth body: ", string(body))
	fmt.Println("Auth resp: ", authResp)

	err = json.Unmarshal(body, &authResp)

	return &authResp, nil
}

func (r *AmocrmRequest) GetContacts(accessToken string) (*ContactsResponse, error) {
	var contactsResp ContactsResponse

	body, err := r.get(contacts, accessToken)

	if err != nil {
		fmt.Println(err)
		return nil, errors.New("Ошибка получения данных")
	}

	err = json.Unmarshal(body, &contactsResp)

	if err != nil {
		fmt.Println(err)
		return nil, errors.New("Ошибка преобразования данных")
	}

	return &contactsResp, nil
}

func (r *AmocrmRequest) GetAccountID(accessToken string) (*AccountResponse, error) {
	var accountResp AccountResponse

	body, err := r.get(account, accessToken)

	if err != nil {
		fmt.Println(err)
		return nil, errors.New("Ошибка получения данных")
	}

	err = json.Unmarshal(body, &accountResp)

	if err != nil {
		fmt.Println(err)
		return nil, errors.New("Ошибка преобразования данных")
	}

	return &accountResp, nil
}

func (r *AmocrmRequest) get(resource, accessToken string) ([]byte, error) {
	req, err := r.makeRequest(http.MethodGet, resource, nil)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	token := fmt.Sprintf("Bearer %v", accessToken)
	req.Header.Set("Authorization", token)

	body, err := r.Exec(req)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return body, nil
}

func (r *AmocrmRequest) CheckAuth(v any, auth *AuthRequest) error {
	return nil
}
