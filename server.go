package main

import (
	"io"
	"os"

	"github.com/Nishith-Savla/golang-gin-poc/controller"
	"github.com/Nishith-Savla/golang-gin-poc/middlewares"
	"github.com/Nishith-Savla/golang-gin-poc/service"
	"github.com/gin-gonic/gin"
	"github.com/tpkeeper/gin-dump"
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

	server.Use(gin.Recovery(), middlewares.Logger(), middlewares.BasicAuth(), gindump.Dump())

	server.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.FindAll())
	})

	server.POST("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.Save(ctx))
	})

	server.Run("localhost:8080")
}
