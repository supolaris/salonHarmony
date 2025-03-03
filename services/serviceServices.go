package services

import (
	"github.com/supolaris/salonHarmony/db"
	"github.com/supolaris/salonHarmony/models"
)

func CreateService(service *models.Services) error {

	result := db.DB.Create(&service)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
