package handlers

import (
	"github.com/gin-gonic/gin"
	"wechat-bind/handlers/wechat"
	"wechat-bind/middleware"
)

func Handler(engine *gin.Engine) {

	router := engine.Group("/")
	router.Use(middleware.MongoConnect)
	router.Use(middleware.Cors)
	// 注册所有顶层handler
	wechat.Handler(router)

}
