package wechat

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/json"
	"net/http"
)

func SendCustomTextMsg(ctx *gin.Context, openid string, text string) {
	body, _ := json.Marshal(map[string]interface{}{"touser": openid,
		"msgtype": "text",
		"text":    map[string]string{"content": text}})
	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token=%s", GetAccessToken(ctx))
	_, _ = http.Post(url, "applicant/json", bytes.NewReader(body))
}
