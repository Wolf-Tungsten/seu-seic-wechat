package wechat

import (
	"encoding/xml"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"wechat-bind/pkg/wechat"
)

func Handler(router *gin.RouterGroup) {

	route := router.Group("/wechat")
	route.GET("", GET)
	route.POST("", POST)

}

func GET(ctx *gin.Context) {
	// 处理 GET - /wechat 请求

	echostr := ctx.Query("echostr")

	if wechat.CheckWechatSignature(ctx) {
		ctx.String(http.StatusOK, echostr)
	} else {
		ctx.String(http.StatusBadRequest, "awsl")
	}

}

type wxMsg struct {
	ToUserName   string
	FromUserName string
	CreateTime   string
	MsgType      string
	Content      string
	MsgId        string
}

func POST(ctx *gin.Context) {
	// 处理 POST - /wechat 请求

	if !wechat.CheckWechatSignature(ctx) {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	buf, _ := ioutil.ReadAll(ctx.Request.Body)
	reqBody := string(buf[0:])
	fmt.Println(reqBody)
	var msg wxMsg

	_ = xml.Unmarshal(buf, &msg)
	fmt.Println(msg)
	openid := ctx.Query("openid")
	ctx.String(http.StatusOK, "success")

	wechat.SendCustomTextMsg(ctx, openid, "🤬")
}
