package routes

import (
	"gin/controller"

	"github.com/gin-gonic/gin"
)

func Test(apiGroup *gin.RouterGroup) {
	apiUserGroup := apiGroup.Group("/test")
	// 用户登入
	apiUserGroup.GET("/login", controller.TestGet)

}
