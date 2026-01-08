package main

import (
	"RealTimeChatApplication/headlers"
	"RealTimeChatApplication/middleware"
	"RealTimeChatApplication/repository"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	if err := repository.Connect(); err != nil {
		log.Fatal(err)
	}
	if err := repository.ConnectRedis(); err != nil {
		log.Fatal(err)
	}

	r := gin.Default()

	r.POST("/login", headlers.Login)

	auth := r.Group("/")

	auth.Use(middleware.MidwareAuth())
	{
		auth.GET("/listRoom", headlers.ListRoom)
		auth.GET("/showChat", headlers.ShowChat)
	}

	r.Run(":8080")

}
