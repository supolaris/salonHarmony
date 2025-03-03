package services

import (
	"github.com/supolaris/salonHarmony/db"
	"github.com/supolaris/salonHarmony/models"
)

func CreateUser(user *models.Users) error {
	result := db.DB.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
