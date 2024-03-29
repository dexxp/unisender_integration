package entity

type Email struct {
	Email     string `gorm:"primaryKey" json:"email"`
	Name      string `json:"name"`
	Sync      bool
	ContactID uint64 `gorm:"foreignKey:ContactID" json:"contact_id"`
}
