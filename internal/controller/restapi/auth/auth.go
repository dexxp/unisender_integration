package auth

import (
	"net/http"
)

func (i *Impl) FistAuth(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	code, referer := query.Get("code"), query.Get("referer")

	err := i.authUseCase.GetToken(code, referer)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.WriteHeader(200)
	w.Write([]byte("Аккаунт и интеграция успешно созданы!"))
}
