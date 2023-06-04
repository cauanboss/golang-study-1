package http

import (
	"github.com/gin-gonic/gin"
	"log"
	"main/adapter/env"
	"main/domain/user"
)

func Start() {
	r := gin.Default()

	controller := user.NewController()
	r.GET("/user", controller.GetAll)
	r.POST("/user", controller.Create)
	r.GET("/user/:id", controller.Get)
	r.DELETE("/user/:id", controller.Delete)
	r.PUT("/user/:id", controller.Update)

	r.GET("/ping", ping)
	err := r.Run("127.0.0.1:" + env.GetHttpPort())
	if err != nil {
		log.Fatal(err)
	}
}

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
