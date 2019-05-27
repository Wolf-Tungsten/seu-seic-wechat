package wechat

import (
	JSON "encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"io/ioutil"
	"log"
	"net/http"
	"time"
	"wechat-bind/secret"
)

func GetAccessToken(ctx *gin.Context) string {

	db := ctx.MustGet("db").(*mongo.Database)

	accessTokenStruct := struct {
		Errcode     int64  `json:"errcode"`
		AccessToken string `bson:"accessToken" json:"access_token"`
		ExpiresTime int64  `bson:"expiresTime" json:"expires_in"`
	}{}

	err := db.Collection("accessToken").FindOne(ctx, bson.M{}).Decode(&accessTokenStruct)

	if err == mongo.ErrNoDocuments {
		// 不存在accessToken或者过期
		fmt.Println("更新accessToken")
		url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s", secret.WechatAppId, secret.WechatAppSecret)
		res, _ := http.Get(url)
		wxRespBody, _ := ioutil.ReadAll(res.Body)
		_ = JSON.Unmarshal(wxRespBody, &accessTokenStruct)
		if accessTokenStruct.Errcode == 0 {
			accessTokenStruct.ExpiresTime = accessTokenStruct.ExpiresTime + time.Now().Unix()
			_, _ = db.Collection("accessToken").DeleteMany(ctx, bson.M{})
			_, _ = db.Collection("accessToken").InsertOne(ctx, accessTokenStruct)
		} else {
			fmt.Println(string(wxRespBody))
			log.Fatalln("AccessToken获取失败")
		}
	}
	return accessTokenStruct.AccessToken
}
