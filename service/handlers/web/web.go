package web

import (
	"github.com/gin-gonic/gin"
	"wechat-bind/handlers/web/user"
)

func Handler(router *gin.RouterGroup) {

	route := router.Group("/web")
	user.Handler(route)

}