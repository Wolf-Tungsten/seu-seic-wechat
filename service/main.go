package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
	"wechat-bind/database"
	"wechat-bind/handlers"
	"wechat-bind/pkg/wechat"
)

const ServerPort = "3002"

func main() {

	database.Connect()
	wechat.CustomMenu() // 更新自定义菜单
	engine := gin.Default()
	engine.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"PUT", "POST", "DELETE"},
		AllowHeaders:  []string{"Origin", "Token", "Content-Length", "Content-Type"},
		ExposeHeaders: []string{"Content-Length", "Content-Type"},
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	handlers.Handler(engine)
	engine.Run(":" + ServerPort)

}
