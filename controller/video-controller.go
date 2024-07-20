package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"learn-gin/entity"
	"learn-gin/service"
	validator2 "learn-gin/validator"
	"net/http"
)

type VideoController interface {
	FindAll() []entity.Video
	Save(context *gin.Context) error
	ShowAll(context *gin.Context)
}
type controller struct {
	service service.VideoService
}

var validate *validator.Validate

func New(service service.VideoService) VideoController {
	validate = validator.New()
	validate.RegisterValidation("is-cool", validator2.ValidationIsCool)
	return &controller{
		service: service,
	}
}
func (c *controller) FindAll() []entity.Video {
	return c.service.FindAll()
}

func (c *controller) Save(context *gin.Context) error {
	var video entity.Video
	err := context.ShouldBindJSON(&video)
	if err != nil {
		return err
	}
	err = validate.Struct(video)
	if err != nil {
		return err
	}
	c.service.Save(video)
	return nil
}
func (c *controller) ShowAll(context *gin.Context) {
	videos := c.service.FindAll()
	data := gin.H{
		"title":  "Video Page",
		"videos": videos,
	}
	context.HTML(http.StatusOK, "index.html", data)
}
