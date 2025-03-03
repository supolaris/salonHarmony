package main

import (
	"github.com/gin-gonic/gin"
	"github.com/supolaris/salonHarmony/db"
	"github.com/supolaris/salonHarmony/models"
	"github.com/supolaris/salonHarmony/routes"
)

func main() {
	db.InitDB()
	db.DB.AutoMigrate(&models.Users{}, &models.Services{}, &models.Staffs{}, &models.Appointments{})
	server := gin.Default()
	routes.InitRoutes(server)
	server.Run(":8080")
}
