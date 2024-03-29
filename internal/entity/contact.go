package entity

type Contact struct {
	ContactID uint64 `gorm:"primaryKey" json:"contact_id"`
	Name      string `json:"name"`
	Sync      bool   `json:"bool"`
	AccountID uint64 `gorm:"foreignKey:AccountID" json:"account_id"`
}
