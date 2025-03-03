package models

type Staffs struct {
	ID          int64  `gorm:"primaryKey;autoIncrement"`
	OwnerId     int64  `gorm:"not null;index" json:"ownerId"`
	Owner       Users  `gorm:"foreignKey:OwnerId; references:ID" json:"owner"`
	Name        string `gorm:"not null" json:"name"`
	Role        string `gorm:"not null" json:"role"`
	PhoneNumber string `gorm:"not null;type:varchar(13)" json:"phoneNumber"`
	Email       string `gorm:"not null" json:"email"`
	CreatedAt   int64  `gorm:"createAtTime" json:"createdAt"`
	UpdatedAt   int64  `gorm:"updatedAtTime" json:"updatedAt"`
}
