package models

import (
	"time"
)

// Owner Table
type Owner struct {
	ID           uint          `gorm:"primaryKey" json:"id"`
	Name         string        `json:"name"`
	Email        string        `gorm:"unique;not null" json:"email"`
	Password     string        `json:"-"`
	CreatedAt    time.Time     `json:"created_at"`
	UpdatedAt    time.Time     `json:"updated_at"`
	Services     []Service     `gorm:"foreignKey:OwnerID"`
	Staff        []Staff       `gorm:"foreignKey:OwnerID"`
	Appointments []Appointment `gorm:"foreignKey:OwnerID"`
}

// Service Table
type Service struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"not null" json:"name"`
	Price     float64   `gorm:"not null" json:"price"`
	OwnerID   uint      `json:"owner_id"`
	Owner     Owner     `gorm:"foreignKey:OwnerID"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Staff Table
type Staff struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"not null" json:"name"`
	Role      string    `json:"role"`
	OwnerID   uint      `json:"owner_id"`
	Owner     Owner     `gorm:"foreignKey:OwnerID"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Appointment Table
type Appointment struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	DateTime  time.Time `gorm:"not null" json:"date_time"`
	Status    string    `gorm:"type:enum('pending','confirmed','completed');default:'pending'" json:"status"`
	OwnerID   uint      `json:"owner_id"`
	Owner     Owner     `gorm:"foreignKey:OwnerID"`
	ServiceID uint      `json:"service_id"`
	Service   Service   `gorm:"foreignKey:ServiceID"`
	StaffID   uint      `json:"staff_id"`
	Staff     Staff     `gorm:"foreignKey:StaffID"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
