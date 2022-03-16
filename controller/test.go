package controller

import (
	"gin/context"

	"github.com/gin-gonic/gin"
)

func TestGet(c *gin.Context) {
	context.Success(c, nil, nil)
}
