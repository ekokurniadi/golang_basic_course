package main

import (
	"fmt"
	"log"
	"project_pertama/handler"
	"project_pertama/repository"
	"project_pertama/service"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/project_pertama?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	router := gin.Default()
	router.Use(CORSMiddleware())
	router.Static("/images", "./images")
	api := router.Group("/api/v1")

	api.POST("/users", userHandler.Save)
	api.POST("/users/avatar", userHandler.UploadProfilePicture)
	api.PUT("/users/:id", userHandler.UpdateUser)
	api.GET("/users", userHandler.GetAllUser)
	api.GET("/users/:id", userHandler.GetByID)
	api.DELETE("/users/:id", userHandler.DeleteUser)

	fmt.Println("Database berhasil konek")
	router.Run()

	//user -> handler -> service -> repository -> database
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT,DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
