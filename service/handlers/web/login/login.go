package login

import "github.com/gin-gonic/gin"

func Handler(router *gin.RouterGroup) {

	route := router.Group("/login")
	route.GET("", GET)
	route.POST("", POST)

}

/**
 * @api {get} /web/login 使用微信网页授权 code 登录
 * @apiName Login
 * @apiGroup Login
 *
 * @apiParam {String} code 微信网页授权 code
 *
 * @apiSuccess {String} sessionToken 后续会话使用的 sessionToken
 * @apiSuccess {Bool} needInfo 是否需要补全用户信息
 */
func GET(ctx *gin.Context){

}

/**
 * @api {post} /web/login 补全用户信息
 * @apiName UpdateUserInfo
 * @apiGroup Login
 *
 * @apiParam {String} name 真实姓名
 * @apiParam {String} cardnum 一卡通号
 * @apiParam {String} phoneNumber 联系电话
 */
func POST(ctx *gin.Context){

}

