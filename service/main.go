package main

import (
	"github.com/gin-gonic/gin"
	"wechat-bind/database"
	"wechat-bind/handlers"
	"wechat-bind/pkg/wechat"
)

const ServerPort = "3002"

func main() {

	database.Connect()
	wechat.CustomMenu() // 更新自定义菜单
	engine := gin.Default()
	handlers.Handler(engine)
	engine.Run(":" + ServerPort)

}
