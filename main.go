package main

import (
	"github.com/gin-gonic/gin"
	"wechat-bind/database"
	"wechat-bind/handlers"
)

const ServerPort = "3002"

func main() {

	database.Connect()
	engine := gin.Default()
	handlers.Handler(engine)
	engine.Run(":" + ServerPort)

}
