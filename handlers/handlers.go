package handlers

import (
	"github.com/gin-gonic/gin"
	"go-common/handlers/wechat"
)

func Handler(engine *gin.Engine){

	router := engine.Group("/")

	// 注册所有顶层handler
	wechat.Handler(router)

}