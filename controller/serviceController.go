package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/supolaris/salonHarmony/models"
	"github.com/supolaris/salonHarmony/services"
)

func CreateService(ctx *gin.Context) {
	var service models.Services
	err := ctx.ShouldBindJSON(&service)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "unable to bind service data",
		})
		return
	}

	// var user models.Users
	//   if err := db.DB.First(&user, service.OwnerId).Error; err != nil {
	//       ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid ownerId"})
	//       return
	//   }

	err = services.CreateService(&service)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "unable to create service" + err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"result": service,
	})
}
