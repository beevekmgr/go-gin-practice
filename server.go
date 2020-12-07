package main

import (
	"golang-gin-poc/controller"
	"golang-gin-poc/middlewares"
	"golang-gin-poc/service"
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {
	setupLogOutput()

	server := gin.New()

	server.Use(gin.Recovery(), middlewares.Logger(), middlewares.BasicAuth(),
		gindump.Dump())

	server.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Ok!!",
		})
	})

	server.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.FindAll())
	})

	server.POST("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.Save(ctx))
	})

	server.Run(":8080")

}
