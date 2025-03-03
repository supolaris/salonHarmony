package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/supolaris/salonHarmony/models"
	"github.com/supolaris/salonHarmony/services"
)

func CreateUser(ctx *gin.Context) {
	var user models.Users
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "unable to bind user data",
		})
		return
	}
	err = services.CreateUser(&user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "unable to create user" + err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"result": user,
	})
}
