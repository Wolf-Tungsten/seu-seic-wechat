package wechat

import (
	"encoding/xml"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"io/ioutil"
	"net/http"
	"os"
	"wechat-bind/models"
	"wechat-bind/pkg/wechat"
	"wechat-bind/secret"
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
	Event        string
	EventKey     string
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

	if msg.MsgType == "event" && msg.Event == "CLICK" {
		wechat.SendCustomTextMsg(ctx, openid, msg.EventKey)
	}

	if msg.MsgType == "text" {
		switch msg.Content {
		case "初始化管理员":
			{
				initSuperAdmin(ctx, openid)
			}

		}
	}
}

// 初始化超级管理员

func initSuperAdmin(ctx *gin.Context, openid string) {
	db := ctx.MustGet("db").(*mongo.Database)
	superAdminCount, _ := db.Collection("user").CountDocuments(ctx, bson.M{"adminLevel": -1})
	if superAdminCount != 0 {
		wechat.SendCustomTextMsg(ctx, openid, "无权操作")
	} else {
		// 说明没有超级管理员
		// 删除该用户已有记录
		_, _ = db.Collection("user").DeleteMany(ctx, bson.M{"openId": openid})
		// 创建一个新记录
		_, _ = db.Collection("user").InsertOne(ctx, models.User{OpenId: openid, AdminLevel: -1})
		wechat.SendCustomTextMsg(ctx, openid, "您已成为该系统管理员！")
		wechat.SendCustomTextMsg(ctx, openid, fmt.Sprintf("<a href=\"https://open.weixin.qq.com/connect/oauth2/authorize?appid=%s&redirect_uri=%s/login/admin&response_type=code&scope=snsapi_base#wechat_redirect\">设置其他管理人员</a>", secret.WechatAppId, os.Getenv("HOMEPAGE")))
	}
}
