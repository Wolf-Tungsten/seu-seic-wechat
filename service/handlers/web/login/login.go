package login

import (
	JSON "encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"io/ioutil"
	"log"
	"net/http"
	"wechat-bind/models"
	"wechat-bind/pkg"
	"wechat-bind/secret"
)

func Handler(router *gin.RouterGroup) {

	route := router.Group("/login")
	route.GET("", GET)
	route.POST("", POST)

}

/**
 * @api {get} /web/login 微信网页授权登录
 * @apiName Login
 * @apiGroup Login
 *
 * @apiParam {String} code 微信网页授权 code
 *
 * @apiSuccess {String} sessionToken 后续会话使用的 sessionToken
 * @apiSuccess {Bool} needInfo 是否需要补全用户信息
 */
func GET(ctx *gin.Context) {

	db := ctx.MustGet("db").(*mongo.Database)
	code := ctx.Query("code")
	// 换取 openid
	wxAuthUrl := fmt.Sprintf("https://api.weixin.qq.com/sns/oauth2/access_token?appid=%s&secret=%s&code=%s&grant_type=authorization_code", secret.WechatAppId, secret.WechatAppSecret, code)
	wxResp, _ := http.Get(wxAuthUrl)
	wxRespBody, _ := ioutil.ReadAll(wxResp.Body)
	fmt.Println(string(wxRespBody))
	wxRespJSON := &struct {
		OpenId      string `json:"openid"`
		AccessToken string `json:"access_token"`
		ErrMsg      string `json:"errmsg"`
	}{}

	_ = JSON.Unmarshal(wxRespBody, &wxRespJSON)

	if wxRespJSON.ErrMsg != "" {
		pkg.Return(ctx, 400, "微信认证出错")
		ctx.Abort()
		return
	}
	// 生成 sessionToken
	sessionTokenUUID, _ := uuid.NewRandom()
	sessionTokenStr := sessionTokenUUID.String()

	recordCount, _ := db.Collection("user").CountDocuments(ctx, bson.M{"openId": wxRespJSON.OpenId})

	switch recordCount {
	case 0:
		// 新用户注册的情况
		_, err := db.Collection("user").InsertOne(ctx, models.User{OpenId: wxRespJSON.OpenId, SessionToken: sessionTokenStr})
		if err != nil {
			log.Fatal(err)
		}
	case 1:
		// 已经注册，更新sessionToken
		_, _ = db.Collection("user").UpdateOne(ctx,
			bson.M{"openId": wxRespJSON.OpenId},
			bson.M{"$set": bson.M{"sessionToken": sessionTokenStr}})
	default:
		// 理论上不应该出现这种情况，如果出现则清除所有记录并报警
		_, _ = db.Collection("user").DeleteMany(ctx, bson.M{"openId": wxRespJSON.OpenId})
		pkg.Return(ctx, 400, "身份认证出现问题，请重试")
		return
	}

	var UserRecord models.User

	_ = db.Collection("user").FindOne(ctx, bson.M{"openId": wxRespJSON.OpenId}).Decode(&UserRecord)

	if UserRecord.Name == "" || UserRecord.Cardnum == "" {
		pkg.Return(ctx, 200, gin.H{"sessionToken": sessionTokenStr, "needUserInfo": true})
	} else {
		pkg.Return(ctx, 200, gin.H{"sessionToken": sessionTokenStr, "needUserInfo": false})
	}
}

/**
 * @api {post} /web/login 补全用户信息
 * @apiName UpdateUserInfo
 * @apiGroup Login
 *
 * @apiParam {String} name 真实姓名
 * @apiParam {String} cardnum 一卡通号
 */
func POST(ctx *gin.Context) {

}
