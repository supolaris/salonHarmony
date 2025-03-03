package models

type Users struct {
	ID           int64          `gorm:"primaryKey;autoIncrement" json:"id"`
	Name         string         `gorm:"not null" json:"name"`
	Email        string         `gorm:"not null" json:"email"`
	Password     string         `gorm:"not null" json:"password"`
	PhoneNumber  string         `gorm:"not null;type:varchar(13)" json:"phoneNumber"`
	Address      string         `gorm:"not null" json:"address"`
	CreatedAt    int64          `gorm:"autoCreatedAt" json:"createdAt"`
	UpdatedAt    int64          `gorm:"autoUpdateTime" json:"updatedAt"`
	Services     []Services     `gorm:"foreignKey:OwnerId" json:"services"`
	Staffs       []Staffs       `gorm:"foreignKey:OwnerId" json:"staffs"`
	Appointments []Appointments `gorm:"foreignKey:OwnerId" json:"appointments"`
}
