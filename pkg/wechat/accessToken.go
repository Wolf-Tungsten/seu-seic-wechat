package wechat

import (
	JSON "encoding/json"
	"fmt"
	"go-common/secret"
	"io/ioutil"
	"net/http"
	"time"
)

func GetAccessToken() string {

	if time.Now().Unix() >= secret.WechatAccessTokenWExpires {

		url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s", secret.WechatAppId, secret.WechatAppSecret)
		res, _ := http.Get(url)
		wxRespBody, _ := ioutil.ReadAll(res.Body)

		wxRespJSON := &struct {
			AccessToken string `json:"access_token"`
			ExpiresIn int64 `json:"expires_in"`
		}{}

		_ = JSON.Unmarshal(wxRespBody, &wxRespJSON)

		secret.WechatAccessToken = wxRespJSON.AccessToken

		secret.WechatAccessTokenWExpires = time.Now().Unix() + wxRespJSON.ExpiresIn

	}

	return secret.WechatAccessToken

}