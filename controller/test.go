package controller

import (
	"gin/context"
	"gin/service"

	"github.com/gin-gonic/gin"
)

func TestGet(c *gin.Context) {
	p := new(service.TestServices)
	if err := c.ShouldBind(p); err != nil {
		context.ValidateError(c)
		return
	}
	service.Test()
	context.Success(c, nil)
}
