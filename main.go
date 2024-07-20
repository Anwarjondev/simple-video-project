package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"learn-gin/controller"
	"learn-gin/middlewares"
	"learn-gin/service"
	"net/http"
	"os"
)

var (
	videoService   service.VideoService       = service.New()
	videoContrller controller.VideoController = controller.New(videoService)
)

func LoggerFile() {
	f, _ := os.Create("log.go")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {
	LoggerFile()

	r := gin.New()

	r.Static("/css", "./templates/css")

	r.LoadHTMLGlob("templates/*.html")

	r.Use(gin.Recovery(), middlewares.Logger(), middlewares.BasicAuth())

	apiRoutes := r.Group("/api")
	{
		apiRoutes.GET("/videos", func(context *gin.Context) {
			context.JSON(200, videoContrller.FindAll())
		})

		apiRoutes.POST("/videos", func(context *gin.Context) {
			err := videoContrller.Save(context)
			if err != nil {
				context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				context.JSON(http.StatusOK, gin.H{"message": "Video is Successfully saved"})
			}
		})
	}

	viewRoutes := r.Group("/view")
	{
		viewRoutes.GET("/videos", videoContrller.ShowAll)
	}
	r.Run(":8080")
}
