package web

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"wechat-bind/handlers/web/admin"
	"wechat-bind/handlers/web/login"
	"wechat-bind/pkg"
	"wechat-bind/secret"
)

func Handler(router *gin.RouterGroup) {

	route := router.Group("/web")
	route.GET("", GET)
	login.Handler(route)
	admin.Handler(route)

}

func GET(ctx *gin.Context) {
	currentPath := ctx.Query("path")
	pkg.Return(ctx, 200, fmt.Sprintf("https://open.weixin.qq.com/connect/oauth2/authorize?appid=%s&redirect_uri=%s/login/%s&response_type=code&scope=snsapi_base#wechat_redirect", secret.WechatAppId, os.Getenv("HOMEPAGE"), currentPath))
}
