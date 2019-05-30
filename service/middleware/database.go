package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"wechat-bind/database"
	"wechat-bind/secret"
)

func MongoConnect(context *gin.Context) {
	fmt.Print("数据库中间件进入")
	context.Set("db", database.Client.Database(secret.DatabaseName))
	context.Next()
	fmt.Print("数据库中间件离开")
}
