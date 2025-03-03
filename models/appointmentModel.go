package models

import "time"

type Appointments struct {
	ID              int64      `gorm:"primaryKey;autoIncrement" json:"id"`
	OwnerId         int64      `gorm:"not null;index" json:"ownerId"`
	Owner           Users      `gorm:"foreignKey:OwnerId;references:ID" json:"owner"`
	ClientName      string     `gorm:"not null" json:"clientName"`
	ClientPhone     string     `json:"clientPhone"`
	ServiceId       int64      `gorm:"not null,index" json:"serviceId"`
	Service         []Services `gorm:"foreignKey:ServiceId;references:ID" json:"service"`
	StaffId         int64      `gorm:"not null,index" json:"staffId"`
	Staff           []Staffs   `gorm:"foreignKey:StaffId;references:ID" json:"staff"`
	AppointmentTime time.Time  `gorm:"not null" json:"time"`
	Status          string     `gorm:"not null" json:"status"`
	CreatedAt       int64      `gorm:"createAtTime" json:"createdAt"`
	UpdatedAt       int64      `gorm:"updatedAtTime" json:"updatedAt"`
}
