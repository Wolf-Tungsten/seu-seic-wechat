package wechat

import (
	JSON "encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"io/ioutil"
	"net/http"
	"time"
	"wechat-bind/secret"
)

func GetAccessToken(ctx *gin.Context) string {

	db := ctx.MustGet("db").(*mongo.Database)

	accessTokenStruct := struct {
		AccessToken string `bson:"accessToken" json:"access_token"`
		ExpiresTime int64  `bson:"expiresTime" json:"expires_in"`
	}{}

	err := db.Collection("accessToken").FindOne(ctx, bson.M{}).Decode(&accessTokenStruct)

	if err == mongo.ErrNoDocuments || time.Now().Unix() >= accessTokenStruct.ExpiresTime {
		// 不存在accessToken或者过期
		url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s", secret.WechatAppId, secret.WechatAppSecret)
		res, _ := http.Get(url)
		wxRespBody, _ := ioutil.ReadAll(res.Body)
		_ = JSON.Unmarshal(wxRespBody, &accessTokenStruct)
		accessTokenStruct.ExpiresTime = accessTokenStruct.ExpiresTime + time.Now().Unix()
		_, _ = db.Collection("accessToken").DeleteMany(ctx, bson.M{})
		_, _ = db.Collection("accessToken").InsertOne(ctx, accessTokenStruct)
	}

	return accessTokenStruct.AccessToken

}
