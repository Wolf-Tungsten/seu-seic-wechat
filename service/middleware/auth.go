package middleware

import "github.com/gin-gonic/gin"

func userInfo(ctx *gin.Context){
	sessionToken := ctx.Request.Header.Get("session-token")
	

}
