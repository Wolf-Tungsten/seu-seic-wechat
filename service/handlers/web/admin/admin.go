package admin

import (
	"github.com/gin-gonic/gin"
	"wechat-bind/handlers/web/admin/admin"
	"wechat-bind/middleware"
)

func Handler(router *gin.RouterGroup) {

	route := router.Group("/admin")
	route.Use(middleware.AuthMiddleware)
	admin.Handler(route)

}
