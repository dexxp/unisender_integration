package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

const (
	importContacts = "/importContacts"
)

type UnisenderRequest struct {
	*Request
	apiKey string
}

type ImportContacts struct {
	FieldNames []string
	Data       [][]string
}

type ImportContactsRequest struct {
	apiKey string
	ImportContacts
}

type ImportContactsResponse struct {
	Result struct {
		Total     int `json:"total"`
		Inserted  int `json:"inserted"`
		Updated   int `json:"updated"`
		Deleted   int `json:"deleted"`
		NewEmails int `json:"new_emails"`
		Invalid   int `json:"invalid"`
		Log       []struct {
			Index   string `json:"index"`
			Code    string `json:"code"`
			Message string `json:"message"`
		} `json:"log"`
	} `json:"result"`
	Warnings []struct {
		Warning string `json:"warning"`
	} `json:"warnings"`
}

func NewUnisenderRequest(request *Request) *UnisenderRequest {
	return &UnisenderRequest{
		Request: request,
	}
}

func (u *UnisenderRequest) SetApiKey(apiKey string) {
	u.apiKey = apiKey
}

func (u *UnisenderRequest) ImportContacts(ic ImportContacts) (*ImportContactsResponse, error) {
	icReq := ImportContactsRequest{
		apiKey:         u.apiKey,
		ImportContacts: ic,
	}

	fmt.Println(icReq)

	params := icReq.Params()

	req, err := u.makeRequest(http.MethodPost, importContacts, strings.NewReader(params.Encode()))
	if err != nil {
		fmt.Println("ImportContacts: ", err)
		return nil, errors.New("Не удалось создать запрос")
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	body, err := u.Exec(req)

	fmt.Println("Body UNISENDER: ", string(body))

	if err != nil {
		fmt.Println("ImportContacts: ", err)
		return nil, errors.New("Не удалось выполнить запрос")
	}

	var icResp ImportContactsResponse

	err = json.Unmarshal(body, &icResp)

	if err != nil {
		fmt.Println("ImportContacts: ", err)
		return nil, errors.New("Не удалось преобразовать данные")
	}

	return &icResp, nil
}

func (icr *ImportContactsRequest) Params() url.Values {
	res := url.Values{}
	res.Set("api_key", icr.apiKey)
	for i, value := range icr.FieldNames {
		res.Set(fmt.Sprintf("field_names[%d]", i), value)
	}

	for i, line := range icr.Data {
		for j, value := range line {
			res.Set(fmt.Sprintf("data[%d][%d]", i, j), value)
		}
	}
	return res
}
