package wechat

import (
	JSON "encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strings"
)

func SendCustomTextMsg(ctx *gin.Context, openid string, text string) {
	bodyBytes, _ := JSON.Marshal(map[string]interface{}{"touser": openid,
		"msgtype": "text",
		"text":    map[string]string{"content": text}})
	body := string(bodyBytes)
	body = strings.Replace(body, "\\u003c", "<", -1)
	body = strings.Replace(body, "\\u003e", ">", -1)
	body = strings.Replace(body, "\\u0026", "&", -1)
	fmt.Println(body)
	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token=%s", GetAccessToken(ctx))
	res, _ := http.Post(url, "application/json", strings.NewReader(body))
	wxRespBody, _ := ioutil.ReadAll(res.Body)
	wxRespBodyStruct := struct {
		Errcode int `json:"errcode"`
	}{}
	_ = JSON.Unmarshal(wxRespBody, &wxRespBodyStruct)
	if wxRespBodyStruct.Errcode != 0 {
		fmt.Println(string(wxRespBody))
	}
}
