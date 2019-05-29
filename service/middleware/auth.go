package middleware

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"wechat-bind/models"
	"wechat-bind/pkg"
)

func AuthMiddleware(c *gin.Context) {
	token := c.Request.Header.Get("session-token")
	db := c.MustGet("db").(*mongo.Database)
	// 如果token存在则进行权限鉴定
	if token != "" {
		var user models.User
		err := db.Collection("user").FindOne(c, bson.M{"sessionToken": token}).Decode(&user)
		if err == nil {
			c.Set("user", user)
			c.Next()
		} else {
			pkg.Return(c, 401, "身份认证无效")
			c.Abort()
		}
		// 处理请求

	} else {
		pkg.Return(c, 401, "需要身份认证")
		c.Abort()
	}

}
