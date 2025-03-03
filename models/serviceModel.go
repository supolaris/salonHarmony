package models

type Services struct {
	ID          int64   `gorm:"primaryKey;autoIncrement" json:"id"`
	OwnerId     int64   `gorm:"not null;index" json:"ownerId"`
	Owner       Users   `gorm:"foreignKey:OwnerId;references:ID" json:"owner"`
	Name        string  `gorm:"not null" json:"name"`
	Description string  `json:"description"`
	Price       float64 `gorm:"not null" json:"price"`
	Duration    int64   `gorm:"not null" json:"duration"`
	CreatedAt   int64   `gorm:"createAtTime" json:"createdAt"`
	UpdatedAt   int64   `gorm:"updatedAtTime" json:"updatedAt"`
}
