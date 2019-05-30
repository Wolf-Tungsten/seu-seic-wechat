package middleware

import (
	"github.com/gin-gonic/gin"
	"wechat-bind/database"
	"wechat-bind/secret"
)

func MongoConnect(context *gin.Context) {
	context.Set("db", database.Client.Database(secret.DatabaseName))
	context.Next()
}
