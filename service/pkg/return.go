package pkg

import "github.com/gin-gonic/gin"

func Return(ctx *gin.Context, code int, result interface{}) {

	if code >= 300 {
		ctx.JSON(200, gin.H{"success": false, "code": code, "reason": result})
	} else {
		ctx.JSON(200, gin.H{"success": true, "code": code, "result": result})
	}

}
