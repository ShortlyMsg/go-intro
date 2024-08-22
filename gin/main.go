package main

import (
	"gin/controllers"
	"gin/db"
	"gin/models"

	"github.com/gin-gonic/gin"
)

func main() {

	db.ConnectDatabase()

	db.DB.AutoMigrate(&models.User{})

	r := gin.Default()

	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	r.Run(":8080")
}
