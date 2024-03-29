package unisender

import (
	"io"
	"net/http"
	"net/url"
	"strconv"
)

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func (i *Impl) FirstSync(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Ошибка чтения тела запроса", http.StatusBadRequest)
		return
	}
	values, err := url.ParseQuery(string(body))
	accountID, err := strconv.Atoi(values.Get("account_id"))

	if err != nil {
		http.Error(w, "Неверный аккаунт ID", http.StatusBadRequest)
		return
	}

	unisenderKey := values.Get("unisender_key")

	err = i.unisenderUseCase.SaveApiKeyInAccount(unisenderKey, uint64(accountID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	contacts, err := i.unisenderUseCase.GetContacts(uint64(accountID))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	icResp, err := i.unisenderUseCase.ImportContacts(contacts)

	contactsToSave, emailsToSave := i.unisenderUseCase.CheckContacts(icResp)

	err = i.contactUseCase.CreateContacts(contactsToSave)
	err = i.emailUseCase.CreateEmails(emailsToSave)
}
