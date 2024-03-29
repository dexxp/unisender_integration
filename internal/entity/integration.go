package entity

type Integration struct {
	ClientID           string `gorm:"primaryKey" json:"client_id"`
	SecretKey          string `json:"secret_key"`
	RedirectURL        string `json:"redirect_url"`
	AuthenticationCode string `json:"authentication_code"`
	AccountID          uint64 `gorm:"foreignKey:AccountID" json:"account_id"`
}
