package entity

type Account struct {
	AccountID    uint64 `gorm:"primaryKey" json:"account_id"`
	Referer      string `json:"referer"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	Expires      int64  `json:"expires"`
	ApiKey       string `json:"api_key"`
	Disabled     bool   `json:"disabled"`
}
