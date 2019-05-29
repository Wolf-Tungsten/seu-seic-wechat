package user

import "github.com/gin-gonic/gin"

func Handler(router *gin.RouterGroup) {

	route := router.Group("/user")
	route.GET("", GET)
	route.POST("", POST)

}

func GET(ctx *gin.Context){

}

func POST(ctx *gin.Context) {

}