package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mrandi96/personal-gin/controller"
	"github.com/mrandi96/personal-gin/service"
)

var (
	videoService    service.VideoService       = service.New()
	VideoController controller.VideoController = controller.New(videoService)
)

func main() {
	server := gin.Default()

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "PONG",
		})
	})

	server.GET("/post", func(ctx *gin.Context) {
		ctx.JSON(200, VideoController.FindAll())
	})

	server.POST("/post", func(ctx *gin.Context) {
		ctx.JSON(201, VideoController.Save(ctx))
	})

	server.Run(":8080")
}
