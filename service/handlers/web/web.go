package web

import (
	"github.com/gin-gonic/gin"
	"wechat-bind/handlers/web/login"
)

func Handler(router *gin.RouterGroup) {

	route := router.Group("/web")
	login.Handler(route)

}
