package api

import (
	"errors"
	"fmt"
	"io"
	"net/http"
)

type Request struct {
	client  *http.Client
	baseUrl string
}

func NewRequest(client *http.Client) *Request {
	return &Request{
		client: client,
	}
}

func (r *Request) SetBaseUrl(url string) {
	r.baseUrl = url
}

func (r *Request) makeRequest(method, resource string, reader io.Reader) (*http.Request, error) {
	endpoint := r.baseUrl + resource
	fmt.Println("makeRequest endpoint: ", endpoint)
	req, err := http.NewRequest(method, endpoint, reader)

	if err != nil {
		fmt.Println(err)
		return nil, errors.New("Не удалось создать запрос")
	}

	return req, nil
}

func (r *Request) Exec(req *http.Request) ([]byte, error) {
	resp, err := r.client.Do(req)

	if err != nil {
		fmt.Println(err)
		return nil, errors.New("Не удалось выполнить запрос")
	}

	defer resp.Body.Close()

	out, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
		return nil, errors.New("Не удалось прочитать тело ответа")
	}

	return out, nil
}
