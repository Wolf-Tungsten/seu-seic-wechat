package main

import (
	"github.com/gin-gonic/gin"
	"go-common/handlers"
)

const ServerPort = "3002"

func main() {

	engine := gin.Default()
	handlers.Handler(engine)
	engine.Run(":" + ServerPort)

}
